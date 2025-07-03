import { useState, useEffect } from 'react';
import { GenerateReport } from './wailsjs/go/pdf/Generator';
import { CreateProject, ListIncomes, AddExpense, AddIncome, ListExpenses } from './wailsjs/go/service/DataService';
import './index.css';

function App() {
  const [revenue, setRevenue] = useState('');
  const [expenses, setExpenses] = useState('');
  const [filePath, setFilePath] = useState('');
  const [error, setError] = useState('');
  const [mode, setMode] = useState('light');
  const [projectName, setProjectName] = useState('');
  const [projectID, setProjectID] = useState(null);
  const [expenseCategory, setExpenseCategory] = useState('');
  const [expenseAmount, setExpenseAmount] = useState('');
  const [incomes, setIncomes] = useState([]);
  const [incomeSource, setIncomeSource] = useState('');
  const [incomeAmount, setIncomeAmount] = useState('');
  const [expensesList, setExpensesList] = useState([]);

  const toggleMode = () => {
    setMode(prev => (prev === 'light' ? 'dark' : 'light'));
  };

  useEffect(() => {
    const fetchIncomes = async () => {
      if (!projectID) return;
      try {
        const list = await ListIncomes(projectID);
        setIncomes(list);
      } catch (err) {
        console.error(err);
      }
    };
    fetchIncomes();
  }, [projectID]);

  useEffect(() => {
    const fetchExpenses = async () => {
      if (!projectID) return;
      try {
        const list = await ListExpenses(projectID);
        setExpensesList(list);
      } catch (err) {
        console.error(err);
      }
    };
    fetchExpenses();
  }, [projectID]);

  const handleCreateProject = async () => {
    try {
      const p = await CreateProject(projectName);
      setProjectID(p.ID);
    } catch (err) {
      setError(err.message || 'Failed to create project');
    }
  };

  const handleAddExpense = async () => {
    if (!projectID) return;
    try {
      await AddExpense(projectID, expenseCategory, parseFloat(expenseAmount));
      setExpenseCategory('');
      setExpenseAmount('');
      const list = await ListIncomes(projectID);
      setIncomes(list);
      const expList = await ListExpenses(projectID);
      setExpensesList(expList);
    } catch (err) {
      setError(err.message || 'Failed to add expense');
    }
  };

  const handleAddIncome = async () => {
    if (!projectID) return;
    try {
      await AddIncome(projectID, incomeSource, parseFloat(incomeAmount));
      setIncomeSource('');
      setIncomeAmount('');
      const incList = await ListIncomes(projectID);
      setIncomes(incList);
      const expList = await ListExpenses(projectID);
      setExpensesList(expList);
    } catch (err) {
      setError(err.message || 'Failed to add income');
    }
  };

  const handleReport = async () => {
    setError('');
    setFilePath('');
    try {
      if (!projectID) {
        setError('Please create a project first.');
        return;
      }
      const path = await GenerateReport(projectID);
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
          <div className="mb-6">
            <label htmlFor="project" className="block text-sm font-medium mb-2">
              Project Name
            </label>
            <input
              id="project"
              value={projectName}
              onChange={e => setProjectName(e.target.value)}
              className={`w-full p-2 border rounded ${mode === 'dark' ? 'bg-gray-700 border-gray-600' : 'border-gray-300'}`}
              placeholder="My Project"
            />
            <button
              onClick={handleCreateProject}
              className="mt-2 px-4 py-2 bg-blue-500 text-white rounded"
            >
              Create Project
            </button>
          </div>

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

          {projectID && (
            <div className="mb-6">
              <h2 className="font-semibold mb-2">Add Expense</h2>
              <input
                type="text"
                value={expenseCategory}
                onChange={e => setExpenseCategory(e.target.value)}
                placeholder="Category"
                className={`w-full p-2 mb-2 border rounded ${mode === 'dark' ? 'bg-gray-700 border-gray-600' : 'border-gray-300'}`}
              />
              <input
                type="number"
                value={expenseAmount}
                onChange={e => setExpenseAmount(e.target.value)}
                placeholder="Amount"
                className={`w-full p-2 mb-2 border rounded ${mode === 'dark' ? 'bg-gray-700 border-gray-600' : 'border-gray-300'}`}
              />
              <button
                onClick={handleAddExpense}
                className="px-4 py-2 bg-purple-500 text-white rounded"
              >
                Add Expense
              </button>
            </div>
          )}

          {projectID && (
            <div className="mb-6">
              <h2 className="font-semibold mb-2">Add Income</h2>
              <input
                type="text"
                value={incomeSource}
                onChange={e => setIncomeSource(e.target.value)}
                placeholder="Source"
                className={`w-full p-2 mb-2 border rounded ${mode === 'dark' ? 'bg-gray-700 border-gray-600' : 'border-gray-300'}`}
              />
              <input
                type="number"
                value={incomeAmount}
                onChange={e => setIncomeAmount(e.target.value)}
                placeholder="Amount"
                className={`w-full p-2 mb-2 border rounded ${mode === 'dark' ? 'bg-gray-700 border-gray-600' : 'border-gray-300'}`}
              />
              <button
                onClick={handleAddIncome}
                className="px-4 py-2 bg-blue-500 text-white rounded"
              >
                Add Income
              </button>
            </div>
          )}

          <button
            onClick={handleReport}
            className="w-full bg-green-600 hover:bg-green-700 text-white font-bold py-3 px-4 rounded transition-transform transform hover:scale-105"
          >
            Generate PDF Report
          </button>

          {projectID && incomes.length > 0 && (
            <div className="mt-6">
              <h2 className="font-semibold mb-2">Incomes</h2>
              <ul className="list-disc list-inside">
                {incomes.map(i => (
                  <li key={i.ID}>{i.Source} - {i.Amount}€</li>
                ))}
              </ul>
            </div>
          )}

          {projectID && expensesList.length > 0 && (
            <div className="mt-6">
              <h2 className="font-semibold mb-2">Expenses</h2>
              <ul className="list-disc list-inside">
                {expensesList.map(e => (
                  <li key={e.ID}>{e.Category} - {e.Amount}€</li>
                ))}
              </ul>
            </div>
          )}

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