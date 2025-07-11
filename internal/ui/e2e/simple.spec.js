import { test, expect } from '@playwright/test'

test.beforeEach(async ({ page }) => {
  await page.addInitScript(() => {
    window.backend = {
      Generator: { GenerateReport: async () => 'test.pdf' },
    }
  })
})

test('basic interactions', async ({ page }) => {
  await page.goto('/')

  // add income
  await page.getByPlaceholder('Quelle').fill('Demo income')
  await page.getByPlaceholder('Betrag').first().fill('10')
  await page.getByRole('button', { name: 'Hinzufügen' }).first().click()
  await expect(page.getByText('Demo income')).toBeVisible()

  // add expense
  await page.getByPlaceholder('Beschreibung').fill('Demo expense')
  await page.getByPlaceholder('Betrag').nth(1).fill('5')
  await page.getByRole('button', { name: 'Hinzufügen' }).nth(1).click()
  await expect(page.getByText('Demo expense')).toBeVisible()

  // open PDF preview
  await page.getByRole('button', { name: 'PDF Vorschau' }).click()
  await expect(page.getByRole('button', { name: 'PDF erzeugen' })).toBeVisible()

  await page.getByRole('button', { name: 'PDF erzeugen' }).click()
  await expect(page.getByText('Download')).toBeVisible()
})
