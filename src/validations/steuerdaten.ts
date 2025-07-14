import { z } from 'zod';

// Base schema for a single financial entry
export const EintragSchema = z.object({
  bezeichnung: z.string(),
  betrag: z.number().nonnegative(),
  ustSatz: z.number().optional(),
  alsSpendeBehandeln: z.boolean().optional(),
});
export type Eintrag = z.infer<typeof EintragSchema>;

// Schema for one of the four tax sections of a club
export const SteuerbereichSchema = z.object({
  einnahmen: z.array(EintragSchema).default([]),
  ausgaben: z.array(EintragSchema).default([]),
  angewandtePauschalen: z.number().default(0),
});
export type Steuerbereich = z.infer<typeof SteuerbereichSchema>;

// Schema for the club's basic information
export const VereinStammdatenSchema = z.object({
  name: z.string().default(''),
  adresse: z.object({
    strasse: z.string().default(''),
    plz: z.string().default(''),
    ort: z.string().default(''),
  }).default({}),
  steuernummer: z.string().optional(),
  finanzamtAdresse: z.string().default(''),
  hebesatz: z.number().default(400),
  kleinunternehmerregelung: z.boolean().default(true),
});
export type VereinStammdaten = z.infer<typeof VereinStammdatenSchema>;

// The main schema for the entire club tax declaration
export const VereinSteuererklaerungSchema = z.object({
  id: z.string().optional(),
  jahr: z.number().int().min(2000).max(2100),
  verein: VereinStammdatenSchema.default({}),
  ideellerBereich: SteuerbereichSchema.default({}),
  vermoegensverwaltung: SteuerbereichSchema.default({}),
  zweckbetrieb: SteuerbereichSchema.default({}),
  wirtschaftlicherGeschaeftsbetrieb: SteuerbereichSchema.default({}),
  createdAt: z.date().optional(),
  updatedAt: z.date().optional(),
});

// The main type derived from the schema
export type VereinSteuererklaerung = z.infer<typeof VereinSteuererklaerungSchema>;
