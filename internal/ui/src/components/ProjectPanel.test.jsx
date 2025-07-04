import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom';
import { vi, beforeEach } from 'vitest';
import ProjectPanel from './ProjectPanel';

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


