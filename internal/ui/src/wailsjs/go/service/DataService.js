let incomes = []
let expenses = []

export const Backend = {
  CreateProject: async (name) => ({ id: 1 }),
  ListIncomes: async () => incomes,
  AddIncome: async (projectId, source, amount) => {
    incomes.push({ source, amount: parseFloat(amount) })
  },
  ListExpenses: async () => expenses,
  AddExpense: async (projectId, desc, amount) => {
    expenses.push({ desc, amount: parseFloat(amount) })
  },
}

