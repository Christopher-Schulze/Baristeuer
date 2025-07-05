import { test, expect } from "@playwright/test";
import { fileURLToPath } from "url";
import path from "path";

test.beforeEach(async ({ page }) => {
  const dir = path.dirname(fileURLToPath(import.meta.url));
  const scriptPath = path.join(dir, "mockDataService.js");
  await page.addInitScript({ path: scriptPath });
});

test("create project and manage income and expenses", async ({ page }) => {
  await page.goto("/");

  await page.getByRole("textbox", { name: /Neues Projekt/i }).fill("Demo");
  await page.getByRole("button", { name: /Erstellen/i }).click();
  await expect(page.getByText("Demo")).toBeVisible();

  await page.getByRole("tab", { name: /Einnahmen/i }).click();
  await page.getByLabel("Quelle").fill("Spende");
  await page.getByLabel("Betrag").fill("10");
  await page.getByRole("button", { name: /Hinzufügen/i }).click();
  await expect(page.getByText("Spende")).toBeVisible();
  await page
    .getByRole("button", { name: /Löschen/i })
    .first()
    .click();
  await expect(page.getByText("Spende")).not.toBeVisible();

  await page.getByRole("tab", { name: /Ausgaben/i }).click();
  await page.getByLabel("Beschreibung").fill("Miete");
  await page.getByLabel(/^Betrag/).fill("5");
  await page.getByRole("button", { name: /Hinzufügen/i }).click();
  await expect(page.getByText("Miete")).toBeVisible();
  await page
    .getByRole("button", { name: /Löschen/i })
    .first()
    .click();
  await expect(page.getByText("Miete")).not.toBeVisible();
});
