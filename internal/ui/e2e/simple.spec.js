import { test, expect } from '@playwright/test'

test('add income', async ({ page }) => {
  await page.goto('/')
  await page.getByPlaceholder('Quelle').fill('Demo')
  await page.getByPlaceholder('Betrag').fill('10')
  await page.getByRole('button', { name: 'Hinzufügen' }).click()
  await expect(page.getByText('Demo')).toBeVisible()
})
