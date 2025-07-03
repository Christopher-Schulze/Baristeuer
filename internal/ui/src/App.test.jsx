import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom';
import { vi } from 'vitest';
import App from './App';

// mock the DataService module used by App
vi.mock('./wailsjs/go/service/DataService', () => ({
  AddExpense: vi.fn(),
  ListExpenses: vi.fn(() => Promise.resolve([])),
}), { virtual: true });

test('renders app heading', async () => {
  render(<App />);
  expect(await screen.findByRole('heading', { name: /Baristeuer/i })).toBeInTheDocument();
});

