(() => {
  const projects = [];
  const incomes = {};
  const expenses = {};
  function ensure(pid, map) {
    if (!map[pid]) map[pid] = [];
    return map[pid];
  }
  window.go = { service: { DataService: {} } };
  const svc = window.go.service.DataService;
  svc.ListProjects = async () => projects;
  svc.CreateProject = async (name) => {
    const id = projects.length + 1;
    const p = { id, name };
    projects.push(p);
    return p;
  };
  svc.DeleteProject = async (id) => {
    const i = projects.findIndex((p) => p.id === id);
    if (i > -1) projects.splice(i, 1);
    delete incomes[id];
    delete expenses[id];
  };
  svc.ListIncomes = async (pid) => ensure(pid, incomes);
  svc.AddIncome = async (pid, source, amount) => {
    const list = ensure(pid, incomes);
    const id = list.length + 1;
    const it = { id, source, amount };
    list.push(it);
    return it;
  };
  svc.DeleteIncome = async (id) => {
    for (const pid of Object.keys(incomes)) {
      const list = incomes[pid];
      const idx = list.findIndex((i) => i.id === id);
      if (idx > -1) {
        list.splice(idx, 1);
        break;
      }
    }
  };
  svc.ListExpenses = async (pid) => ensure(pid, expenses);
  svc.AddExpense = async (pid, desc, amount) => {
    const list = ensure(pid, expenses);
    const id = list.length + 1;
    const it = { id, description: desc, amount };
    list.push(it);
    return it;
  };
  svc.DeleteExpense = async (id) => {
    for (const pid of Object.keys(expenses)) {
      const list = expenses[pid];
      const idx = list.findIndex((e) => e.id === id);
      if (idx > -1) {
        list.splice(idx, 1);
        break;
      }
    }
  };

  const exports = [];
  svc.ExportDatabase = async (dest) => {
    exports.push({ type: "db", dest });
  };
  svc.ExportProjectCSV = async (pid, dest) => {
    exports.push({ type: "csv", pid, dest });
  };
  svc.RestoreDatabase = async () => {};
  svc.SetLogLevel = async () => {};
  window.__exports = exports;
})();
