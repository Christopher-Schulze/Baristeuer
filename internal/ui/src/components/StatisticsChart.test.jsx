import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import { vi } from "vitest";
import StatisticsChart from "./StatisticsChart";
import "../i18n";

vi.mock(
  "../wailsjs/go/service/DataService",
  () => ({
    GenerateStatistics: vi.fn(),
  }),
  { virtual: true },
);

import { GenerateStatistics } from "../wailsjs/go/service/DataService";

beforeEach(() => {
  vi.clearAllMocks();
});

test("renders line chart", async () => {
  GenerateStatistics.mockResolvedValue({
    averageIncome: 10,
    averageExpense: 5,
    medianIncome: 9,
    medianExpense: 4,
    stdDevIncome: 0,
    stdDevExpense: 0,
    trend: 6,
  });
  render(<StatisticsChart projectId={1} />);
  expect(await screen.findByRole("img")).toBeInTheDocument();
});
