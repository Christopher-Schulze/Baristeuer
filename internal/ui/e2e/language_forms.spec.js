import { test, expect } from "@playwright/test";
import { fileURLToPath } from "url";
import path from "path";

test.beforeEach(async ({ page }) => {
  const dir = path.dirname(fileURLToPath(import.meta.url));
  await page.addInitScript({ path: path.join(dir, "mockDataService.js") });
  await page.addInitScript({ path: path.join(dir, "mockPDFGenerator.js") });
});

test("switch languages and generate forms", async ({ page }) => {
  await page.goto("/");

  await page.getByRole("textbox", { name: /Neues Projekt/i }).fill("Demo");
  await page.getByRole("button", { name: /Erstellen/i }).click();

  // switch to English and generate
  await page.getByRole("button", { name: "DE" }).click();
  await page.getByRole("option", { name: "EN" }).click();
  await page.getByRole("tab", { name: "Forms" }).click();
  await page.getByRole("button", { name: /Generate All Forms/i }).click();

  // switch back to German and generate again
  await page.getByRole("button", { name: "EN" }).click();
  await page.getByRole("option", { name: "DE" }).click();
  await page.getByRole("tab", { name: /Formulare/i }).click();
  await page.getByRole("button", { name: /Alle Formulare erstellen/i }).click();

  const calls = await page.evaluate(() => window.__pdfCalls);
  const count = calls.filter((c) => c.name === "GenerateAllForms").length;
  expect(count).toBe(2);
});
