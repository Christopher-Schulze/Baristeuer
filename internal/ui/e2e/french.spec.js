import { test, expect } from "@playwright/test";
import { fileURLToPath } from "url";
import path from "path";

test.beforeEach(async ({ page }) => {
  const dir = path.dirname(fileURLToPath(import.meta.url));
  const scriptPath = path.join(dir, "mockDataService.js");
  await page.addInitScript({ path: scriptPath });
});

test("switch UI language to French", async ({ page }) => {
  await page.goto("/");

  await page.getByRole("button", { name: "DE" }).click();
  await page.getByRole("option", { name: "FR" }).click();

  await expect(page.getByRole("tab", { name: "Projets" })).toBeVisible();
  await expect(page.getByRole("tab", { name: "Revenus" })).toBeVisible();
  await expect(page.getByRole("tab", { name: "Dépenses" })).toBeVisible();

  await page.getByRole("tab", { name: "Projets" }).click();
  await expect(page.getByLabel("Nouveau projet")).toBeVisible();
  await expect(page.getByRole("button", { name: "Créer" })).toBeVisible();

  await page.getByRole("tab", { name: "Revenus" }).click();
  await expect(page.getByRole("heading", { name: "Nouveau revenu" })).toBeVisible();
  await expect(page.getByLabel("Source")).toBeVisible();
  await expect(page.getByLabel("Montant (€)" )).toBeVisible();
  await expect(page.getByRole("button", { name: "Ajouter" })).toBeVisible();
});
