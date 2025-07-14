import { VereinSteuererklaerung as Steuerdaten } from '../../validations/steuerdaten';
import { SteuerEngine, TAX_CONSTANTS } from './engine';

// The new structure for the calculation result
export interface BerechnungsErgebnisVerein {
  ideellerBereich: number;
  vermoegensverwaltung: number;
  zweckbetrieb: number;
  wirtschaftlicherGeschaeftsbetrieb: number;
  koerperschaftssteuer: number;
  gewerbesteuer: number;
  umsatzsteuer: number;
  gesamtsteuerlast: number;
}

/**
 * Acts as a facade for the tax calculation logic.
 * It uses the SteuerEngine to perform the actual calculations
 * and formats the results.
 */
export class SteuerService {
  private engine: SteuerEngine;

  constructor() {
    this.engine = new SteuerEngine();
  }

  /**
   * Main method to calculate taxes for a non-profit organization.
   * @param steuerdaten - The complete tax data for the organization.
   * @returns The calculated tax results for each area.
   */
  public berechneSteuern(steuerdaten: Steuerdaten): BerechnungsErgebnisVerein {
    if (!steuerdaten) {
      throw new Error('Keine Steuerdaten übergeben');
    }

    const { verein, ...bereiche } = steuerdaten;

    // 1. Calculate the initial results for all four spheres
    const einzelErgebnisse = {
      ideellerBereich: this.engine.berechneErgebnisFuerBereich(bereiche.ideellerBereich),
      vermoegensverwaltung: this.engine.berechneErgebnisFuerBereich(bereiche.vermoegensverwaltung),
      zweckbetrieb: this.engine.berechneErgebnisFuerBereich(bereiche.zweckbetrieb),
      wirtschaftlicherGeschaeftsbetrieb: this.engine.berechneErgebnisWirtschaftlicherGeschaeftsbetrieb(bereiche.wirtschaftlicherGeschaeftsbetrieb),
    };

    // 2. Pass these results to the offsetting logic (now part of Gesamtergebnis)
    // The `berechneGesamtergebnis` method now encapsulates the offsetting logic.
    const { finaleErgebnisse } = this.engine.berechneGesamtergebnis(steuerdaten);

    // 3. Use the final, offset result of the WGB for tax calculations
    let einkommenFuerSteuerberechnung = finaleErgebnisse.wirtschaftlicherGeschaeftsbetrieb;

    // Check if the gross income of WGB is below the tax-exempt limit.
    // If so, the profit is not taxed, even if it was used for offsetting losses.
    const bruttoEinnahmenWGB = this.engine.summiereBetraege(bereiche.wirtschaftlicherGeschaeftsbetrieb.einnahmen);
    if (bruttoEinnahmenWGB <= TAX_CONSTANTS.FREIGRENZE_WGB) {
        einkommenFuerSteuerberechnung = 0;
    }

    // 4. Calculate the actual taxes based on the final, potentially adjusted income
    const { koerperschaftsteuer, solidaritätszuschlag } = this.engine.berechneKoerperschaftsteuer(einkommenFuerSteuerberechnung);
    
    // Assuming a default Hebesatz if not provided. This should come from `VereinStammdaten` in the future.
    const hebesatz = steuerdaten.verein.hebesatz || 400;
    const gewerbesteuer = this.engine.berechneGewerbesteuer(einkommenFuerSteuerberechnung, hebesatz);

    // 5. Calculate VAT
    const { zahllast: umsatzsteuer } = this.engine.berechneUmsatzsteuer(steuerdaten);

    const kstGesamt = koerperschaftsteuer + solidaritätszuschlag;
    const gesamtsteuerlast = kstGesamt + gewerbesteuer + umsatzsteuer;

    return {
      ideellerBereich: finaleErgebnisse.ideellerBereich,
      vermoegensverwaltung: finaleErgebnisse.vermoegensverwaltung,
      zweckbetrieb: finaleErgebnisse.zweckbetrieb,
      wirtschaftlicherGeschaeftsbetrieb: finaleErgebnisse.wirtschaftlicherGeschaeftsbetrieb,
      koerperschaftssteuer: kstGesamt,
      gewerbesteuer,
      umsatzsteuer,
      gesamtsteuerlast,
    };
  }
}
