import { render, screen, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom';
import { vi, beforeEach } from 'vitest';
import SettingsPanel from './SettingsPanel';
import '../i18n';

vi.mock('../wailsjs/go/service/DataService', () => ({
  ExportDatabase: vi.fn(),
  RestoreDatabase: vi.fn(),
  SetLogLevel: vi.fn(),
  SetLogFormat: vi.fn(),
  ExportProjectCSV: vi.fn(),
  GetFormName: vi.fn().mockResolvedValue('Club'),
  SetFormName: vi.fn(),
  GetFormTaxNumber: vi.fn().mockResolvedValue('11/111/11111'),
  SetFormTaxNumber: vi.fn(),
  GetFormAddress: vi.fn().mockResolvedValue('Street'),
  SetFormAddress: vi.fn(),
  GetTaxYear: vi.fn().mockResolvedValue(2025),
  SetTaxYear: vi.fn(),
  GetCloudUploadURL: vi.fn().mockResolvedValue('u'),
  SetCloudUploadURL: vi.fn(),
  GetCloudDownloadURL: vi.fn().mockResolvedValue('d'),
  SetCloudDownloadURL: vi.fn(),
  GetCloudToken: vi.fn().mockResolvedValue('t'),
  SetCloudToken: vi.fn(),
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

test('loads settings from service', async () => {
  render(<SettingsPanel projectId={1} />);
  expect(await screen.findByDisplayValue('Club')).toBeInTheDocument();
  expect(await screen.findByDisplayValue('11/111/11111')).toBeInTheDocument();
  expect(await screen.findByDisplayValue('Street')).toBeInTheDocument();
});

test('applies form name change', async () => {
  const { SetFormName } = await import('../wailsjs/go/service/DataService');
  render(<SettingsPanel projectId={1} />);
  const input = await screen.findByLabelText(/Vereinsname/i);
  fireEvent.change(input, { target: { value: 'New' } });
  const buttons = screen.getAllByRole('button', { name: /Anwenden/i });
  fireEvent.click(buttons[3]);
  expect(SetFormName).toHaveBeenCalledWith('New');
});
