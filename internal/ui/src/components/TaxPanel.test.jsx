import { render, screen, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom';
import TaxPanel from './TaxPanel';
import '../i18n';

test('year dropdown includes 2027', () => {
  render(<TaxPanel projectId={1} />);
  fireEvent.mouseDown(screen.getByLabelText(/Jahr/i));
  expect(screen.getByRole('option', { name: '2027' })).toBeInTheDocument();
});
