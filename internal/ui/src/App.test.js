import { render, fireEvent } from '@testing-library/svelte'
import { describe, it, expect } from 'vitest'
import '@testing-library/jest-dom'
import App from './App.svelte'

describe('App component', () => {
  it('allows basic interactions', async () => {
    const { getByText, getByPlaceholderText, getAllByPlaceholderText, getAllByText, getByTitle } = render(App)

    // income
    await fireEvent.input(getByPlaceholderText('Quelle'), { target: { value: 'Job' } })
    await fireEvent.input(getAllByPlaceholderText('Betrag')[0], { target: { value: '100' } })
    await fireEvent.click(getAllByText('Hinzufügen')[0])
    expect(getByText('Job')).toBeInTheDocument()

    // expense
    await fireEvent.input(getByPlaceholderText('Beschreibung'), { target: { value: 'Food' } })
    await fireEvent.input(getAllByPlaceholderText('Betrag')[1], { target: { value: '50' } })
    await fireEvent.click(getAllByText('Hinzufügen')[1])
    expect(getByText('Food')).toBeInTheDocument()

    // PDF preview toggle
    await fireEvent.click(getByText('PDF'))
    await fireEvent.click(getByText('PDF Vorschau'))
    expect(getByText('PDF erzeugen')).toBeInTheDocument()
  })
})
