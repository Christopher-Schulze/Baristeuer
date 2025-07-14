import { createTRPCRouter } from '@/server/trpc';
import { steuerRouter } from './steuer';

export const appRouter = createTRPCRouter({
  steuer: steuerRouter,
});

export type AppRouter = typeof appRouter;