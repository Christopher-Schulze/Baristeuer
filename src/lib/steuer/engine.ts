// src/lib/steuer/engine.ts

/**
 * This file contains the core tax calculation engine.
 * It is designed to be modular and testable, separating the logic
 * for different tax areas of a non-profit organization.
 */
import {
 VereinSteuererklaerung,
 Steuerbereich,
 Eintrag as SteuerEintrag,
} from '@/validations/steuerdaten';

/**
 * Central repository for all tax-related constants.
 * This improves maintainability and provides context for the legal basis of each value.
 */
export const TAX_CONSTANTS = {
  // § 64 Abs. 3 AO: The tax-exempt limit for income from economic business operations.
  // If total gross income is below this, any resulting profit is not subject to KSt or GewSt.
  FREIGRENZE_WGB: 45000,

  // § 24 KStG, § 11 Abs. 1 S. 3 Nr. 2 GewStG: The tax-free allowance for corporate and trade tax.
  // This amount is deducted from the taxable income before calculating the tax.
  FREIBETRAG_KST_GEWST: 5000,

  // § 23 Abs. 1 KStG: The standard corporate income tax rate in Germany.
  KST_SATZ: 0.15,

  // SolzG 1995: The solidarity surcharge rate, calculated on the corporate income tax.
  SOLI_SATZ: 0.055,

  // § 11 Abs. 2 GewStG: The basic federal rate for trade tax calculation.
  GEWST_MESSZAHL: 0.035,

  // A general constant for percentage calculations and rounding.
  BASIS_HUNDERT: 100,
};

export class SteuerEngine {
  /**
   * Calculates the sum of a list of amounts.
   * @param betraege - An array of Betrag objects.
   * @returns The total sum of the 'wert' properties.
   */
  public summiereBetraege(betraege: SteuerEintrag[]): number {
    return betraege.reduce((summe, posten) => summe + posten.betrag, 0);
  }

  /**
   * Calculates the result (surplus or deficit) for a single tax area.
   * @param bereich - A Steuerbereich object with income and expenses.
   * @returns The calculated result (Einnahmen - Ausgaben).
   */
  public berechneErgebnisFuerBereich(bereich: Steuerbereich): number {
    const summeEinnahmen = this.summiereBetraege(bereich.einnahmen);
    const summeAusgaben = this.summiereBetraege(bereich.ausgaben);
    const pauschalen = bereich.angewandtePauschalen ?? 0;
    return summeEinnahmen - summeAusgaben - pauschalen;
  }

  /**
   * Calculates the result for the economic business segment.
   * This method now calculates the true result (profit or loss) without applying the tax-exempt limit.
   * The 45,000 EUR limit is applied later in the process to determine if a profit is taxable.
   * @param bereich - The Steuerbereich for the economic business segment.
   * @returns The actual result (Einnahmen - Ausgaben).
   */
  public berechneErgebnisWirtschaftlicherGeschaeftsbetrieb(bereich: Steuerbereich): number {
    // This now correctly returns the true result, including potential losses.
    // The check against the FREIGRENZE_WGB (45,000€) is handled by the consuming service
    // to decide if a *profit* is subject to taxation.
    return this.berechneErgebnisFuerBereich(bereich);
  }

  /**
   * Calculates the corporate income tax (Körperschaftsteuer) and the solidarity surcharge (Solidaritätszuschlag).
   * @param einkommenWGB - The income from the economic business segment.
   * @returns An object containing the calculated taxes.
   */
  public berechneKoerperschaftsteuer(einkommenWGB: number): { koerperschaftsteuer: number; solidaritätszuschlag: number } {
    if (einkommenWGB <= TAX_CONSTANTS.FREIBETRAG_KST_GEWST) {
      return { koerperschaftsteuer: 0, solidaritätszuschlag: 0 };
    }

    const steuerpflichtigesEinkommen = einkommenWGB - TAX_CONSTANTS.FREIBETRAG_KST_GEWST;

    const koerperschaftsteuer = steuerpflichtigesEinkommen * TAX_CONSTANTS.KST_SATZ;
    const solidaritätszuschlag = koerperschaftsteuer * TAX_CONSTANTS.SOLI_SATZ;

    return { koerperschaftsteuer, solidaritätszuschlag };
  }

  /**
   * Calculates the trade tax (Gewerbesteuer).
   * @param einkommenWGB - The income from the economic business segment.
   * @param hebesatz - The municipal multiplier (Hebesatz) in percent (e.g., 400).
   * @returns The calculated trade tax.
   */
  public berechneGewerbesteuer(einkommenWGB: number, hebesatz: number): number {
    if (einkommenWGB <= TAX_CONSTANTS.FREIBETRAG_KST_GEWST) {
      return 0;
    }

    // Round down to the nearest 100
    const gerundeterErtrag = Math.floor(einkommenWGB / TAX_CONSTANTS.BASIS_HUNDERT) * TAX_CONSTANTS.BASIS_HUNDERT;
    const steuerbemessungsgrundlage = gerundeterErtrag - TAX_CONSTANTS.FREIBETRAG_KST_GEWST;

    const hebesatzFaktor = hebesatz / TAX_CONSTANTS.BASIS_HUNDERT;

    const gewerbesteuer = steuerbemessungsgrundlage * TAX_CONSTANTS.GEWST_MESSZAHL * hebesatzFaktor;

    return gewerbesteuer > 0 ? gewerbesteuer : 0;
  }

