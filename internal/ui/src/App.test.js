import { render, fireEvent } from '@testing-library/svelte'
import { describe, it, expect } from 'vitest'
import '@testing-library/jest-dom'
import App from './App.svelte'

describe('App component', () => {
  it('allows basic interactions', async () => {
    const { getByText, getByPlaceholderText, getAllByPlaceholderText, getAllByText, getByTitle } = render(App)

    // income
    await fireEvent.input(getByPlaceholderText('Source'), { target: { value: 'Job' } })
    await fireEvent.input(getAllByPlaceholderText('Amount')[0], { target: { value: '100' } })
    await fireEvent.click(getAllByText('Add')[0])
    expect(getByText('Job')).toBeInTheDocument()

    // expense
    await fireEvent.input(getByPlaceholderText('Description'), { target: { value: 'Food' } })
    await fireEvent.input(getAllByPlaceholderText('Amount')[1], { target: { value: '50' } })
    await fireEvent.click(getAllByText('Add')[1])
    expect(getByText('Food')).toBeInTheDocument()

    // PDF preview toggle
    await fireEvent.click(getByText('PDF Preview'))
    expect(getByTitle('PDF Preview')).toBeInTheDocument()
  })
})
