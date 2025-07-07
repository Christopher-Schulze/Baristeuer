import { test, expect } from '@playwright/test'

test('basic interactions', async ({ page }) => {
  await page.goto('/')

  // add income
  await page.getByPlaceholder('Source').fill('Demo income')
  await page.getByPlaceholder('Amount').first().fill('10')
  await page.getByRole('button', { name: 'Add' }).first().click()
  await expect(page.getByText('Demo income')).toBeVisible()

  // add expense
  await page.getByPlaceholder('Description').fill('Demo expense')
  await page.getByPlaceholder('Amount').nth(1).fill('5')
  await page.getByRole('button', { name: 'Add' }).nth(1).click()
  await expect(page.getByText('Demo expense')).toBeVisible()

  // open PDF preview
  await page.getByRole('button', { name: 'PDF Preview' }).click()
  await expect(page.getByTitle('PDF Preview')).toBeVisible()
})
