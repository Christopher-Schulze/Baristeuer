const state = {
  incomes: [],
  expenses: [],
  members: [],
  nextMemberId: 1,
};

export const Backend = {
  async CreateProject(name) {
    return { id: 1 };
  },
  async ListIncomes() {
    return state.incomes;
  },
  async AddIncome(projectId, source, amount) {
    state.incomes.push({ source, amount });
  },
  async ListExpenses() {
    return state.expenses;
  },
  async AddExpense(projectId, desc, amount) {
    state.expenses.push({ desc, amount });
  },
  async ListMembers() {
    return state.members;
  },
  async AddMember(name, email, joinDate) {
    const m = { id: state.nextMemberId++, name, email, joinDate };
    state.members.push(m);
    return m;
  },
  async UpdateMember(id, name, email, joinDate) {
    const m = state.members.find((x) => x.id === id);
    if (m) {
      m.name = name;
      m.email = email;
      m.joinDate = joinDate;
    }
  },
  async DeleteMember(id) {
    state.members = state.members.filter((x) => x.id !== id);
  },
};

export function __reset() {
  state.incomes = [];
  state.expenses = [];
  state.members = [];
  state.nextMemberId = 1;
}
