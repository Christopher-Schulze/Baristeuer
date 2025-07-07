let incomes = []
let expenses = []

export const Backend = {
  async CreateProject() { return { id: 1 } },
  async ListIncomes() { return incomes },
  async ListExpenses() { return expenses },
  async AddIncome(_projectId, source, amount) {
    incomes.push({ source, amount })
  },
  async AddExpense(_projectId, desc, amount) {
    expenses.push({ desc, amount })
  },
};
export const Generator = {
  async GenerateReport() { return 'test.pdf'; }
};
if (typeof window !== 'undefined') {
  window.backend = { Generator };
}
