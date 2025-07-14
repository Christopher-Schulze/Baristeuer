import { createTRPCRouter, publicProcedure } from '@/server/trpc';
import { z } from 'zod';
import { VereinSteuererklaerungSchema } from '@/validations/steuerdaten';
import { PdfGenerator } from '@/lib/pdf/pdfGenerator';

export const steuerRouter = createTRPCRouter({
  get: publicProcedure
    .input(
      z.object({
        vereinId: z.string(),
        jahr: z.number(),
      })
    )
    .query(async ({ ctx, input }) => {
      return ctx.prisma.steuererklaerung.findUnique({
        where: {
          vereinId_jahr: {
            vereinId: input.vereinId,
            jahr: input.jahr,
          },
        },
      });
    }),

  upsert: publicProcedure
    .input(
      VereinSteuererklaerungSchema.extend({
        vereinId: z.string(),
      })
    )
    .mutation(async ({ ctx, input }) => {
      const { vereinId, jahr, verein, ...steuerData } = input;
      
      // Update Verein data first
      await ctx.prisma.verein.update({
        where: { id: vereinId },
        data: {
          name: verein.name,
          adresse: verein.adresse,
          steuernummer: verein.steuernummer,
          finanzamtAdresse: verein.finanzamtAdresse,
          hebesatz: verein.hebesatz,
          kleinunternehmerregelung: verein.kleinunternehmerregelung,
        },
      });
      
      // Then upsert the Steuererklärung
      return ctx.prisma.steuererklaerung.upsert({
        where: {
          vereinId_jahr: {
            vereinId: vereinId,
            jahr: jahr,
          },
        },
        update: {
          ideellerBereich: steuerData.ideellerBereich,
          vermoegensverwaltung: steuerData.vermoegensverwaltung,
          zweckbetrieb: steuerData.zweckbetrieb,
          wirtschaftlicherGeschaeftsbetrieb: steuerData.wirtschaftlicherGeschaeftsbetrieb,
        },
        create: {
          jahr: jahr,
          vereinId: vereinId,
          ideellerBereich: steuerData.ideellerBereich,
          vermoegensverwaltung: steuerData.vermoegensverwaltung,
          zweckbetrieb: steuerData.zweckbetrieb,
          wirtschaftlicherGeschaeftsbetrieb: steuerData.wirtschaftlicherGeschaeftsbetrieb,
        },
      });
    }),

  generatePdf: publicProcedure
    .input(VereinSteuererklaerungSchema)
    .mutation(async ({ input }) => {
      const generator = new PdfGenerator(input);
      const pdfBytes = await generator.generate();
      return Buffer.from(pdfBytes).toString('base64');
    }),
});