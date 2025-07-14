import { PDFDocument, rgb, StandardFonts } from 'pdf-lib';
import { VereinSteuererklaerung, Steuerbereich, Eintrag } from '../../validations/steuerdaten';
import { SteuerEngine } from '../steuer/engine';
import { PDFPage, PDFFont, cmyk } from 'pdf-lib';

const FONT_SIZE_H1 = 22;
const FONT_SIZE_H2 = 18;
const FONT_SIZE_H3 = 14;
const FONT_SIZE_TEXT = 10;
const PADDING = 50;
const LINE_HEIGHT = 18;
const PAGE_WIDTH = 595.28;
const PAGE_HEIGHT = 841.89;

export class PdfGenerator {
    private doc!: PDFDocument;
    private page!: PDFPage;
    private yPosition = 0;
    private font!: PDFFont;
    private boldFont!: PDFFont;
    private steuerEngine: SteuerEngine;

    constructor(private data: VereinSteuererklaerung) {
        this.steuerEngine = new SteuerEngine();
    }

    public async generate(): Promise<Uint8Array> {
        this.doc = await PDFDocument.create();
        this.font = await this.doc.embedFont(StandardFonts.Helvetica);
        this.boldFont = await this.doc.embedFont(StandardFonts.HelveticaBold);

        this.drawCoverPage();
        this.drawDetailedPages();
        this.addPageNumbers();

        return this.doc.save();
    }

    private formatCurrency(value: number | undefined | null): string {
        if (typeof value !== 'number') {
            return 'N/A';
        }
        // Directly return the correctly formatted string from Intl.NumberFormat
        return new Intl.NumberFormat('de-DE', {
            style: 'currency',
            currency: 'EUR',
            minimumFractionDigits: 2,
            maximumFractionDigits: 2,
        }).format(value);
    }
    
    private addNewPage() {
        this.page = this.doc.addPage();
        this.yPosition = PAGE_HEIGHT - PADDING;
    }

    private drawCoverPage() {
        this.addNewPage();
        
        // Title
        this.page.drawText(`Steuererklärung ${this.data.jahr} für ${this.data.verein.name}`, {
            x: PADDING,
            y: this.yPosition,
            font: this.boldFont,
            size: FONT_SIZE_H1,
            color: rgb(0, 0, 0),
        });
        this.yPosition -= LINE_HEIGHT * 3;

        // Club Data Box
        this.drawBox('Stammdaten des Vereins', [
            `Name: ${this.data.verein.name}`,
            `Adresse: ${this.data.verein.adresse.strasse}, ${this.data.verein.adresse.plz} ${this.data.verein.adresse.ort}`,
            `Steuernummer: ${this.data.verein.steuernummer || 'Nicht angegeben'}`,
            `Zuständiges Finanzamt: ${this.data.verein.finanzamtAdresse || 'Nicht angegeben'}`,
            `Erstellt am: ${new Intl.DateTimeFormat('de-DE', { dateStyle: 'long', timeStyle: 'short' }).format(new Date())}`
        ]);
        
        // Tax Summary Box
        const results = this.steuerEngine.berechneGesamtergebnis(this.data);
        const kst = this.steuerEngine.berechneKoerperschaftsteuer(results.finaleErgebnisse.wirtschaftlicherGeschaeftsbetrieb);
        const gewerbesteuer = this.steuerEngine.berechneGewerbesteuer(results.finaleErgebnisse.wirtschaftlicherGeschaeftsbetrieb, this.data.verein.hebesatz);
        const ust = this.steuerEngine.berechneUmsatzsteuer(this.data);

        this.drawBox('Zusammenfassung der Steuerlast (KSt & GewSt)', [
            `Körperschaftsteuer: ${this.formatCurrency(kst.koerperschaftsteuer)}`,
            `Solidaritätszuschlag: ${this.formatCurrency(kst.solidaritätszuschlag)}`,
            `Gewerbesteuer: ${this.formatCurrency(gewerbesteuer)}`,
        ]);

        // Detailed VAT Summary Box as requested by the audit
        this.drawBox('Umsatzsteuer-Detailberechnung', [
            `Umsatzsteuer (vereinnahmt): ${this.formatCurrency(ust.umsatzsteuer)}`,
            `Vorsteuer (abziehbar): ${this.formatCurrency(ust.vorsteuer)}`,
            `--------------------------------------------------`,
            `Zahllast / Erstattungsanspruch: ${this.formatCurrency(ust.zahllast)}`,
        ]);
    }

