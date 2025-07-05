import { render, screen, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom';
import SettingsPanel from './SettingsPanel';
import '../i18n';

test('tax year options include 2027', () => {
  render(<SettingsPanel projectId={1} />);
  fireEvent.mouseDown(screen.getByLabelText(/Steuerjahr|Tax Year/i));
  expect(screen.getByRole('option', { name: '2027' })).toBeInTheDocument();
});
