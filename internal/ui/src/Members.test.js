import { render, fireEvent } from '@testing-library/svelte';
import { describe, it, expect, beforeEach } from 'vitest';
import '@testing-library/jest-dom';
import Members from './Members.svelte';
import { __reset } from './wailsjs/go/service/DataService';

describe('Members component', () => {
  beforeEach(() => {
    __reset();
  });

  it('handles add, edit and delete', async () => {
    const { getByText, getByPlaceholderText, queryByText, getByDisplayValue } = render(Members);

    await fireEvent.input(getByPlaceholderText('Name'), { target: { value: 'Alice' } });
    await fireEvent.input(getByPlaceholderText('Email'), { target: { value: 'alice@example.com' } });
    await fireEvent.input(getByPlaceholderText('Datum'), { target: { value: '2024-01-01' } });
    await fireEvent.click(getByText('Hinzufügen'));
    expect(getByText('alice@example.com')).toBeInTheDocument();

    await fireEvent.click(getByText('Bearbeiten'));
    await fireEvent.input(getByDisplayValue('Alice'), { target: { value: 'Alicia' } });
    await fireEvent.click(getByText('Speichern'));
    expect(getByText('Alicia')).toBeInTheDocument();

    await fireEvent.click(getByText('Löschen'));
    expect(queryByText('Alicia')).not.toBeInTheDocument();
  });
});
