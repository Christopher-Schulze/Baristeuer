import { render, screen, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom';
import IncomeForm from './IncomeForm';
import '../i18n';

const onSubmit = vi.fn();

beforeEach(() => {
  vi.clearAllMocks();
});

test('shows validation errors', async () => {
  render(<IncomeForm onSubmit={onSubmit} />);
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));
  expect(await screen.findByText(/Quelle und Betrag erforderlich/i)).toBeInTheDocument();

  fireEvent.change(screen.getByLabelText(/Quelle/i), { target: { value: 'Sponsor' } });
  fireEvent.change(screen.getByLabelText(/Betrag/i), { target: { value: '-5' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));
  expect(await screen.findByText(/Betrag muss eine positive Zahl sein/i)).toBeInTheDocument();
});

test('submits valid data', async () => {
  render(<IncomeForm onSubmit={onSubmit} />);
  fireEvent.change(screen.getByLabelText(/Quelle/i), { target: { value: 'Job' } });
  fireEvent.change(screen.getByLabelText(/Betrag/i), { target: { value: '10' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));
  expect(onSubmit).toHaveBeenCalled();
});

test('shows submit error', async () => {
  onSubmit.mockImplementation(async (_s, _a, setErr) => {
    setErr('fail');
  });
  render(<IncomeForm onSubmit={onSubmit} />);
  fireEvent.change(screen.getByLabelText(/Quelle/i), { target: { value: 'Foo' } });
  fireEvent.change(screen.getByLabelText(/Betrag/i), { target: { value: '2' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));
  expect(await screen.findByText('fail')).toBeInTheDocument();
});
