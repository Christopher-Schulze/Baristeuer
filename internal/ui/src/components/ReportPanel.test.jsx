import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom';
import { vi } from 'vitest';
import ReportPanel from './ReportPanel';
import '../i18n';

vi.mock('../wailsjs/go/service/DataService', () => ({
  GenerateStatistics: vi.fn(),
}), { virtual: true });

import { GenerateStatistics } from '../wailsjs/go/service/DataService';

beforeEach(() => {
  vi.clearAllMocks();
});

test('renders statistics', async () => {
  GenerateStatistics.mockResolvedValue({ averageIncome: 10, averageExpense: 5, trend: 5 });
  render(<ReportPanel projectId={1} />);
  expect(await screen.findByText(/Trend/i)).toBeInTheDocument();
});
