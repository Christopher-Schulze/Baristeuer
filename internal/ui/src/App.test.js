import { render, fireEvent } from '@testing-library/svelte'
import { describe, it, expect, vi, beforeEach } from 'vitest'
import '@testing-library/jest-dom'
import App from './App.svelte'

vi.mock('./wailsjs/go/service/DataService', () => {
  const incomes = []
  const expenses = []
  return {
    Backend: {
      CreateProject: vi.fn().mockResolvedValue({ id: 1 }),
      ListIncomes: vi.fn(() => Promise.resolve(incomes)),
      ListExpenses: vi.fn(() => Promise.resolve(expenses)),
      AddIncome: vi.fn((_p, source, amount) => {
        incomes.push({ source, amount })
        return Promise.resolve()
      }),
      AddExpense: vi.fn((_p, desc, amount) => {
        expenses.push({ desc, amount })
        return Promise.resolve()
      }),
    },
    Generator: {
      GenerateReport: vi.fn().mockResolvedValue('test.pdf'),
    },
  }
})

beforeEach(() => {
  window.backend = { Generator: { GenerateReport: vi.fn().mockResolvedValue('test.pdf') } }
  window.open = vi.fn()
})

describe('App component', () => {
  it('allows basic interactions', async () => {
    const { getByText, getByPlaceholderText, getAllByPlaceholderText, getAllByText, getByTitle } = render(App)

    // create project
    await fireEvent.input(getByPlaceholderText('Projektname'), { target: { value: 'Test' } })
    await fireEvent.click(getByText('Anlegen'))

    // income
    await fireEvent.input(getByPlaceholderText('Quelle'), { target: { value: 'Job' } })
    await fireEvent.input(getAllByPlaceholderText('Betrag')[0], { target: { value: '100' } })
    await fireEvent.click(getAllByText('Hinzufügen')[0])
    await Promise.resolve()
    expect(getByText('Job')).toBeInTheDocument()

    // expense
    await fireEvent.input(getByPlaceholderText('Beschreibung'), { target: { value: 'Food' } })
    await fireEvent.input(getAllByPlaceholderText('Betrag')[1], { target: { value: '50' } })
    await fireEvent.click(getAllByText('Hinzufügen')[1])
    await Promise.resolve()
    expect(getByText('Food')).toBeInTheDocument()

    // PDF preview toggle
    await fireEvent.click(getByText('PDF Vorschau'))
    expect(getByText('PDF erzeugen')).toBeInTheDocument()
  })

  it('generates a PDF and shows download link', async () => {
    const { getByText, getByTitle, findByText } = render(App)

    await fireEvent.click(getByText('PDF Vorschau'))
    await fireEvent.click(getByText('PDF erzeugen'))

    await findByText('Download')
    const link = getByText('Download')
    expect(link).toHaveAttribute('href', expect.stringContaining('file://test.pdf'))
  })

  it('shows an error when PDF generation fails', async () => {
    const errorMessage = 'boom'
    window.backend.Generator.GenerateReport = vi.fn().mockRejectedValue(new Error(errorMessage))

    const { getByText, findByText } = render(App)
    await fireEvent.click(getByText('PDF Vorschau'))
    await fireEvent.click(getByText('PDF erzeugen'))

    expect(await findByText(errorMessage)).toBeInTheDocument()
  })
})
