import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom';
import { vi, beforeEach } from 'vitest';
import LoginDialog from './components/LoginDialog';

vi.mock('./wailsjs/go/service/DataService', () => ({
  AuthenticateUser: vi.fn(),
  CreateUser: vi.fn(),
}), { virtual: true });

import { AuthenticateUser, CreateUser } from './wailsjs/go/service/DataService';

beforeEach(() => {
  vi.clearAllMocks();
});

test('logs in successfully', async () => {
  AuthenticateUser.mockResolvedValueOnce('tok');
  render(<LoginDialog open={true} onSuccess={vi.fn()} />);
  fireEvent.change(screen.getByLabelText(/Benutzername/i), { target: { value: 'a' } });
  fireEvent.change(screen.getByLabelText(/Passwort/i), { target: { value: 'b' } });
  fireEvent.click(screen.getByRole('button', { name: 'OK' }));
  await waitFor(() => expect(AuthenticateUser).toHaveBeenCalledWith('a', 'b'));
});

test('registers new user', async () => {
  CreateUser.mockResolvedValueOnce();
  AuthenticateUser.mockResolvedValueOnce('tok');
  const onSuccess = vi.fn();
  render(<LoginDialog open={true} onSuccess={onSuccess} />);
  fireEvent.click(screen.getByRole('button', { name: /Registrieren/i }));
  fireEvent.change(screen.getByLabelText(/Benutzername/i), { target: { value: 'x' } });
  fireEvent.change(screen.getByLabelText(/Passwort/i), { target: { value: 'y' } });
  fireEvent.click(screen.getByRole('button', { name: 'OK' }));
  await waitFor(() => expect(CreateUser).toHaveBeenCalledWith('x', 'y'));
  await waitFor(() => expect(onSuccess).toHaveBeenCalled());
});
