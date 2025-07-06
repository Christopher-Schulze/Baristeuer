import { test, expect } from "@playwright/test";
import { fileURLToPath } from "url";
import path from "path";

test.beforeEach(async ({ page }) => {
  const dir = path.dirname(fileURLToPath(import.meta.url));
  await page.addInitScript({ path: path.join(dir, "mockDataService.js") });
  await page.addInitScript({ path: path.join(dir, "mockFailCSV.js") });
});

test("show error if CSV export fails", async ({ page }) => {
  await page.goto("/");

  await page.getByRole("textbox", { name: /Neues Projekt/i }).fill("Demo");
  await page.getByRole("button", { name: /Erstellen/i }).click();

  await page.getByRole("tab", { name: /Einstellungen/i }).click();
  await page.getByLabel(/Projekt-CSV exportieren/i).fill("fail.csv");
  await page.getByRole("button", { name: /Projekt-CSV exportieren/i }).click();

  await expect(page.getByText("CSV export failed")).toBeVisible();
});
