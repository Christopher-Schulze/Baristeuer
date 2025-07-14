'use client';

import { useAtom } from 'jotai';
import { currentStepAtom, steuererklaerungAtom } from '@/lib/steuer/formState';
import { Card, CardBody, CardFooter, CardHeader } from '@heroui/react';
import { Input, Button, Checkbox, Tooltip } from '@heroui/react';
import { Info } from 'lucide-react';
 
 export function Schritt1_Grunddaten() {
   const [steuererklaerung, setSteuererklaerung] = useAtom(steuererklaerungAtom);
  const [, setCurrentStep] = useAtom(currentStepAtom);

  return (
    <div className="space-y-4">
      <div className="space-y-2">
        <label htmlFor="vereinsname">Name des Vereins</label>
        <Input
          id="vereinsname"
          value={steuererklaerung.verein.name}
          onChange={(e) =>
            setSteuererklaerung((prev) => ({
              ...prev,
              verein: { ...prev.verein, name: e.target.value },
            }))
          }
          placeholder="z.B. Musterverein e.V."
        />
      </div>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
       <div className="space-y-2 md:col-span-2">
           <label htmlFor="strasse">Straße und Hausnummer</label>
           <Input
           id="strasse"
           value={steuererklaerung.verein.adresse.strasse}
           onChange={(e) =>
               setSteuererklaerung((prev) => ({
               ...prev,
               verein: { ...prev.verein, adresse: { ...prev.verein.adresse, strasse: e.target.value } },
               }))
           }
           placeholder="Musterstraße 123"
           />
       </div>
       <div className="space-y-2">
           <label htmlFor="plz">PLZ</label>
           <Input
           id="plz"
           value={steuererklaerung.verein.adresse.plz}
           onChange={(e) =>
               setSteuererklaerung((prev) => ({
               ...prev,
               verein: { ...prev.verein, adresse: { ...prev.verein.adresse, plz: e.target.value } },
               }))
           }
           placeholder="12345"
           />
       </div>
       <div className="space-y-2 md:col-span-2">
           <label htmlFor="ort">Ort</label>
           <Input
           id="ort"
           value={steuererklaerung.verein.adresse.ort}
           onChange={(e) =>
               setSteuererklaerung((prev) => ({
               ...prev,
               verein: { ...prev.verein, adresse: { ...prev.verein.adresse, ort: e.target.value } },
               }))
           }
           placeholder="Musterstadt"
           />
       </div>
      </div>
      <div className="space-y-2">
        <div className="flex items-center">
            <label htmlFor="steuernummer">Steuernummer</label>
            <Tooltip content="Die vom Finanzamt für den Verein vergebene Steuernummer.">
              <Info className="h-4 w-4 text-muted-foreground cursor-pointer ml-2" />
            </Tooltip>
        </div>
          <Input
          id="steuernummer"
          value={steuererklaerung.verein.steuernummer || ''}
          onChange={(e) =>
              setSteuererklaerung((prev) => ({
              ...prev,
              verein: { ...prev.verein, steuernummer: e.target.value },
              }))
          }
          placeholder="z.B. 123/456/7890"
          />
      </div>
      <div className="space-y-2">
           <label htmlFor="finanzamtAdresse">Adresse des zuständigen Finanzamts</label>
           <Input
           id="finanzamtAdresse"
           value={steuererklaerung.verein.finanzamtAdresse}
           onChange={(e) =>
               setSteuererklaerung((prev) => ({
               ...prev,
               verein: { ...prev.verein, finanzamtAdresse: e.target.value },
               }))
           }
           placeholder="Finanzamt Musterstadt, Musterweg 1, 12345 Musterstadt"
           />
       </div>
      <div className="space-y-2">
        <div className="flex items-center">
          <label htmlFor="hebesatz">Hebesatz der Gemeinde (%)</label>
            <Tooltip content="Der Hebesatz ist ein von der Gemeinde festgelegter Prozentsatz zur Berechnung der Gewerbesteuer.">
              <Info className="h-4 w-4 text-muted-foreground cursor-pointer ml-2" />
            </Tooltip>
        </div>
        <Input
          id="hebesatz"
          type="number"
          min="0"
          value={steuererklaerung.verein.hebesatz.toString()}
          onChange={(e) => {
            const value = parseFloat(e.target.value);
             // Allow clearing the field, otherwise enforce non-negative
            if (e.target.value === '') {
                 setSteuererklaerung((prev) => ({
                    ...prev,
                    verein: { ...prev.verein, hebesatz: 0 }, // Or handle as null/undefined if schema allows
                }));
            } else if (!isNaN(value) && value >= 0) {
              setSteuererklaerung((prev) => ({
                ...prev,
                verein: { ...prev.verein, hebesatz: value },
              }));
            }
          }}
          placeholder="z.B. 400"
        />
      </div>
      <div className="flex items-center space-x-2 pt-2">
        <Checkbox
          id="kleinunternehmer"
          checked={steuererklaerung.verein.kleinunternehmerregelung}
          onChange={(e) => {
            setSteuererklaerung((prev) => ({
              ...prev,
              verein: { ...prev.verein, kleinunternehmerregelung: e.target.checked },
            }));
          }}
        />
        <div className="flex items-center">
          <label htmlFor="kleinunternehmer">
            Kleinunternehmerregelung in Anspruch nehmen
          </label>
            <Tooltip content="Wenn die umsatzsteuerpflichtigen Einnahmen im Vorjahr unter 22.000€ und im laufenden Jahr voraussichtlich unter 50.000€ liegen, kann diese Regelung genutzt werden. Es wird keine Umsatzsteuer fällig, aber auch kein Vorsteuerabzug ist möglich.">
              <Info className="h-4 w-4 text-muted-foreground cursor-pointer ml-2" />
            </Tooltip>
        </div>
      </div>
    </div>
  );
}