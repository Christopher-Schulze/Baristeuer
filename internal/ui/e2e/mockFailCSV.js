(() => {
  const svc = window.go && window.go.service && window.go.service.DataService;
  if (svc) {
    svc.ExportProjectCSV = async () => {
      throw new Error("CSV export failed");
    };
  }
})();
