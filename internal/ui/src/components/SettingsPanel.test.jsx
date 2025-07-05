import { render, screen, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom';
import { vi, beforeEach } from 'vitest';
import SettingsPanel from './SettingsPanel';
import '../i18n';

vi.mock('../wailsjs/go/service/DataService', () => ({
  ExportDatabase: vi.fn(),
  RestoreDatabase: vi.fn(),
  SetLogLevel: vi.fn(),
  ExportProjectCSV: vi.fn(),
}), { virtual: true });

vi.mock('../wailsjs/go/pdf/Generator', () => ({
  SetTaxYear: vi.fn(),
}), { virtual: true });

beforeEach(() => {
  vi.clearAllMocks();
});

test('shows success message when log format applied', async () => {
  render(<SettingsPanel projectId={1} />);
  const applyButtons = screen.getAllByRole('button', { name: /Anwenden/i });
  fireEvent.click(applyButtons[1]);
  expect(await screen.findByText('settings.applied')).toBeInTheDocument();
});