  /**
   * Calculates VAT (Umsatzsteuer) liability, input tax (Vorsteuer), and the final payment (Zahllast).
   * @param daten - The complete tax data for the club.
   * @returns An object with umsatzsteuer, vorsteuer, and zahllast.
   */
  public berechneUmsatzsteuer(daten: VereinSteuererklaerung): { umsatzsteuer: number; vorsteuer: number; zahllast: number } {
    if (daten.verein.kleinunternehmerregelung) {
      return { umsatzsteuer: 0, vorsteuer: 0, zahllast: 0 };
    }

    const steuerpflichtigeBereiche = [
      daten.zweckbetrieb,
      daten.wirtschaftlicherGeschaeftsbetrieb,
    ];

    let umsatzsteuer = 0;
    let vorsteuer = 0;

    for (const bereich of steuerpflichtigeBereiche) {
      for (const einnahme of bereich.einnahmen) {
        if (einnahme.ustSatz) {
          umsatzsteuer += einnahme.betrag * (einnahme.ustSatz / TAX_CONSTANTS.BASIS_HUNDERT);
        }
      }
      for (const ausgabe of bereich.ausgaben) {
        if (ausgabe.ustSatz) {
          vorsteuer += ausgabe.betrag * (ausgabe.ustSatz / TAX_CONSTANTS.BASIS_HUNDERT);
        }
      }
    }

    const zahllast = umsatzsteuer - vorsteuer;

    return { umsatzsteuer, vorsteuer, zahllast };
  }

  /**
   * Offsets the results of the four spheres against each other according to tax law.
   * @param ergebnisse - An object with the results of the four spheres.
   * @returns A new object with the final, offset results per sphere.
   */
  private verrechneErgebnisse(ergebnisse: Record<string, number>): Record<string, number> {
    const verrechneteErgebnisse = { ...ergebnisse };

    // Verlustverrechnung aus Zweckbetrieb und Vermögensverwaltung mit WGB
    const verlustquellen = ['zweckbetrieb', 'vermoegensverwaltung'];
    for (const quelle of verlustquellen) {
      if (verrechneteErgebnisse[quelle] < 0 && verrechneteErgebnisse.wirtschaftlicherGeschaeftsbetrieb > 0) {
        const verrechnungsbetrag = Math.min(
          Math.abs(verrechneteErgebnisse[quelle]),
          verrechneteErgebnisse.wirtschaftlicherGeschaeftsbetrieb
        );
        verrechneteErgebnisse[quelle] += verrechnungsbetrag;
        verrechneteErgebnisse.wirtschaftlicherGeschaeftsbetrieb -= verrechnungsbetrag;
      }
    }

    // Verlustverrechnung aus WGB mit Zweckbetrieb
    if (verrechneteErgebnisse.wirtschaftlicherGeschaeftsbetrieb < 0 && verrechneteErgebnisse.zweckbetrieb > 0) {
      const verrechnungsbetrag = Math.min(
        Math.abs(verrechneteErgebnisse.wirtschaftlicherGeschaeftsbetrieb),
        verrechneteErgebnisse.zweckbetrieb
      );
      verrechneteErgebnisse.wirtschaftlicherGeschaeftsbetrieb += verrechnungsbetrag;
      verrechneteErgebnisse.zweckbetrieb -= verrechnungsbetrag;
    }

    return verrechneteErgebnisse;
  }

  /**
   * Main calculation function that orchestrates all steps.
   * @param daten - The complete tax data for the club.
   * @returns A comprehensive result object.
   */
  public berechneGesamtergebnis(daten: VereinSteuererklaerung) {
    const einzelErgebnisse = {
      ideellerBereich: this.berechneErgebnisFuerBereich(daten.ideellerBereich),
      vermoegensverwaltung: this.berechneErgebnisFuerBereich(daten.vermoegensverwaltung),
      zweckbetrieb: this.berechneErgebnisFuerBereich(daten.zweckbetrieb),
      wirtschaftlicherGeschaeftsbetrieb: this.berechneErgebnisWirtschaftlicherGeschaeftsbetrieb(daten.wirtschaftlicherGeschaeftsbetrieb),
    };

    const finaleErgebnisse = this.verrechneErgebnisse(einzelErgebnisse);
    
    return {
      einzelErgebnisse,
      finaleErgebnisse,
    };
  }
}