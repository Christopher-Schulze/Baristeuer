import { useState, useEffect } from "react";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import { CssBaseline, Container, FormControlLabel, Switch, AppBar, Toolbar, Typography, Tabs, Tab, Paper } from "@mui/material";
import { ListExpenses, ListIncomes, AddIncome, UpdateIncome, DeleteIncome, AddExpense, UpdateExpense, DeleteExpense } from "./wailsjs/go/service/DataService";
import IncomeForm from "./components/IncomeForm";
import IncomeTable from "./components/IncomeTable";
import ExpenseForm from "./components/ExpenseForm";
import ExpenseTable from "./components/ExpenseTable";
import TaxPanel from "./components/TaxPanel";
import FormsPanel from "./components/FormsPanel";

export default function App() {
  const [incomes, setIncomes] = useState([]);
  const [expenses, setExpenses] = useState([]);
  const [editIncome, setEditIncome] = useState(null);
  const [editExpense, setEditExpense] = useState(null);
  const [darkMode, setDarkMode] = useState(false);
  const [tab, setTab] = useState(0);

  const theme = createTheme({
    palette: {
      mode: darkMode ? "dark" : "light",
      primary: { main: "#1976d2" },
      secondary: { main: "#9c27b0" },
    },
  });

  const fetchExpenses = async () => {
    const list = await ListExpenses(1);
    setExpenses(list || []);
  };
  const fetchIncomes = async () => {
    const list = await ListIncomes(1);
    setIncomes(list || []);
  };

  useEffect(() => {
    fetchExpenses();
    fetchIncomes();
  }, []);

  const submitIncome = async (source, amount, setError) => {
    try {
      if (editIncome) {
        await UpdateIncome(editIncome.id, 1, source, amount);
        setEditIncome(null);
      } else {
        await AddIncome(1, source, amount);
      }
      setError("");
      fetchIncomes();
    } catch (err) {
      setError(err.message || "Fehler beim Hinzufügen");
    }
  };

  const submitExpense = async (desc, amount, setError) => {
    try {
      if (editExpense) {
        await UpdateExpense(editExpense.id, 1, desc, amount);
        setEditExpense(null);
      } else {
        await AddExpense(1, desc, amount);
      }
      setError("");
      fetchExpenses();
    } catch (err) {
      setError(err.message || "Fehler beim Hinzufügen");
    }
  };

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <AppBar position="static" color="primary">
        <Toolbar>
          <Typography variant="h6" sx={{ flexGrow: 1 }}>
            Baristeuer
          </Typography>
          <FormControlLabel
            control={<Switch checked={darkMode} onChange={() => setDarkMode(!darkMode)} color="default" />}
            label={darkMode ? "Dunkel" : "Hell"}
          />
        </Toolbar>
        <Tabs value={tab} onChange={(_, v) => setTab(v)} textColor="inherit" indicatorColor="secondary" centered>
          <Tab label="Einnahmen" />
          <Tab label="Ausgaben" />
          <Tab label="Formulare" />
          <Tab label="Steuern" />
        </Tabs>
      </AppBar>
      <Container maxWidth="md" sx={{ py: 4 }}>
        {tab === 0 && (
          <>
            <Paper sx={{ p: 3, mb: 4 }}>
              <Typography variant="h6" component="h2" gutterBottom>
                Neue Einnahme
              </Typography>
              <IncomeForm onSubmit={submitIncome} editItem={editIncome} />
            </Paper>
            <Paper>
              <IncomeTable
                incomes={incomes}
                onEdit={(i) => setEditIncome(i)}
                onDelete={async (id) => {
                  await DeleteIncome(id);
                  fetchIncomes();
                }}
              />
            </Paper>
          </>
        )}
        {tab === 1 && (
          <>
            <Paper sx={{ p: 3, mb: 4 }}>
              <Typography variant="h6" component="h2" gutterBottom>
                Neue Ausgabe
              </Typography>
              <ExpenseForm onSubmit={submitExpense} editItem={editExpense} />
            </Paper>
            <Paper>
              <ExpenseTable
                expenses={expenses}
                onEdit={(e) => setEditExpense(e)}
                onDelete={async (id) => {
                  await DeleteExpense(id);
                  fetchExpenses();
                }}
              />
            </Paper>
          </>
        )}
        {tab === 2 && <FormsPanel />}
        {tab === 3 && (
          <Paper sx={{ p: 3 }}>
            <TaxPanel />
          </Paper>
        )}
      </Container>
    </ThemeProvider>
  );
}
