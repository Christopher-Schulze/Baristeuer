import { test, expect } from "@playwright/test";
import { fileURLToPath } from "url";
import path from "path";

test.beforeEach(async ({ page }) => {
  const dir = path.dirname(fileURLToPath(import.meta.url));
  const scriptPath = path.join(dir, "mockDataService.js");
  await page.addInitScript({ path: scriptPath });
});

test("switch UI language to English", async ({ page }) => {
  await page.goto("/");

  await page.getByRole("button", { name: "DE" }).click();
  await page.getByRole("option", { name: "EN" }).click();

  await expect(page.getByRole("tab", { name: "Projects" })).toBeVisible();
  await expect(page.getByRole("tab", { name: "Income" })).toBeVisible();
  await expect(page.getByRole("tab", { name: "Expenses" })).toBeVisible();

  await page.getByRole("tab", { name: "Projects" }).click();
  await expect(page.getByLabel("New Project")).toBeVisible();
  await expect(page.getByRole("button", { name: "Create" })).toBeVisible();

  await page.getByRole("tab", { name: "Income" }).click();
  await expect(page.getByRole("heading", { name: "New Income" })).toBeVisible();
  await expect(page.getByLabel("Source")).toBeVisible();
  await expect(page.getByLabel("Amount (€)" )).toBeVisible();
  await expect(page.getByRole("button", { name: "Add" })).toBeVisible();
});
