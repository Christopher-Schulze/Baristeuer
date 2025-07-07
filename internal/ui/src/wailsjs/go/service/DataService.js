const projects = [];
const incomes = [];
const expenses = [];
let nextProjectId = 1;
let nextIncomeId = 1;
let nextExpenseId = 1;

export const Backend = {
  async CreateProject(name) {
    const p = { id: nextProjectId++, name };
    projects.push(p);
    return p;
  },
  async ListProjects() {
    return projects;
  },
  async DeleteProject(id) {
    const idx = projects.findIndex(p => p.id === id);
    if (idx !== -1) projects.splice(idx, 1);
    for (let i = incomes.length - 1; i >= 0; i--) {
      if (incomes[i].projectId === id) incomes.splice(i, 1);
    }
    for (let i = expenses.length - 1; i >= 0; i--) {
      if (expenses[i].projectId === id) expenses.splice(i, 1);
    }
  },
  async AddIncome(projectId, source, amount) {
    const rec = { id: nextIncomeId++, projectId, source, amount: parseFloat(amount) };
    incomes.push(rec);
    return rec;
  },
  async ListIncomes(projectId) {
    return incomes.filter(i => i.projectId === projectId);
  },
  async AddExpense(projectId, desc, amount) {
    const rec = { id: nextExpenseId++, projectId, desc, amount: parseFloat(amount) };
    expenses.push(rec);
    return rec;
  },
  async ListExpenses(projectId) {
    return expenses.filter(e => e.projectId === projectId);
  }
};
