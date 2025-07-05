import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom';
import { vi, beforeEach } from 'vitest';
import ProjectPanel from './ProjectPanel';
import '../i18n';

vi.mock('../wailsjs/go/service/DataService', () => ({
  ListProjects: vi.fn(),
  CreateProject: vi.fn(),
  DeleteProject: vi.fn(),
}), { virtual: true });

import { ListProjects, CreateProject, DeleteProject } from '../wailsjs/go/service/DataService';

beforeEach(() => {
  vi.clearAllMocks();
});

test('selects a project', async () => {
  ListProjects.mockResolvedValueOnce([
    { id: 1, name: 'One' },
    { id: 2, name: 'Two' },
  ]);
  const onSelect = vi.fn();
  render(<ProjectPanel activeId={1} onSelect={onSelect} />);
  await screen.findByText('One');
  fireEvent.click(screen.getByText('Two'));
  expect(onSelect).toHaveBeenCalledWith(2);
});

test('creates a project and shows it in the list', async () => {
  ListProjects
    .mockResolvedValueOnce([])
    .mockResolvedValueOnce([{ id: 1, name: 'Foo' }]);
  CreateProject.mockResolvedValueOnce({ id: 1, name: 'Foo' });
  const onSelect = vi.fn();
  render(<ProjectPanel activeId={0} onSelect={onSelect} />);
  await screen.findByRole('textbox', { name: /Neues Projekt/i });
  fireEvent.change(screen.getByLabelText(/Neues Projekt/i), { target: { value: 'Foo' } });
  fireEvent.click(screen.getByRole('button', { name: /Erstellen/i }));

  await waitFor(() => expect(CreateProject).toHaveBeenCalledWith('Foo'));
  expect(await screen.findByText('Foo')).toBeInTheDocument();
  expect(onSelect).toHaveBeenCalledWith(1);
});

test('deletes a project', async () => {
  ListProjects
    .mockResolvedValueOnce([
      { id: 1, name: 'Alpha' },
      { id: 2, name: 'Beta' },
    ])
    .mockResolvedValueOnce([{ id: 2, name: 'Beta' }]);
  DeleteProject.mockResolvedValueOnce();
  render(<ProjectPanel activeId={1} />);
  await screen.findByText('Alpha');
  fireEvent.click(screen.getAllByRole('button', { name: /Löschen/i })[0]);

  await waitFor(() => expect(DeleteProject).toHaveBeenCalledWith(1));
  await waitFor(() => expect(screen.queryByText('Alpha')).not.toBeInTheDocument());
});

test('shows create error', async () => {
  ListProjects.mockResolvedValueOnce([]);
  CreateProject.mockRejectedValueOnce(new Error('fail'));
  render(<ProjectPanel activeId={0} />);
  await screen.findByRole('textbox', { name: /Neues Projekt/i });
  fireEvent.change(screen.getByLabelText(/Neues Projekt/i), { target: { value: 'Foo' } });
  fireEvent.click(screen.getByRole('button', { name: /Erstellen/i }));
  expect(await screen.findByText('fail')).toBeInTheDocument();
});


