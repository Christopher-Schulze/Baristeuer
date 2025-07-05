import { render, screen, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom';
import MemberTable from './MemberTable';
import '../i18n';

const onDelete = vi.fn();

beforeEach(() => {
  vi.clearAllMocks();
});

test('renders empty state', () => {
  render(<MemberTable members={[]} onDelete={onDelete} />);
  expect(screen.getByText(/Keine Mitglieder vorhanden/i)).toBeInTheDocument();
});

test('calls delete callback', () => {
  const members = [{ id: 1, name: 'Alice', email: 'a@example.com', joinDate: '2024-01-01' }];
  render(<MemberTable members={members} onDelete={onDelete} />);
  fireEvent.click(screen.getByRole('button', { name: /Löschen/i }));
  expect(onDelete).toHaveBeenCalledWith(1);
});
