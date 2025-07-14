import { atom } from 'jotai';
// The task instruction mentions 'Steuererklaerung', but the actual exported type from the validation file
// is 'Steuerdaten'. We use the correct type but name the atom as requested.
import type { VereinSteuererklaerung } from '../../validations/steuerdaten';

/**
 * Defines the initial state for a new tax declaration.
 * This state is derived from the Zod schema defaults in `steuerdaten.ts`.
 * It ensures that the form state is fully initialized with type-safe default values
 * before any user interaction.
*/
const initialState: VereinSteuererklaerung = {
  jahr: new Date().getFullYear(),
  verein: {
    name: '',
    adresse: {
      strasse: '',
      plz: '',
      ort: '',
    },
    finanzamtAdresse: '',
    hebesatz: 400,
    kleinunternehmerregelung: true,
  },
  ideellerBereich: {
    einnahmen: [],
    ausgaben: [],
    angewandtePauschalen: 0,
  },
  vermoegensverwaltung: {
    einnahmen: [],
    ausgaben: [],
    angewandtePauschalen: 0,
  },
  zweckbetrieb: {
    einnahmen: [],
    ausgaben: [],
    angewandtePauschalen: 0,
  },
  wirtschaftlicherGeschaeftsbetrieb: {
    einnahmen: [],
    ausgaben: [],
    angewandtePauschalen: 0,
  },
};

/**
 * Jotai atom to hold the entire state of the tax declaration form data.
 */
export const steuererklaerungAtom = atom<VereinSteuererklaerung>(initialState);

/**
 * Jotai atom to track the current step in the multi-step interview process.
 * It starts at 0, representing the beginning of the interview.
 */
export const currentStepAtom = atom(0);