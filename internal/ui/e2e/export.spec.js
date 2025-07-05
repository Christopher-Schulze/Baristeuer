import { test, expect } from "@playwright/test";
import { fileURLToPath } from "url";
import path from "path";

test.beforeEach(async ({ page }) => {
  const dir = path.dirname(fileURLToPath(import.meta.url));
  await page.addInitScript({ path: path.join(dir, "mockDataService.js") });
});

test("export data", async ({ page }) => {
  await page.goto("/");

  await page.getByRole("textbox", { name: /Neues Projekt/i }).fill("Demo" );
  await page.getByRole("button", { name: /Erstellen/i }).click();

  await page.getByRole("tab", { name: /Einstellungen/i }).click();
  await page.getByLabel(/Datenbank exportieren/i).fill("db.sqlite");
  await page.getByRole("button", { name: /Datenbank exportieren/i }).click();
  await page.getByLabel(/Projekt-CSV exportieren/i).fill("proj.csv");
  await page.getByRole("button", { name: /Projekt-CSV exportieren/i }).click();

  const exports = await page.evaluate(() => window.__exports);
  expect(exports).toEqual([
    { type: "db", dest: "db.sqlite" },
    { type: "csv", pid: 1, dest: "proj.csv" },
  ]);
});
