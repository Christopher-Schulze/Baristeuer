import { test, expect } from "@playwright/test";
import { fileURLToPath } from "url";
import path from "path";

test.beforeEach(async ({ page }) => {
  const dir = path.dirname(fileURLToPath(import.meta.url));
  await page.addInitScript({ path: path.join(dir, "mockDataService.js") });
  await page.addInitScript({ path: path.join(dir, "mockPDFGenerator.js") });
});

test("generate forms", async ({ page }) => {
  await page.goto("/");

  await page.getByRole("textbox", { name: /Neues Projekt/i }).fill("Demo");
  await page.getByRole("button", { name: /Erstellen/i }).click();

  await page.getByRole("tab", { name: /Formulare/i }).click();
  await page.getByRole("button", { name: /Alle Formulare erstellen/i }).click();

  const calls = await page.evaluate(() => window.__pdfCalls);
  expect(calls.map((c) => c.name)).toContain("GenerateAllForms");
});
