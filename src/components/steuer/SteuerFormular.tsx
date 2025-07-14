'use client';

import { useAtom } from 'jotai';
import { useEffect, useRef, useState } from 'react';
import { currentStepAtom, steuererklaerungAtom } from '@/lib/steuer/formState';
import { api } from '@/lib/trpc/client';
import { Toaster, toast } from 'sonner';
import { VereinSteuererklaerungSchema } from '@/validations/steuerdaten';
import { Schritt1_Grunddaten } from './interview-steps/Schritt1_Grunddaten';
import { Schritt2_IdeellerBereich } from './interview-steps/Schritt2_IdeellerBereich';
import { Schritt3_Vermoegensverwaltung } from './interview-steps/Schritt3_Vermoegensverwaltung';
import { Schritt4_Zweckbetrieb } from './interview-steps/Schritt4_Zweckbetrieb';
import { Schritt5_WirtschaftlicherGeschaeftsbetrieb } from './interview-steps/Schritt5_WirtschaftlicherGeschaeftsbetrieb';
import { Button } from '@heroui/react';
import {
  Card,
  CardBody as CardContent,
  CardFooter,
  CardHeader,
} from '@heroui/react';
import { Loader2 } from 'lucide-react';
// CardTitle and CardDescription are not separate components in HeroUI, but props of CardHeader.
 
 interface SteuerFormularProps {
   vereinId: string;
  jahr: number;
}

export function SteuerFormular({ vereinId, jahr }: SteuerFormularProps) {
  const [currentStep, setCurrentStep] = useAtom(currentStepAtom);
  const [steuererklaerung, setSteuererklaerung] = useAtom(steuererklaerungAtom);
  const [isSaving, setIsSaving] = useState(false);
  
  const { mutate: upsertSteuererklaerung } = api.steuer.upsert.useMutation({
    onMutate: () => {
      setIsSaving(true);
    },
    onSuccess: () => {
      toast.success('Änderungen erfolgreich gespeichert.');
      setIsSaving(false);
    },
    onError: (error) => {
      toast.error('Fehler beim Speichern', {
        description: error.message,
      });
      setIsSaving(false);
    },
  });
  const { mutateAsync: generatePdf, isPending: isGeneratingPdf } = api.steuer.generatePdf.useMutation();
 
  const debounceTimer = useRef<NodeJS.Timeout | null>(null);
  const isInitialMount = useRef(true);

  const { data: initialData, isLoading } = api.steuer.get.useQuery(
    { vereinId, jahr },
    {
      enabled: !!vereinId && !!jahr,
      refetchOnWindowFocus: false,
    }
  );

  useEffect(() => {
    if (initialData) {
      try {
        const parsedData = VereinSteuererklaerungSchema.parse(initialData);
        setSteuererklaerung(parsedData);
      } catch (error) {
        console.error("Validation error on initial data:", error);
        toast.error("Fehler beim Laden der Daten", {
          description: "Die vom Server empfangenen Daten sind nicht im erwarteten Format.",
        });
      }
    }
  }, [initialData, setSteuererklaerung]);

  useEffect(() => {
   // Prevent auto-saving when the form is first loading initial data.
   if (isLoading) {
     return;
   }

   if (debounceTimer.current) {
     clearTimeout(debounceTimer.current);
   }

   debounceTimer.current = setTimeout(() => {
     const dataToSave = {
       ...steuererklaerung,
       vereinId: vereinId,
       jahr: jahr,
     };
     upsertSteuererklaerung(dataToSave);
   }, 1500);

   return () => {
     if (debounceTimer.current) {
       clearTimeout(debounceTimer.current);
     }
   };
 }, [steuererklaerung, upsertSteuererklaerung, isLoading, vereinId, jahr]);

  const allSteps = [
    {
      id: 0,
      label: 'Grunddaten',
      component: <Schritt1_Grunddaten />,
      isUmsatzsteuerRelevant: false,
    },
    {
      id: 1,
      label: 'Ideeller Bereich',
      component: <Schritt2_IdeellerBereich />,
      isUmsatzsteuerRelevant: false,
    },
    {
      id: 2,
      label: 'Vermögensverwaltung',
      component: <Schritt3_Vermoegensverwaltung />,
      isUmsatzsteuerRelevant: false,
    },
    {
      id: 3,
      label: 'Zweckbetrieb',
      component: <Schritt4_Zweckbetrieb />,
      isUmsatzsteuerRelevant: true,
    },
    {
      id: 4,
      label: 'Wirtschaftlicher Geschäftsbetrieb',
      component: <Schritt5_WirtschaftlicherGeschaeftsbetrieb />,
      isUmsatzsteuerRelevant: true,
    },
  ];

  const visibleSteps = steuererklaerung.verein.kleinunternehmerregelung
    ? allSteps.filter(step => !step.isUmsatzsteuerRelevant)
    : allSteps;

  const currentVisibleIndex = visibleSteps.findIndex(step => step.id === currentStep);
  const isLastStep = currentVisibleIndex === visibleSteps.length - 1;
  const isFirstStep = currentVisibleIndex === 0;

  const handleNext = () => {
    if (!isLastStep) {
      setCurrentStep(visibleSteps[currentVisibleIndex + 1].id);
    }
  };

  const handlePrev = () => {
    if (!isFirstStep) {
      setCurrentStep(visibleSteps[currentVisibleIndex - 1].id);
    }
  };

  const renderCurrentStep = () => {
    const step = visibleSteps.find(s => s.id === currentStep);
    return step ? step.component : <div>Schritt nicht gefunden</div>;
  };

  const downloadPdf = (base64String: string, filename: string) => {
    const byteCharacters = atob(base64String);
    const byteNumbers = new Array(byteCharacters.length);
    for (let i = 0; i < byteCharacters.length; i++) {
      byteNumbers[i] = byteCharacters.charCodeAt(i);
    }
    const byteArray = new Uint8Array(byteNumbers);
    const blob = new Blob([byteArray], { type: 'application/pdf' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = filename;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  };

  const handleGeneratePdf = async () => {
    try {
      const base64String = await generatePdf(steuererklaerung);
      downloadPdf(base64String, `Steuererklaerung_${vereinId}_${jahr}.pdf`);
    } catch (error: any) {
      toast.error("PDF-Generierung fehlgeschlagen", {
        description: error.message || 'Ein unbekannter Fehler ist aufgetreten.',
      });
    }
  };

  return (
    <>
    <Toaster position="top-right" />
    <Card>
      <CardHeader>
        <div>
          <h2 className="text-xl font-semibold">Steuererklärung für Vereine</h2>
          <p className="text-sm text-muted-foreground">Führt Sie schrittweise durch die Steuererklärung.</p>
          {isSaving && (
              <div className="flex items-center text-muted-foreground mt-2">
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  <span>Speichern...</span>
              </div>
          )}
        </div>
      </CardHeader>
      <CardContent>{renderCurrentStep()}</CardContent>
      <CardFooter>
        <Button
          variant="bordered"
          onClick={handlePrev}
          disabled={isFirstStep}
        >
          Zurück
        </Button>
        <div>
          {isLastStep && (
            <Button onClick={handleGeneratePdf} disabled={isGeneratingPdf} variant="ghost" className="mr-4">
              {isGeneratingPdf ? 'PDF wird generiert...' : 'PDF generieren'}
            </Button>
          )}
          <Button
            onClick={handleNext}
            disabled={isLastStep}
          >
            {isLastStep ? 'Abschließen' : 'Weiter'}
          </Button>
        </div>
      </CardFooter>
    </Card>
    </>
  );
}
