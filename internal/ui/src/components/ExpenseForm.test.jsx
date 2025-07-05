import { render, screen, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom';
import ExpenseForm from './ExpenseForm';
import '../i18n';

const onSubmit = vi.fn();

beforeEach(() => {
  vi.clearAllMocks();
});

test('shows validation errors', async () => {
  render(<ExpenseForm onSubmit={onSubmit} />);
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));
  expect(await screen.findByText(/Beschreibung und Betrag erforderlich/i)).toBeInTheDocument();

  fireEvent.change(screen.getByLabelText(/Beschreibung/i), { target: { value: 'Fee' } });
  fireEvent.change(screen.getByLabelText(/Betrag/i), { target: { value: '-1' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));
  expect(await screen.findByText(/Betrag muss eine positive Zahl sein/i)).toBeInTheDocument();
});

test('submits valid data', async () => {
  render(<ExpenseForm onSubmit={onSubmit} />);
  fireEvent.change(screen.getByLabelText(/Beschreibung/i), { target: { value: 'Rent' } });
  fireEvent.change(screen.getByLabelText(/Betrag/i), { target: { value: '5' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));
  expect(onSubmit).toHaveBeenCalled();
});
