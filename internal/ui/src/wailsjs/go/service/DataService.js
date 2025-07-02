export function CreateProject(name) {
  return window.go.service.DataService.CreateProject(name);
}

export function ListIncomes(projectID) {
  return window.go.service.DataService.ListIncomes(projectID);
}

export function AddExpense(projectID, category, amount) {
  return window.go.service.DataService.AddExpense(projectID, category, amount);
}
