import { describe, it, expect } from 'vitest';
import { SteuerEngine } from './engine';
import {
  VereinSteuererklaerung,
  Steuerbereich,
  Eintrag as SteuerEintrag,
} from '@/validations/steuerdaten';

describe('SteuerEngine', () => {
  describe('Berechnungsgrundlagen', () => {
    it('sollte die Summe von Einnahmen oder Ausgaben korrekt berechnen', () => {
      const engine = new SteuerEngine();
      const betraege: SteuerEintrag[] = [
        { bezeichnung: 'Posten 1', betrag: 100.5 },
        { bezeichnung: 'Posten 2', betrag: 50.25 },
        { bezeichnung: 'Posten 3', betrag: 200 },
      ];
      const summe = engine.summiereBetraege(betraege);
      expect(summe).toBe(350.75);
    });

    it('sollte das Ergebnis (Überschuss/Fehlbetrag) eines Bereichs korrekt berechnen', () => {
      const engine = new SteuerEngine();
      const bereich: Steuerbereich = {
        einnahmen: [
          { bezeichnung: 'Einnahme 1', betrag: 1000 },
          { bezeichnung: 'Einnahme 2', betrag: 500 },
        ],
        ausgaben: [
          { bezeichnung: 'Ausgabe 1', betrag: 300 },
          { bezeichnung: 'Ausgabe 2', betrag: 200 },
        ],
        angewandtePauschalen: 0,
      };
      const ergebnis = engine.berechneErgebnisFuerBereich(bereich);
      expect(ergebnis).toBe(1000); // 1500 - 500
    });
  });

  describe('Wirtschaftlicher Geschäftsbetrieb', () => {
    it('sollte das tatsächliche Ergebnis (Gewinn) berechnen, auch wenn die Einnahmen unter der Freigrenze liegen', () => {
      const engine = new SteuerEngine();
      const bereich: Steuerbereich = {
        einnahmen: [{ bezeichnung: 'Verkäufe', betrag: 40000, }],
        ausgaben: [{ bezeichnung: 'Wareneinsatz', betrag: 10000 }],
        angewandtePauschalen: 0
      };
      const ergebnis = engine.berechneErgebnisWirtschaftlicherGeschaeftsbetrieb(bereich);
      // The engine now correctly returns the actual result. The tax service is responsible for applying the limit.
      expect(ergebnis).toBe(30000);
    });

    it('sollte das tatsächliche Ergebnis (Verlust) korrekt ausweisen', () => {
      const engine = new SteuerEngine();
      const bereich: Steuerbereich = {
        einnahmen: [{ bezeichnung: 'Verkäufe', betrag: 10000 }],
        ausgaben: [{ bezeichnung: 'Wareneinsatz', betrag: 15000 }],
        angewandtePauschalen: 0
      };
      const ergebnis = engine.berechneErgebnisWirtschaftlicherGeschaeftsbetrieb(bereich);
      expect(ergebnis).toBe(-5000);
    });

    it('sollte das korrekte Ergebnis ausweisen, wenn die Einnahmen über der Freigrenze von 45.000€ liegen', () => {
      const engine = new SteuerEngine();
      const bereich: Steuerbereich = {
        einnahmen: [{ bezeichnung: 'Verkäufe', betrag: 50000 }],
        ausgaben: [{ bezeichnung: 'Wareneinsatz', betrag: 10000 }],
        angewandtePauschalen: 0
      };
      const ergebnis = engine.berechneErgebnisWirtschaftlicherGeschaeftsbetrieb(bereich);
      expect(ergebnis).toBe(40000);
    });
  });

  describe('Körperschaftsteuer-Berechnung', () => {
    it('sollte 0€ KSt und SoliZ zurückgeben, wenn das Einkommen den Freibetrag von 5.000€ nicht übersteigt', () => {
      const engine = new SteuerEngine();
      const { koerperschaftsteuer, solidaritätszuschlag } = engine.berechneKoerperschaftsteuer(4999.99);
      expect(koerperschaftsteuer).toBe(0);
      expect(solidaritätszuschlag).toBe(0);
    });

    it('sollte die korrekte KSt und den SoliZ berechnen, wenn das Einkommen über 5.000€ liegt', () => {
      const engine = new SteuerEngine();
      const { koerperschaftsteuer, solidaritätszuschlag } = engine.berechneKoerperschaftsteuer(10000);
      // (10000 - 5000) * 0.15 = 750
      expect(koerperschaftsteuer).toBe(750);
      // 750 * 0.055 = 41.25
      expect(solidaritätszuschlag).toBe(41.25);
    });

    it('sollte 0€ KSt und SoliZ zurückgeben, wenn das Einkommen genau 5.000€ beträgt', () => {
      const engine = new SteuerEngine();
      const { koerperschaftsteuer, solidaritätszuschlag } = engine.berechneKoerperschaftsteuer(5000);
      expect(koerperschaftsteuer).toBe(0);
      expect(solidaritätszuschlag).toBe(0);
    });

    it('sollte bei negativem Einkommen (Verlust) 0€ KSt und SoliZ zurückgeben', () => {
      const engine = new SteuerEngine();
      const { koerperschaftsteuer, solidaritätszuschlag } = engine.berechneKoerperschaftsteuer(-1000);
      expect(koerperschaftsteuer).toBe(0);
      expect(solidaritätszuschlag).toBe(0);
    });
  });
describe('Gewerbesteuer-Berechnung', () => {
    it('sollte 0€ GewSt zurückgeben, wenn das Einkommen den Freibetrag von 5.000€ nicht übersteigt', () => {
      const engine = new SteuerEngine();
      const gewerbesteuer = engine.berechneGewerbesteuer(5000, 400);
      expect(gewerbesteuer).toBe(0);
    });

    it('sollte die korrekte GewSt für ein Einkommen über 5.000€ berechnen', () => {
      const engine = new SteuerEngine();
      const gewerbesteuer = engine.berechneGewerbesteuer(15050, 400);
      // (15000 - 5000) * 0.035 * 4 = 1400
      expect(gewerbesteuer).toBe(1400);
    });

    it('sollte den Gewerbeertrag bei .99 Werten korrekt auf volle 100€ abrunden (§ 9 Nr. 1 GewStG)', () => {
      const engine = new SteuerEngine();
      const gewerbesteuer = engine.berechneGewerbesteuer(15099.99, 400);
      // Gewerbeertrag wird auf volle 100 abgerundet: 15099.99 -> 15000
      // (15000 - 5000) * 0.035 * 4 = 1400
      expect(gewerbesteuer).toBe(1400);
    });

    it('sollte bei negativem Einkommen (Verlust) 0€ GewSt zurückgeben', () => {
      const engine = new SteuerEngine();
      const gewerbesteuer = engine.berechneGewerbesteuer(-10000, 400);
      expect(gewerbesteuer).toBe(0);
    });
  });

  describe('Umsatzsteuer-Berechnung', () => {
   it('sollte 0€ USt-Schuld zurückgeben, wenn die Kleinunternehmerregelung aktiv ist', () => {
     const engine = new SteuerEngine();
      const daten: VereinSteuererklaerung = {
        jahr: 2023,
        verein: { name: 'Test', adresse: { strasse: 'Teststraße 1', plz: '12345', ort: 'Teststadt' }, finanzamtAdresse: 'Finanzamt Test', kleinunternehmerregelung: true, hebesatz: 400 },
        ideellerBereich: { einnahmen: [], ausgaben: [], angewandtePauschalen: 0 },
        vermoegensverwaltung: { einnahmen: [], ausgaben: [], angewandtePauschalen: 0 },
        zweckbetrieb: {
          einnahmen: [{ bezeichnung: 'E1', betrag: 1000, ustSatz: 7 }],
          ausgaben: [],
          angewandtePauschalen: 0
        },
        wirtschaftlicherGeschaeftsbetrieb: {
          einnahmen: [{ bezeichnung: 'E2', betrag: 2000, ustSatz: 19 }],
          ausgaben: [],
          angewandtePauschalen: 0
        },
      };
     const { umsatzsteuer, vorsteuer, zahllast } = engine.berechneUmsatzsteuer(daten);
     expect(umsatzsteuer).toBe(0);
     expect(vorsteuer).toBe(0);
     expect(zahllast).toBe(0);
   });

   it('sollte die USt-Last korrekt berechnen, wenn die Kleinunternehmerregelung inaktiv ist', () => {
       const engine = new SteuerEngine();
        const daten: VereinSteuererklaerung = {
          jahr: 2023,
          verein: { name: 'Test', adresse: { strasse: 'Teststraße 1', plz: '12345', ort: 'Teststadt' }, finanzamtAdresse: 'Finanzamt Test', kleinunternehmerregelung: false, hebesatz: 400 },
          ideellerBereich: {
            einnahmen: [{ bezeichnung: 'Spende', betrag: 500 }], // No USt
            ausgaben: [],
            angewandtePauschalen: 0
          },
          vermoegensverwaltung: { einnahmen: [], ausgaben: [], angewandtePauschalen: 0 },
          zweckbetrieb: {
            einnahmen: [{ bezeichnung: 'E1', betrag: 1000, ustSatz: 7 }],
            ausgaben: [],
            angewandtePauschalen: 0
          },
          wirtschaftlicherGeschaeftsbetrieb: {
            einnahmen: [{ bezeichnung: 'E2', betrag: 2000, ustSatz: 19 }],
            ausgaben: [],
            angewandtePauschalen: 0
          },
        };
       const { umsatzsteuer } = engine.berechneUmsatzsteuer(daten);
       expect(umsatzsteuer).toBe(450); // (1000 * 0.07) + (2000 * 0.19) = 70 + 380
   });

   it('sollte den Vorsteuerabzug korrekt berechnen', () => {
       const engine = new SteuerEngine();
        const daten: VereinSteuererklaerung = {
          jahr: 2023,
          verein: { name: 'Test', adresse: { strasse: 'Teststraße 1', plz: '12345', ort: 'Teststadt' }, finanzamtAdresse: 'Finanzamt Test', kleinunternehmerregelung: false, hebesatz: 400 },
          ideellerBereich: { einnahmen: [], ausgaben: [], angewandtePauschalen: 0 },
          vermoegensverwaltung: { einnahmen: [], ausgaben: [], angewandtePauschalen: 0 },
           zweckbetrieb: {
            einnahmen: [],
            ausgaben: [{ bezeichnung: 'A1', betrag: 500, ustSatz: 19 }],
            angewandtePauschalen: 0
          },
          wirtschaftlicherGeschaeftsbetrieb: {
            einnahmen: [],
            ausgaben: [{ bezeichnung: 'A2', betrag: 100, ustSatz: 7 }], // 7% Vorsteuer
            angewandtePauschalen: 0
          },
        };
       const { vorsteuer } = engine.berechneUmsatzsteuer(daten);
       expect(vorsteuer).toBe(102); // (500 * 0.19) + (100 * 0.07) = 95 + 7
   });
   
   it('sollte die finale USt-Zahllast (Umsatzsteuer - Vorsteuer) korrekt berechnen', () => {
       const engine = new SteuerEngine();
        const daten: VereinSteuererklaerung = {
            jahr: 2023,
            verein: { name: 'Test', adresse: { strasse: 'Teststraße 1', plz: '12345', ort: 'Teststadt' }, finanzamtAdresse: 'Finanzamt Test', kleinunternehmerregelung: false, hebesatz: 400 },
            ideellerBereich: { einnahmen: [], ausgaben: [], angewandtePauschalen: 0 },
            vermoegensverwaltung: { einnahmen: [], ausgaben: [], angewandtePauschalen: 0 },
            zweckbetrieb: {
                einnahmen: [{ bezeichnung: 'E1', betrag: 1000, ustSatz: 7 }],
                ausgaben: [{ bezeichnung: 'A1', betrag: 200, ustSatz: 19 }],
                angewandtePauschalen: 0
            },
            wirtschaftlicherGeschaeftsbetrieb: {
                einnahmen: [{ bezeichnung: 'E2', betrag: 2000, ustSatz: 19 }],
                ausgaben: [{ bezeichnung: 'A2', betrag: 300, ustSatz: 19 }],
                angewandtePauschalen: 0
            },
        };
       const { zahllast } = engine.berechneUmsatzsteuer(daten);
       // USt: (1000 * 0.07) + (2000 * 0.19) = 70 + 380 = 450
       // VSt: (200 * 0.19) + (300 * 0.19) = 38 + 57 = 95
       // Zahllast: 450 - 95 = 355
       expect(zahllast).toBe(355);
   });

   it('sollte sicherstellen, dass nur steuerpflichtige Bereiche in die USt-Berechnung einfließen', () => {
       const engine = new SteuerEngine();
        const daten: VereinSteuererklaerung = {
            jahr: 2023,
            verein: { name: 'Test', adresse: { strasse: 'Teststraße 1', plz: '12345', ort: 'Teststadt' }, finanzamtAdresse: 'Finanzamt Test', kleinunternehmerregelung: false, hebesatz: 400 },
            ideellerBereich: {
                einnahmen: [{ bezeichnung: 'Spende', betrag: 1000, ustSatz: 19 }],
                ausgaben: [{ bezeichnung: 'Dankeskarte', betrag: 50, ustSatz: 19 }],
                angewandtePauschalen: 0
            },
            vermoegensverwaltung: {
                einnahmen: [{ bezeichnung: 'Zinsen', betrag: 200, ustSatz: 19 }],
                ausgaben: [{ bezeichnung: 'Kontoführung', betrag: 10, ustSatz: 19 }],
                angewandtePauschalen: 0
            },
            zweckbetrieb: {
                einnahmen: [],
                ausgaben: [],
                angewandtePauschalen: 0
            },
            wirtschaftlicherGeschaeftsbetrieb: {
                einnahmen: [],
                ausgaben: [],
                angewandtePauschalen: 0
            },
        };
       const { umsatzsteuer, vorsteuer } = engine.berechneUmsatzsteuer(daten);
       expect(umsatzsteuer).toBe(0);
       expect(vorsteuer).toBe(0);
   });
  });


 describe('Pauschalen-Verrechnung', () => {
  it('sollte die angewandten Pauschalen korrekt vom Ergebnis des Bereichs abziehen', () => {
    const engine = new SteuerEngine();
    const bereich: Steuerbereich = {
      einnahmen: [{ bezeichnung: 'Einnahme 1', betrag: 10000 }],
      ausgaben: [{ bezeichnung: 'Ausgabe 1', betrag: 2000 }],
      angewandtePauschalen: 840,
    };
    const ergebnis = engine.berechneErgebnisFuerBereich(bereich);
    expect(ergebnis).toBe(7160); // 10000 - 2000 - 840
  });

  it('sollte auch bei einem resultierenden Verlust die Pauschalen korrekt abziehen', () => {
    const engine = new SteuerEngine();
    const bereich: Steuerbereich = {
      einnahmen: [{ bezeichnung: 'Einnahme 1', betrag: 1000 }],
      ausgaben: [{ bezeichnung: 'Ausgabe 1', betrag: 2000 }],
      angewandtePauschalen: 3000,
    };
    const ergebnis = engine.berechneErgebnisFuerBereich(bereich);
    expect(ergebnis).toBe(-4000); // 1000 - 2000 - 3000
  });

  it('sollte die Pauschalen in mehreren Bereichen korrekt verrechnen', () => {
    const engine = new SteuerEngine();
    const ideellerBereich: Steuerbereich = {
      einnahmen: [{ bezeichnung: 'Einnahme 1', betrag: 5000 }],
      ausgaben: [],
      angewandtePauschalen: 840, // Ehrenamt
    };
    const zweckbetrieb: Steuerbereich = {
      einnahmen: [{ bezeichnung: 'Einnahme 2', betrag: 12000 }],
      ausgaben: [{ bezeichnung: 'Ausgabe 2', betrag: 4000 }],
      angewandtePauschalen: 3000, // Übungsleiter
    };

    const ergebnisIdeell = engine.berechneErgebnisFuerBereich(ideellerBereich);
    const ergebnisZweck = engine.berechneErgebnisFuerBereich(zweckbetrieb);

    expect(ergebnisIdeell).toBe(4160); // 5000 - 840
    expect(ergebnisZweck).toBe(5000); // 12000 - 4000 - 3000
  });
 });

 describe('Ergebnisverrechnung der Sphären (via Gesamtergebnis)', () => {
   const mockSteuerdaten = (ergebnisse: Record<string, number>): VereinSteuererklaerung => {
       const createBereich = (ergebnis: number): Steuerbereich => ({
           einnahmen: [{ bezeichnung: 'E', betrag: ergebnis > 0 ? ergebnis : 0 }],
           ausgaben: [{ bezeichnung: 'A', betrag: ergebnis < 0 ? Math.abs(ergebnis) : 0 }],
           angewandtePauschalen: 0,
       });

       return {
           jahr: 2023,
           verein: { name: 'Test', adresse: { strasse: 'Teststraße 1', plz: '12345', ort: 'Teststadt' }, finanzamtAdresse: 'Finanzamt Test', kleinunternehmerregelung: false, hebesatz: 400 },
           ideellerBereich: createBereich(ergebnisse.ideellerBereich),
           vermoegensverwaltung: createBereich(ergebnisse.vermoegensverwaltung),
           zweckbetrieb: createBereich(ergebnisse.zweckbetrieb),
           wirtschaftlicherGeschaeftsbetrieb: createBereich(ergebnisse.wirtschaftlicherGeschaeftsbetrieb),
       };
   };

  it('sollte einen Verlust im Zweckbetrieb mit einem Gewinn im WGB verrechnen', () => {
    const engine = new SteuerEngine();
    const daten = mockSteuerdaten({
      ideellerBereich: 500,
      vermoegensverwaltung: 1000,
      zweckbetrieb: -2000,
      wirtschaftlicherGeschaeftsbetrieb: 5000,
    });
    const { finaleErgebnisse } = engine.berechneGesamtergebnis(daten);
    expect(finaleErgebnisse.zweckbetrieb).toBe(0);
    expect(finaleErgebnisse.wirtschaftlicherGeschaeftsbetrieb).toBe(3000);
  });

  it('sollte einen Verlust in der Vermögensverwaltung mit einem Gewinn im WGB verrechnen', () => {
      const engine = new SteuerEngine();
      const daten = mockSteuerdaten({
        ideellerBereich: 500,
        vermoegensverwaltung: -1500,
        zweckbetrieb: 1000,
        wirtschaftlicherGeschaeftsbetrieb: 5000,
      });
      const { finaleErgebnisse } = engine.berechneGesamtergebnis(daten);
      expect(finaleErgebnisse.vermoegensverwaltung).toBe(0);
      expect(finaleErgebnisse.wirtschaftlicherGeschaeftsbetrieb).toBe(3500);
  });

  it('sollte einen Verlust im WGB mit einem Gewinn im Zweckbetrieb verrechnen', () => {
      const engine = new SteuerEngine();
      const daten = mockSteuerdaten({
        ideellerBereich: 500,
        vermoegensverwaltung: 1000,
        zweckbetrieb: 4000,
        wirtschaftlicherGeschaeftsbetrieb: -2000,
      });
      const { finaleErgebnisse } = engine.berechneGesamtergebnis(daten);
      expect(finaleErgebnisse.zweckbetrieb).toBe(2000);
      expect(finaleErgebnisse.wirtschaftlicherGeschaeftsbetrieb).toBe(0);
  });

  it('[NEU] sollte Verluste aus Zweckbetrieb und VGV korrekt mit Gewinn im WGB verrechnen', () => {
    const engine = new SteuerEngine();
    const daten = mockSteuerdaten({
      ideellerBereich: 100,
      vermoegensverwaltung: -2000, // Verlust
      zweckbetrieb: -3000,         // Verlust
      wirtschaftlicherGeschaeftsbetrieb: 10000, // Gewinn
    });
    const { finaleErgebnisse } = engine.berechneGesamtergebnis(daten);
    expect(finaleErgebnisse.vermoegensverwaltung).toBe(0); // Vollständig verrechnet
    expect(finaleErgebnisse.zweckbetrieb).toBe(0); // Vollständig verrechnet
    expect(finaleErgebnisse.wirtschaftlicherGeschaeftsbetrieb).toBe(5000); // 10000 - 2000 - 3000
  });

   it('sollte keine Verrechnung durchführen, wenn alle drei verrechenbaren Sphären einen Verlust ausweisen', () => {
     const engine = new SteuerEngine();
     const daten = mockSteuerdaten({
       ideellerBereich: 500,
       vermoegensverwaltung: -1000,
       zweckbetrieb: -2000,
       wirtschaftlicherGeschaeftsbetrieb: -3000,
     });
     const { finaleErgebnisse } = engine.berechneGesamtergebnis(daten);
     expect(finaleErgebnisse.vermoegensverwaltung).toBe(-1000); // Unverändert
     expect(finaleErgebnisse.zweckbetrieb).toBe(-2000); // Unverändert
     expect(finaleErgebnisse.wirtschaftlicherGeschaeftsbetrieb).toBe(-3000); // Unverändert
   });

   it('sollte einen Verlust im Zweckbetrieb nur teilweise mit einem kleineren Gewinn im WGB verrechnen', () => {
     const engine = new SteuerEngine();
     const daten = mockSteuerdaten({
       ideellerBereich: 500,
       vermoegensverwaltung: 1000,
       zweckbetrieb: -5000, // Großer Verlust
       wirtschaftlicherGeschaeftsbetrieb: 2000, // Kleiner Gewinn
     });
     const { finaleErgebnisse } = engine.berechneGesamtergebnis(daten);
     // Der WGB-Gewinn wird auf 0 reduziert
     expect(finaleErgebnisse.wirtschaftlicherGeschaeftsbetrieb).toBe(0);
     // Der Verlust im Zweckbetrieb wird um den Betrag des WGB-Gewinns reduziert
     expect(finaleErgebnisse.zweckbetrieb).toBe(-3000); // -5000 + 2000
   });
  });

  describe('Unzulässige Verlustverrechnung (Negative Testfälle)', () => {
    const mockSteuerdaten = (ergebnisse: Record<string, number>): VereinSteuererklaerung => {
        const createBereich = (ergebnis: number): Steuerbereich => ({
            einnahmen: [{ bezeichnung: 'E', betrag: ergebnis > 0 ? ergebnis : 0 }],
            ausgaben: [{ bezeichnung: 'A', betrag: ergebnis < 0 ? Math.abs(ergebnis) : 0 }],
            angewandtePauschalen: 0,
        });

        return {
            jahr: 2023,
            verein: { name: 'Test', adresse: { strasse: 'Teststraße 1', plz: '12345', ort: 'Teststadt' }, finanzamtAdresse: 'Finanzamt Test', kleinunternehmerregelung: false, hebesatz: 400 },
            ideellerBereich: createBereich(ergebnisse.ideellerBereich),
            vermoegensverwaltung: createBereich(ergebnisse.vermoegensverwaltung),
            zweckbetrieb: createBereich(ergebnisse.zweckbetrieb),
            wirtschaftlicherGeschaeftsbetrieb: createBereich(ergebnisse.wirtschaftlicherGeschaeftsbetrieb),
        };
    };

    it('sollte einen Verlust im ideellen Bereich NICHT mit Gewinnen aus anderen Sphären verrechnen', () => {
      const engine = new SteuerEngine();
      const daten = mockSteuerdaten({
        ideellerBereich: -5000, // Relevanter Verlust
        vermoegensverwaltung: 2000,
        zweckbetrieb: 3000,
        wirtschaftlicherGeschaeftsbetrieb: 4000,
      });
      const { finaleErgebnisse } = engine.berechneGesamtergebnis(daten);
      // Der Verlust im ideellen Bereich muss unverändert bleiben.
      expect(finaleErgebnisse.ideellerBereich).toBe(-5000);
      // Die anderen Bereiche dürfen nicht angetastet werden.
      expect(finaleErgebnisse.vermoegensverwaltung).toBe(2000);
      expect(finaleErgebnisse.zweckbetrieb).toBe(3000);
      expect(finaleErgebnisse.wirtschaftlicherGeschaeftsbetrieb).toBe(4000);
    });

    it('sollte einen Gewinn im ideellen Bereich NICHT zur Deckung von Verlusten aus anderen Sphären verwenden', () => {
      const engine = new SteuerEngine();
      const daten = mockSteuerdaten({
        ideellerBereich: 10000, // Relevanter Gewinn
        vermoegensverwaltung: -2000,
        zweckbetrieb: -3000,
        wirtschaftlicherGeschaeftsbetrieb: -1000,
      });
      const { finaleErgebnisse } = engine.berechneGesamtergebnis(daten);

      // Der Gewinn im ideellen Bereich muss unangetastet bleiben.
      expect(finaleErgebnisse.ideellerBereich).toBe(10000);
      
      // Die Verrechnung zwischen den anderen Sphären soll aber stattfinden.
      // VGV (-2000) und WGB (-1000) können nicht verrechnet werden, da ZB auch negativ ist.
      // Aber der Verlust im WGB (-1000) kann mit dem ZB (-3000) nicht verrechnet werden.
      // Die Verluste bleiben also bestehen, da keine positiven Gegenparts existieren.
      expect(finaleErgebnisse.vermoegensverwaltung).toBe(-2000);
      expect(finaleErgebnisse.zweckbetrieb).toBe(-3000);
      expect(finaleErgebnisse.wirtschaftlicherGeschaeftsbetrieb).toBe(-1000);
    });
  });
});