const incomes = [];
const expenses = [];
export const Backend = {
  CreateProject: async () => ({ id: 1 }),
  ListIncomes: async () => incomes,
  ListExpenses: async () => expenses,
  AddIncome: async (projectId, source, amount) => {
    incomes.push({ source, amount: parseFloat(amount) });
  },
  AddExpense: async (projectId, desc, amount) => {
    expenses.push({ desc, amount: parseFloat(amount) });
  },
  ListPlugins: async () => [],
  SetPluginEnabled: async () => {},
};
