import { useState } from 'react';
import { GenerateReport } from './wailsjs/go/pdf/Generator';
import './index.css';

function App() {
  const [revenue, setRevenue] = useState('');
  const [expenses, setExpenses] = useState('');
  const [filePath, setFilePath] = useState('');
  const [error, setError] = useState('');
  const [mode, setMode] = useState('light');

  const toggleMode = () => {
    setMode(prev => (prev === 'light' ? 'dark' : 'light'));
  };

  const handleReport = async () => {
    setError('');
    setFilePath('');
    try {
      const rev = parseFloat(revenue);
      const exp = parseFloat(expenses);
      if (isNaN(rev) || isNaN(exp)) {
        setError('Please enter valid numbers for revenue and expenses.');
        return;
      }
      const path = await GenerateReport(rev, exp);
      setFilePath(path);
    } catch (err) {
      setError(err.message || 'An unknown error occurred.');
    }
  };

  return (
    <div className={`min-h-screen ${mode === 'dark' ? 'bg-gray-900 text-white' : 'bg-gray-100 text-black'} p-8 transition-colors duration-300`}>
      <div className="max-w-2xl mx-auto">
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-4xl font-bold">Tax Report Generator</h1>
          <button
            onClick={toggleMode}
            className={`px-4 py-2 rounded font-semibold ${mode === 'dark' ? 'bg-gray-700 hover:bg-gray-600' : 'bg-blue-500 hover:bg-blue-700 text-white'}`}
          >
            Switch to {mode === 'light' ? 'Dark' : 'Light'} Mode
          </button>
        </div>
        
        <div className={`p-6 rounded-lg shadow-md ${mode === 'dark' ? 'bg-gray-800' : 'bg-white'}`}>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
            <div>
              <label htmlFor="revenue" className="block text-sm font-medium mb-2">
                Revenue (€)
              </label>
              <input
                type="number"
                id="revenue"
                value={revenue}
                onChange={(e) => setRevenue(e.target.value)}
                className={`w-full p-2 border rounded ${mode === 'dark' ? 'bg-gray-700 border-gray-600' : 'border-gray-300'}`}
                placeholder="e.g., 50000"
              />
            </div>
            <div>
              <label htmlFor="expenses" className="block text-sm font-medium mb-2">
                Expenses (€)
              </label>
              <input
                type="number"
                id="expenses"
                value={expenses}
                onChange={(e) => setExpenses(e.target.value)}
                className={`w-full p-2 border rounded ${mode === 'dark' ? 'bg-gray-700 border-gray-600' : 'border-gray-300'}`}
                placeholder="e.g., 20000"
              />
            </div>
          </div>

          <button
            onClick={handleReport}
            className="w-full bg-green-600 hover:bg-green-700 text-white font-bold py-3 px-4 rounded transition-transform transform hover:scale-105"
          >
            Generate PDF Report
          </button>

          {filePath && (
            <div className={`mt-6 p-4 rounded ${mode === 'dark' ? 'bg-green-900' : 'bg-green-100'}`}>
              <p className="font-semibold text-green-700">Report generated successfully!</p>
              <p className="text-sm">
                Saved at: <code className={`${mode === 'dark' ? 'text-gray-400' : 'text-gray-600'}`}>{filePath}</code>
              </p>
            </div>
          )}

          {error && (
            <div className={`mt-6 p-4 rounded ${mode === 'dark' ? 'bg-red-900' : 'bg-red-100'}`}>
              <p className="font-semibold text-red-700">Error</p>
              <p className="text-sm text-red-600">{error}</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
