import React from 'react';
import ReactDOM from 'react-dom/client';

function generateReport() {
  if (window.backend && window.backend.GenerateReport) {
    window.backend.GenerateReport('demo')
      .then(() => console.log('Report generated'))
      .catch((err) => console.error(err));
  } else {
    console.log('Backend not available');
  }
}

function App() {
  return (
    <div>
      Hello from Baristeuer UI
      <button onClick={generateReport}>Generate Report</button>
    </div>
  );
}

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(<App />);