    private drawBox(title: string, lines: string[]) {
        const boxHeight = (lines.length + 2) * LINE_HEIGHT;
        this.yPosition -= LINE_HEIGHT;
        
        this.page.drawRectangle({
            x: PADDING,
            y: this.yPosition - boxHeight + LINE_HEIGHT,
            width: PAGE_WIDTH - 2 * PADDING,
            height: boxHeight,
            borderColor: rgb(0.7, 0.7, 0.7),
            borderWidth: 1,
        });

        this.page.drawText(title, {
            x: PADDING + 10,
            y: this.yPosition,
            font: this.boldFont,
            size: FONT_SIZE_H3,
        });
        this.yPosition -= LINE_HEIGHT * 1.5;

        lines.forEach(line => {
            this.page.drawText(line, {
                x: PADDING + 10,
                y: this.yPosition,
                font: this.font,
                size: FONT_SIZE_TEXT,
            });
            this.yPosition -= LINE_HEIGHT;
        });
        this.yPosition -= LINE_HEIGHT * 2;
    }

    private drawDetailedPages() {
        this.addNewPage();
        this.drawSection('Ideeller Bereich', this.data.ideellerBereich);
        this.drawSection('Vermögensverwaltung', this.data.vermoegensverwaltung);
        this.drawSection('Zweckbetrieb', this.data.zweckbetrieb);
        this.drawSection('Wirtschaftlicher Geschäftsbetrieb', this.data.wirtschaftlicherGeschaeftsbetrieb);
    }
    
    private drawSection(title: string, sectionData: Steuerbereich) {
        const sectionHeight = (sectionData.einnahmen.length + sectionData.ausgaben.length + 8) * LINE_HEIGHT;
        if (this.yPosition < sectionHeight) {
            this.addNewPage();
        }

        // Section Title Background
        this.page.drawRectangle({
            x: PADDING,
            y: this.yPosition - LINE_HEIGHT,
            width: PAGE_WIDTH - 2 * PADDING,
            height: LINE_HEIGHT * 1.5,
            color: cmyk(0.1, 0.05, 0, 0), // Light grey background
        });

        this.page.drawText(title, {
            x: PADDING + 10,
            y: this.yPosition - (LINE_HEIGHT / 2),
            font: this.boldFont,
            size: FONT_SIZE_H2,
        });
        this.yPosition -= LINE_HEIGHT * 2;

        this.drawTable('Einnahmen', sectionData.einnahmen);
        this.drawTable('Ausgaben', sectionData.ausgaben);
        
        const einnahmenTotal = this.steuerEngine.summiereBetraege(sectionData.einnahmen);
        const ausgabenTotal = this.steuerEngine.summiereBetraege(sectionData.ausgaben);
        const ergebnis = einnahmenTotal - ausgabenTotal;

        this.yPosition -= LINE_HEIGHT / 2;
        this.page.drawLine({
             start: { x: PADDING, y: this.yPosition },
             end: { x: PAGE_WIDTH - PADDING, y: this.yPosition },
             thickness: 1,
             color: rgb(0.8, 0.8, 0.8)
        });
        this.yPosition -= LINE_HEIGHT;

        this.page.drawText(`Ergebnis der Sphäre:`, {
            x: PADDING + 10,
            y: this.yPosition,
            font: this.boldFont,
            size: FONT_SIZE_TEXT,
        });
        this.page.drawText(this.formatCurrency(ergebnis), {
            x: 450,
            y: this.yPosition,
            font: this.boldFont,
            size: FONT_SIZE_TEXT,
        });
        this.yPosition -= LINE_HEIGHT * 3;
    }

    private drawTable(title: string, items: Eintrag[]) {
        if (this.yPosition < PADDING * 3) {
            this.addNewPage();
        }

        this.page.drawText(title, {
            x: PADDING + 10,
            y: this.yPosition,
            font: this.boldFont,
            size: FONT_SIZE_H3,
        });
        this.yPosition -= LINE_HEIGHT * 1.5;

        if (items.length === 0) {
            this.page.drawText('Keine Einträge vorhanden.', {
                x: PADDING + 20,
                y: this.yPosition,
                font: this.font,
                size: FONT_SIZE_TEXT,
                color: rgb(0.5, 0.5, 0.5),
            });
            this.yPosition -= LINE_HEIGHT * 1.5;
            return;
        }

        items.forEach(item => {
             if (this.yPosition < PADDING) {
                this.addNewPage();
            }
            this.page.drawText(item.bezeichnung, {
                x: PADDING + 20,
                y: this.yPosition,
                font: this.font,
                size: FONT_SIZE_TEXT,
            });
            this.page.drawText(this.formatCurrency(item.betrag), {
                x: 450,
                y: this.yPosition,
                font: this.font,
                size: FONT_SIZE_TEXT,
            });
            this.yPosition -= LINE_HEIGHT;
        });
        this.yPosition -= LINE_HEIGHT;
    }
    
    private addPageNumbers() {
        const pages = this.doc.getPages();
        for (let i = 0; i < pages.length; i++) {
            const page = pages[i];
            page.drawText(`Seite ${i + 1} von ${pages.length}`, {
                x: PADDING,
                y: PADDING / 2,
                font: this.font,
                size: FONT_SIZE_TEXT - 2,
                color: rgb(0.5, 0.5, 0.5),
            });
        }
    }
}