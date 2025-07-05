import { test, expect } from "@playwright/test";
import { fileURLToPath } from "url";
import path from "path";

// Setup mocks and record service calls

test.beforeEach(async ({ page }) => {
  const dir = path.dirname(fileURLToPath(import.meta.url));
  await page.addInitScript({ path: path.join(dir, "mockDataService.js") });
  await page.addInitScript({ path: path.join(dir, "mockPDFGenerator.js") });
  await page.addInitScript(() => {
    window.__settingsCalls = [];
    const svc = window.go.service.DataService;
    const pdf = window.go.pdf.Generator;
    const origSetLogLevel = svc.SetLogLevel;
    svc.SetLogLevel = async (level) => {
      window.__settingsCalls.push({ func: "SetLogLevel", args: [level] });
      return origSetLogLevel(level);
    };
    const origSetTaxYear = pdf.SetTaxYear || (async () => {});
    pdf.SetTaxYear = async (year) => {
      window.__settingsCalls.push({ func: "SetTaxYear", args: [year] });
      return origSetTaxYear(year);
    };
  });
});

test("update settings", async ({ page }) => {
  await page.goto("/");

  await page.getByRole("tab", { name: /Einstellungen/i }).click();

  // change log level
  await page.getByRole("button", { name: "info" }).click();
  await page.getByRole("option", { name: "debug" }).click();
  await page.getByRole("button", { name: /Anwenden/i }).nth(0).click();

  // change log format
  await page.getByRole("button", { name: "text" }).click();
  await page.getByRole("option", { name: "json" }).click();
  await page.getByRole("button", { name: /Anwenden/i }).nth(1).click();

  // change tax year to 2027
  await page.getByLabel(/Steuerjahr/i).fill("2027");
  await page.getByRole("button", { name: /Anwenden/i }).nth(2).click();

  const calls = await page.evaluate(() => window.__settingsCalls);
  expect(calls).toContainEqual({ func: "SetLogLevel", args: ["debug"] });
  expect(calls).toContainEqual({ func: "SetTaxYear", args: [2027] });
});
