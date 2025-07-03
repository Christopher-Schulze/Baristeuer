import { useState, useEffect } from 'react';
import { useState, useEffect } from 'react';
import { AddIncome, ListExpenses } from './wailsjs/go/data/DataService';
import './index.css';

function App() {
  const [expenses, setExpenses] = useState([]);
  const [description, setDescription] = useState('');
  const [amount, setAmount] = useState('');
  const [error, setError] = useState('');
  const [mode, setMode] = useState('light');

  const toggleMode = () => {
    setMode(prev => (prev === 'light' ? 'dark' : 'light'));
  };

  const fetchExpenses = async () => {
    try {
      const expenseList = await ListExpenses();
      setExpenses(expenseList || []);
    } catch (err) {
      setError(err.message || 'Failed to fetch expenses');
    }
  };

  useEffect(() => {
    fetchExpenses();
  }, []);

  const handleAddIncome = async (e) => {
    e.preventDefault();
    if (!description || !amount) {
      setError('Description and amount are required.');
      return;
    }
    try {
      await AddIncome(description, parseFloat(amount));
      setDescription('');
      setAmount('');
      setError('');
      await fetchExpenses(); // Refresh the list
    } catch (err) {
      setError(err.message || 'Failed to add income');
    }
  };

  return (
    <div className={`min-h-screen ${mode === 'dark' ? 'bg-gray-900 text-white' : 'bg-gray-100 text-black'} p-8 transition-colors duration-300`}>
      <div className="max-w-4xl mx-auto">
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-4xl font-bold">Baristeuer</h1>
          <button
            onClick={toggleMode}
            className={`px-4 py-2 rounded font-semibold ${mode === 'dark' ? 'bg-gray-700 hover:bg-gray-600' : 'bg-blue-500 hover:bg-blue-700 text-white'}`}
          >
            Switch to {mode === 'light' ? 'Dark' : 'Light'} Mode
          </button>
        </div>
        
        <div className={`p-6 rounded-lg shadow-md ${mode === 'dark' ? 'bg-gray-800' : 'bg-white'}`}>
          <h2 className="text-2xl font-semibold mb-4">Add New Income</h2>
          <form onSubmit={handleAddIncome} className="mb-6">
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div>
                <label htmlFor="description" className="block text-sm font-medium mb-1">Description</label>
                <input
                  id="description"
                  type="text"
                  value={description}
                  onChange={e => setDescription(e.target.value)}
                  className={`w-full p-2 border rounded ${mode === 'dark' ? 'bg-gray-700 border-gray-600' : 'border-gray-300'}`}
                  placeholder="e.g., Coffee beans"
                />
              </div>
              <div>
                <label htmlFor="amount" className="block text-sm font-medium mb-1">Amount (€)</label>
                <input
                  id="amount"
                  type="number"
                  value={amount}
                  onChange={e => setAmount(e.target.value)}
                  className={`w-full p-2 border rounded ${mode === 'dark' ? 'bg-gray-700 border-gray-600' : 'border-gray-300'}`}
                  placeholder="e.g., 25.50"
                />
              </div>
              <div className="self-end">
                <button
                  type="submit"
                  className="w-full px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
                >
                  Add Income
                </button>
              </div>
            </div>
          </form>

          {error && (
            <div className={`mb-6 p-4 rounded ${mode === 'dark' ? 'bg-red-900' : 'bg-red-100'}`}>
              <p className="font-semibold text-red-700">Error</p>
              <p className="text-sm text-red-600">{error}</p>
            </div>
          )}

          <h2 className="text-2xl font-semibold mb-4">Expenses</h2>
          <div className="overflow-x-auto">
            <table className={`min-w-full border-collapse ${mode === 'dark' ? 'bg-gray-800' : 'bg-white'}`}>
              <thead>
                <tr>
                  <th className={`px-4 py-2 border-b-2 ${mode === 'dark' ? 'border-gray-700' : 'border-gray-300'} text-left`}>Description</th>
                  <th className={`px-4 py-2 border-b-2 ${mode === 'dark' ? 'border-gray-700' : 'border-gray-300'} text-right`}>Amount (€)</th>
                </tr>
              </thead>
              <tbody>
                {expenses.length > 0 ? (
                  expenses.map((expense, index) => (
                    <tr key={index} className={`${mode === 'dark' ? 'hover:bg-gray-700' : 'hover:bg-gray-50'}`}>
                      <td className={`px-4 py-2 border-b ${mode === 'dark' ? 'border-gray-700' : 'border-gray-200'}`}>{expense.description}</td>
                      <td className={`px-4 py-2 border-b ${mode === 'dark' ? 'border-gray-700' : 'border-gray-200'} text-right`}>{expense.amount.toFixed(2)}</td>
                    </tr>
                  ))
                ) : (
                  <tr>
                    <td colSpan="2" className={`px-4 py-4 text-center ${mode === 'dark' ? 'text-gray-400' : 'text-gray-500'}`}>No expenses recorded yet.</td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;