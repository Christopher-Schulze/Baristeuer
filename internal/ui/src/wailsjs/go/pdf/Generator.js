export function GenerateReport(projectID) {
  return window.go.pdf.Generator.GenerateReport(projectID);
}

export function GenerateDetailedReport(projectID) {
  return window.go.pdf.Generator.GenerateDetailedReport(projectID);
}

export function GenerateKSt1(projectID) {
  return window.go.pdf.Generator.GenerateKSt1(projectID);
}

export function GenerateAnlageGem(projectID) {
  return window.go.pdf.Generator.GenerateAnlageGem(projectID);
}

export function GenerateAnlageGK(projectID) {
  return window.go.pdf.Generator.GenerateAnlageGK(projectID);
}

export function GenerateKSt1F(projectID) {
  return window.go.pdf.Generator.GenerateKSt1F(projectID);
}

export function GenerateAnlageSport(projectID) {
  return window.go.pdf.Generator.GenerateAnlageSport(projectID);
}

export function GenerateAllForms(projectID) {
  return window.go.pdf.Generator.GenerateAllForms(projectID);
}

export function SetTaxYear(year) {
  return window.go.pdf.Generator.SetTaxYear(year);
}
