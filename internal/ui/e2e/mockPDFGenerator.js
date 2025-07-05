(() => {
  window.go = window.go || {};
  window.go.pdf = { Generator: {} };
  const gen = window.go.pdf.Generator;
  const calls = [];
  function record(name) {
    return async (...args) => {
      calls.push({ name, args });
      return `/tmp/${name}.pdf`;
    };
  }
  [
    'GenerateReport',
    'GenerateKSt1',
    'GenerateAnlageGem',
    'GenerateAnlageGK',
    'GenerateKSt1F',
    'GenerateAnlageSport',
    'GenerateAllForms',
  ].forEach((m) => {
    gen[m] = record(m);
  });
  window.__pdfCalls = calls;
})();
