import { useState, useEffect } from "react";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import { CssBaseline, Container, FormControlLabel, Switch, AppBar, Toolbar, Typography, Tabs, Tab, Paper, Select, MenuItem } from "@mui/material";
import { useTranslation } from "react-i18next";
import "./i18n";
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
  const { t, i18n } = useTranslation();
  const [language, setLanguage] = useState(i18n.language);

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

  const handleLanguageChange = (e) => {
    const lng = e.target.value;
    i18n.changeLanguage(lng);
    setLanguage(lng);
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
      setError(err.message || t('add_error'));
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
      setError(err.message || t('add_error'));
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
          <Select
            value={language}
            onChange={handleLanguageChange}
            size="small"
            sx={{ mr: 2, color: 'inherit' }}
            variant="standard"
          >
            <MenuItem value="de">DE</MenuItem>
            <MenuItem value="en">EN</MenuItem>
          </Select>
          <FormControlLabel
            control={<Switch checked={darkMode} onChange={() => setDarkMode(!darkMode)} color="default" />}
            label={darkMode ? t('theme.dark') : t('theme.light')}
          />
        </Toolbar>
        <Tabs value={tab} onChange={(_, v) => setTab(v)} textColor="inherit" indicatorColor="secondary" centered>
          <Tab label={t('tab.projects')} />
          <Tab label={t('tab.incomes')} />
          <Tab label={t('tab.expenses')} />
          <Tab label={t('tab.forms')} />
          <Tab label={t('tab.taxes')} />
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
                {t('income.new')}
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
                {t('expense.new')}
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
