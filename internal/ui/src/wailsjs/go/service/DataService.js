export function CreateProject(name) {
  return window.go.service.DataService.CreateProject(name);
}

export function ListIncomes(projectID) {
  return window.go.service.DataService.ListIncomes(projectID);
}

export function AddIncome(projectID, source, amount) {
  return window.go.service.DataService.AddIncome(projectID, source, amount);
}

export function AddExpense(projectID, category, amount) {
  return window.go.service.DataService.AddExpense(projectID, category, amount);
}

export function ListExpenses(projectID) {
  return window.go.service.DataService.ListExpenses(projectID);
}
