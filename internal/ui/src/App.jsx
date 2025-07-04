import { useState, useEffect } from "react";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import { CssBaseline, Container, FormControlLabel, Switch, AppBar, Toolbar, Typography, Tabs, Tab, Paper } from "@mui/material";
import { ListExpenses, ListIncomes, AddIncome, UpdateIncome, DeleteIncome, AddExpense, UpdateExpense, DeleteExpense } from "./wailsjs/go/service/DataService";
import ProjectPanel from "./components/ProjectPanel";
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
  const [tab, setTab] = useState(1);
  const [projectId, setProjectId] = useState(1);

  const theme = createTheme({
    palette: {
      mode: darkMode ? "dark" : "light",
      primary: { main: "#1976d2" },
      secondary: { main: "#9c27b0" },
    },
  });

  const fetchExpenses = async () => {
    const list = await ListExpenses(projectId);
    setExpenses(list || []);
  };
  const fetchIncomes = async () => {
    const list = await ListIncomes(projectId);
    setIncomes(list || []);
  };

  useEffect(() => {
    fetchExpenses();
    fetchIncomes();
  }, [projectId]);

  const submitIncome = async (source, amount, setError) => {
    try {
      if (editIncome) {
        await UpdateIncome(editIncome.id, projectId, source, amount);
        setEditIncome(null);
      } else {
        await AddIncome(projectId, source, amount);
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
        await UpdateExpense(editExpense.id, projectId, desc, amount);
        setEditExpense(null);
      } else {
        await AddExpense(projectId, desc, amount);
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
          <Tab label="Projekte" />
          <Tab label="Einnahmen" />
          <Tab label="Ausgaben" />
          <Tab label="Formulare" />
          <Tab label="Steuern" />
        </Tabs>
      </AppBar>
      <Container maxWidth="md" sx={{ py: 4 }}>
        {tab === 0 && (
          <ProjectPanel activeId={projectId} onSelect={(id) => setProjectId(id)} />
        )}
        {tab === 1 && (
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
        {tab === 2 && (
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
        {tab === 3 && <FormsPanel projectId={projectId} />}
        {tab === 4 && (
          <Paper sx={{ p: 3 }}>
            <TaxPanel projectId={projectId} />
          </Paper>
        )}
      </Container>
    </ThemeProvider>
  );
}
