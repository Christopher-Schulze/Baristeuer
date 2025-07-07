import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import { configDefaults } from 'vitest/config'

// Vitest runs in a Node environment which would normally resolve the server
// build of Svelte. Alias to the client build so component tests work as
// expected.
const resolve = {
  conditions: ['browser'],
}

export default defineConfig({
  plugins: [svelte()],
  resolve,
  test: {
    environment: 'jsdom',
    globals: true,
    exclude: [...configDefaults.exclude, 'e2e/**'],
    deps: {
      inline: ['svelte', '@sveltejs/vite-plugin-svelte']
    },
  },
})
