import { render, fireEvent } from '@testing-library/svelte'
import { describe, it, expect } from 'vitest'
import '@testing-library/jest-dom'
import Counter from './Counter.svelte'

describe('Counter component', () => {
  it('increments count on click', async () => {
    const { getByText } = render(Counter)
    const button = getByText(/count is 0/i)
    await fireEvent.click(button)
    expect(getByText(/count is 1/i)).toBeInTheDocument()
  })
})
