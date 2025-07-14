import { useState } from 'react';
import { useAtom } from 'jotai';
import { steuererklaerungAtom } from '@/lib/steuer/formState';
import { type Eintrag } from '@/validations/steuerdaten';
import { type ErrorState } from './types';

type BereichKey = 'ideellerBereich' | 'vermoegensverwaltung' | 'zweckbetrieb' | 'wirtschaftlicherGeschaeftsbetrieb';
type EintragType = 'einnahmen' | 'ausgaben';

export const useEintragHandlers = (bereichKey: BereichKey) => {
  const [steuererklaerung, setSteuererklaerung] = useAtom(steuererklaerungAtom);
  const [errors, setErrors] = useState<ErrorState>({});

  const handleEintragChange = (
    type: EintragType,
    index: number,
    field: keyof Eintrag,
    value: string | number
  ) => {
    const bereich = steuererklaerung[bereichKey];
    const updatedEintraege = [...bereich[type]];
    const target = updatedEintraege[index];
    
    if (target) {
      if (field === 'betrag') {
        const rawValue = value.toString();
        (target as any).rawBetrag = rawValue;
        
        // Nur bei onBlur (wenn value eine Zahl ist) validieren und konvertieren
        if (typeof value === 'number') {
          const numValue = value;
          if (isNaN(numValue) || numValue < 0) {
            setErrors(prev => ({
              ...prev,
              [type]: {
                ...prev[type],
                [index]: 'Bitte geben Sie einen gültigen, positiven Betrag ein.'
              }
            }));
            target.betrag = 0;
          } else {
            setErrors(prev => {
              const newErrors = { ...prev[type] } as { [key: number]: string };
              delete newErrors[index];
              return { ...prev, [type]: newErrors };
            });
            target.betrag = numValue;
          }
        } else {
          // Bei String-Input (während der Eingabe) nur rawBetrag setzen
          // target.betrag bleibt unverändert bis onBlur
        }
      } else {
        (target[field] as any) = value;
      }
      
      setSteuererklaerung({
        ...steuererklaerung,
        [bereichKey]: {
          ...bereich,
          [type]: updatedEintraege,
        },
      });
    }
  };

  const addEintrag = (type: EintragType) => {
    const bereich = steuererklaerung[bereichKey];
    setSteuererklaerung({
      ...steuererklaerung,
      [bereichKey]: {
        ...bereich,
        [type]: [...bereich[type], { bezeichnung: '', betrag: 0 }],
      },
    });
  };

  const removeEintrag = (type: EintragType, index: number) => {
    const bereich = steuererklaerung[bereichKey];
    const updatedEintraege = bereich[type].filter((_, i) => i !== index);
    setSteuererklaerung({
      ...steuererklaerung,
      [bereichKey]: {
        ...bereich,
        [type]: updatedEintraege,
      },
    });
  };

  const handleEinnahmeChange = (index: number, field: keyof Eintrag, value: string | number) =>
    handleEintragChange('einnahmen', index, field, value);

  const handleAusgabeChange = (index: number, field: keyof Eintrag, value: string | number) =>
    handleEintragChange('ausgaben', index, field, value);

  const addEinnahme = () => addEintrag('einnahmen');
  const addAusgabe = () => addEintrag('ausgaben');
  const removeEinnahme = (index: number) => removeEintrag('einnahmen', index);
  const removeAusgabe = (index: number) => removeEintrag('ausgaben', index);

  return {
    steuererklaerung,
    errors,
    handleEinnahmeChange,
    handleAusgabeChange,
    addEinnahme,
    addAusgabe,
    removeEinnahme,
    removeAusgabe,
  };
};