import { render, screen, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom';
import MemberForm from './MemberForm';
import '../i18n';

const onSubmit = vi.fn();

beforeEach(() => {
  vi.clearAllMocks();
});

test('shows validation errors', async () => {
  render(<MemberForm onSubmit={onSubmit} />);
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));
  expect(await screen.findByText(/Name und E-Mail erforderlich/i)).toBeInTheDocument();
  fireEvent.change(screen.getByLabelText(/Name/i), { target: { value: 'A' } });
  fireEvent.change(screen.getByLabelText(/E-Mail/i), { target: { value: 'foo' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));
  expect(await screen.findByText(/Ungültige E-Mail-Adresse/i)).toBeInTheDocument();
});

test('submits valid data', async () => {
  render(<MemberForm onSubmit={onSubmit} />);
  fireEvent.change(screen.getByLabelText(/Name/i), { target: { value: 'Max' } });
  fireEvent.change(screen.getByLabelText(/E-Mail/i), { target: { value: 'max@example.com' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));
  expect(onSubmit).toHaveBeenCalled();
});

test('shows submit error', async () => {
  onSubmit.mockImplementation(async (_n, _e, _d, setErr) => {
    setErr('fail');
  });
  render(<MemberForm onSubmit={onSubmit} />);
  fireEvent.change(screen.getByLabelText(/Name/i), { target: { value: 'Bob' } });
  fireEvent.change(screen.getByLabelText(/E-Mail/i), { target: { value: 'bob@example.com' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));
  expect(await screen.findByText('fail')).toBeInTheDocument();
});

test('prefills data when editing', () => {
  const item = { id: 1, name: 'Eve', email: 'eve@example.com', joinDate: '2024-01-05' };
  render(<MemberForm onSubmit={onSubmit} editItem={item} />);
  expect(screen.getByDisplayValue('Eve')).toBeInTheDocument();
  expect(screen.getByDisplayValue('eve@example.com')).toBeInTheDocument();
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));
  expect(onSubmit).toHaveBeenCalled();
});
