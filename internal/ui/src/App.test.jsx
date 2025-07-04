import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom';
import { vi, beforeEach } from 'vitest';
import App from './App';
import './i18n';

// mock the DataService module used by App
vi.mock('./wailsjs/go/service/DataService', () => ({
  CreateProject: vi.fn(),
  ListProjects: vi.fn(),
  GetProject: vi.fn(),
  UpdateProject: vi.fn(),
  DeleteProject: vi.fn(),
  AddExpense: vi.fn(),
  UpdateExpense: vi.fn(),
  DeleteExpense: vi.fn(),
  ListExpenses: vi.fn(),
  AddIncome: vi.fn(),
  UpdateIncome: vi.fn(),
  DeleteIncome: vi.fn(),
  ListIncomes: vi.fn(),
  AddMember: vi.fn(),
  ListMembers: vi.fn(),
  DeleteMember: vi.fn(),
  CalculateProjectTaxes: vi.fn(),
}), { virtual: true });

// import the mocked functions for easier access
import {
  CreateProject,
  ListProjects,
  GetProject,
  UpdateProject,
  DeleteProject,
  AddExpense,
  UpdateExpense,
  DeleteExpense,
  ListExpenses,
  AddIncome,
  UpdateIncome,
  DeleteIncome,
  ListIncomes,
  AddMember,
  ListMembers,
  DeleteMember,
  CalculateProjectTaxes,
} from './wailsjs/go/service/DataService';

beforeEach(() => {
  vi.clearAllMocks();
});

test('renders app heading', async () => {
  ListExpenses.mockResolvedValueOnce([]);
  ListIncomes.mockResolvedValueOnce([]);
  ListMembers.mockResolvedValueOnce([]);
  render(<App />);
  expect(await screen.findByRole('heading', { name: /Baristeuer/i })).toBeInTheDocument();
});

// Add Income

test('adds a new income', async () => {
  ListExpenses.mockResolvedValueOnce([]);
  ListIncomes.mockResolvedValueOnce([]).mockResolvedValueOnce([{ id: 1, source: 'Donation', amount: 50 }]);
  ListMembers.mockResolvedValueOnce([]);
  AddIncome.mockResolvedValueOnce();
  render(<App />);
  await screen.findByRole('heading', { name: /Baristeuer/i });

  fireEvent.change(screen.getByLabelText(/Quelle/i), { target: { value: 'Donation' } });
  fireEvent.change(screen.getByLabelText(/Betrag/i), { target: { value: '50' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));

  await waitFor(() => expect(AddIncome).toHaveBeenCalled());
  expect(await screen.findByText('Donation')).toBeInTheDocument();
  expect(screen.getByText('50.00')).toBeInTheDocument();
});

// Failed add income

test('shows error when adding income fails', async () => {
  ListExpenses.mockResolvedValueOnce([]);
  ListIncomes.mockResolvedValueOnce([]);
  ListMembers.mockResolvedValueOnce([]);
  AddIncome.mockRejectedValueOnce(new Error('fail'));
  render(<App />);
  await screen.findByRole('heading', { name: /Baristeuer/i });

  fireEvent.change(screen.getByLabelText(/Quelle/i), { target: { value: 'Foo' } });
  fireEvent.change(screen.getByLabelText(/Betrag/i), { target: { value: '5' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));

  const alert = await screen.findByRole('alert');
  expect(alert).toHaveTextContent('fail');
});

test('shows error when adding member fails', async () => {
  ListExpenses.mockResolvedValueOnce([]);
  ListIncomes.mockResolvedValueOnce([]);
  ListMembers.mockResolvedValueOnce([]);
  AddMember.mockRejectedValueOnce(new Error('oops'));
  render(<App />);
  await screen.findByRole('heading', { name: /Baristeuer/i });

  fireEvent.click(screen.getByRole('tab', { name: /Mitglieder/i }));
  fireEvent.change(screen.getByLabelText(/Name/i), { target: { value: 'A' } });
  fireEvent.change(screen.getByLabelText(/E-Mail/i), { target: { value: 'a@example.com' } });
  fireEvent.change(screen.getByLabelText(/Beitrittsdatum/i), { target: { value: '2024-01-10' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));

  const alertMember = await screen.findByRole('alert');
  expect(alertMember).toHaveTextContent('oops');
});

// Edit Income

test('edits an income', async () => {
  ListExpenses.mockResolvedValueOnce([]);
  ListIncomes.mockResolvedValueOnce([{ id: 1, source: 'Old', amount: 10 }]).mockResolvedValueOnce([{ id: 1, source: 'New', amount: 20 }]);
  ListMembers.mockResolvedValueOnce([]);
  UpdateIncome.mockResolvedValueOnce();
  render(<App />);
  await screen.findByText('Old');

  fireEvent.click(screen.getByRole('button', { name: /Bearbeiten/i }));
  fireEvent.change(screen.getByLabelText(/Quelle/i), { target: { value: 'New' } });
  fireEvent.change(screen.getByLabelText(/Betrag/i), { target: { value: '20' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));

  await waitFor(() => expect(UpdateIncome).toHaveBeenCalledWith(1, 1, 'New', 20));
  expect(await screen.findByText('New')).toBeInTheDocument();
});

// Delete Income

test('deletes an income', async () => {
  ListExpenses.mockResolvedValueOnce([]);
  ListIncomes.mockResolvedValueOnce([{ id: 1, source: 'Del', amount: 30 }]).mockResolvedValueOnce([]);
  ListMembers.mockResolvedValueOnce([]);
  DeleteIncome.mockResolvedValueOnce();
  render(<App />);
  await screen.findByText('Del');

  fireEvent.click(screen.getByRole('button', { name: /Löschen/i }));

  await waitFor(() => expect(DeleteIncome).toHaveBeenCalledWith(1));
  await waitFor(() => expect(screen.queryByText('Del')).not.toBeInTheDocument());
});

// Add Expense

test('adds a new expense', async () => {
  ListIncomes.mockResolvedValueOnce([]);
  ListExpenses.mockResolvedValueOnce([]).mockResolvedValueOnce([{ id: 1, description: 'Rent', amount: 15 }]);
  ListMembers.mockResolvedValueOnce([]);
  AddExpense.mockResolvedValueOnce();
  render(<App />);
  await screen.findByRole('heading', { name: /Baristeuer/i });

  fireEvent.click(screen.getByRole('tab', { name: /Ausgaben/i }));
  fireEvent.change(screen.getByLabelText(/Beschreibung/i), { target: { value: 'Rent' } });
  fireEvent.change(screen.getByLabelText(/Betrag/i), { target: { value: '15' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));

  await waitFor(() => expect(AddExpense).toHaveBeenCalled());
  expect(await screen.findByText('Rent')).toBeInTheDocument();
  expect(screen.getByText('15.00')).toBeInTheDocument();
});

// Edit Expense

test('edits an expense', async () => {
  ListIncomes.mockResolvedValueOnce([]);
  ListExpenses.mockResolvedValueOnce([{ id: 1, description: 'Coffee', amount: 3 }]).mockResolvedValueOnce([{ id: 1, description: 'Tea', amount: 4 }]);
  ListMembers.mockResolvedValueOnce([]);
  UpdateExpense.mockResolvedValueOnce();
  render(<App />);
  await screen.findByRole('heading', { name: /Baristeuer/i });
  fireEvent.click(screen.getByRole('tab', { name: /Ausgaben/i }));
  await screen.findByText('Coffee');

  fireEvent.click(screen.getByRole('button', { name: /Bearbeiten/i }));
  fireEvent.change(screen.getByLabelText(/Beschreibung/i), { target: { value: 'Tea' } });
  fireEvent.change(screen.getByLabelText(/Betrag/i), { target: { value: '4' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));

  await waitFor(() => expect(UpdateExpense).toHaveBeenCalledWith(1, 1, 'Tea', 4));
  expect(await screen.findByText('Tea')).toBeInTheDocument();
});

// Delete Expense

test('deletes an expense', async () => {
  ListIncomes.mockResolvedValueOnce([]);
  ListExpenses.mockResolvedValueOnce([{ id: 1, description: 'Coffee', amount: 3 }]).mockResolvedValueOnce([]);
  ListMembers.mockResolvedValueOnce([]);
  DeleteExpense.mockResolvedValueOnce();
  render(<App />);
  await screen.findByRole('heading', { name: /Baristeuer/i });
  fireEvent.click(screen.getByRole('tab', { name: /Ausgaben/i }));
  await screen.findByText('Coffee');

  fireEvent.click(screen.getByRole('button', { name: /Löschen/i }));

  await waitFor(() => expect(DeleteExpense).toHaveBeenCalledWith(1));
  await waitFor(() => expect(screen.queryByText('Coffee')).not.toBeInTheDocument());
});

// Calculate taxes

test('shows tax calculation result', async () => {
  ListExpenses.mockResolvedValueOnce([]);
  ListIncomes.mockResolvedValueOnce([]);
  ListMembers.mockResolvedValueOnce([]);
  CalculateProjectTaxes.mockResolvedValueOnce({ revenue: 100, expenses: 20, taxableIncome: 80, totalTax: 10 });
  render(<App />);
  await screen.findByRole('heading', { name: /Baristeuer/i });

  fireEvent.click(screen.getByRole('tab', { name: /Steuern/i }));
  fireEvent.click(screen.getByRole('button', { name: /Steuern berechnen/i }));

  expect(await screen.findByText('Einnahmen: 100.00 \u20AC')).toBeInTheDocument();
  expect(screen.getByText('Ausgaben: 20.00 \u20AC')).toBeInTheDocument();
  expect(screen.getByText('Steuerpflichtiges Einkommen: 80.00 \u20AC')).toBeInTheDocument();
  expect(screen.getByText('Gesamtsteuer: 10.00 \u20AC')).toBeInTheDocument();
});

// Add member

test('adds a new member', async () => {
  ListExpenses.mockResolvedValueOnce([]);
  ListIncomes.mockResolvedValueOnce([]);
  ListMembers.mockResolvedValueOnce([]).mockResolvedValueOnce([{ id: 1, name: 'Alice', email: 'a@example.com', joinDate: '2024-01-10' }]);
  AddMember.mockResolvedValueOnce();
  render(<App />);
  await screen.findByRole('heading', { name: /Baristeuer/i });

  fireEvent.click(screen.getByRole('tab', { name: /Mitglieder/i }));
  fireEvent.change(screen.getByLabelText(/Name/i), { target: { value: 'Alice' } });
  fireEvent.change(screen.getByLabelText(/E-Mail/i), { target: { value: 'a@example.com' } });
  fireEvent.change(screen.getByLabelText(/Beitrittsdatum/i), { target: { value: '2024-01-10' } });
  fireEvent.click(screen.getByRole('button', { name: /Hinzufügen/i }));

  await waitFor(() => expect(AddMember).toHaveBeenCalledWith('Alice', 'a@example.com', '2024-01-10'));
  expect(await screen.findByText('Alice')).toBeInTheDocument();
});

// Delete member

test('deletes a member', async () => {
  ListExpenses.mockResolvedValueOnce([]);
  ListIncomes.mockResolvedValueOnce([]);
  ListMembers.mockResolvedValueOnce([{ id: 1, name: 'Bob', email: 'b@example.com', joinDate: '2024-01-05' }]).mockResolvedValueOnce([]);
  DeleteMember.mockResolvedValueOnce();
  render(<App />);
  await screen.findByRole('heading', { name: /Baristeuer/i });

  fireEvent.click(screen.getByRole('tab', { name: /Mitglieder/i }));
  await screen.findByText('Bob');

  fireEvent.click(screen.getByRole('button', { name: /Löschen/i }));

  await waitFor(() => expect(DeleteMember).toHaveBeenCalledWith(1));
  await waitFor(() => expect(screen.queryByText('Bob')).not.toBeInTheDocument());
});
