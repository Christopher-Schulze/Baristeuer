import { useState, useEffect, useCallback } from "react";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import { CssBaseline, Container, FormControlLabel, Switch, AppBar, Toolbar, Typography, Tabs, Tab, Paper, Select, MenuItem, Snackbar, Alert } from "@mui/material";
import { useTranslation } from "react-i18next";
import "./i18n";
import { ListExpenses, ListIncomes, AddIncome, UpdateIncome, DeleteIncome, AddExpense, UpdateExpense, DeleteExpense, AddMember, UpdateMember, ListMembers, DeleteMember } from "./wailsjs/go/service/DataService";
import ProjectPanel from "./components/ProjectPanel";
import IncomeForm from "./components/IncomeForm";
import IncomeTable from "./components/IncomeTable";
import ExpenseForm from "./components/ExpenseForm";
import ExpenseTable from "./components/ExpenseTable";
import MemberForm from "./components/MemberForm";
import MemberTable from "./components/MemberTable";
import TaxPanel from "./components/TaxPanel";
import FormsPanel from "./components/FormsPanel";
import SettingsPanel from "./components/SettingsPanel";

export default function App() {
  const [incomes, setIncomes] = useState([]);
  const [expenses, setExpenses] = useState([]);
  const [members, setMembers] = useState([]);
  const [editMember, setEditMember] = useState(null);
  const [editIncome, setEditIncome] = useState(null);
  const [editExpense, setEditExpense] = useState(null);
  const [darkMode, setDarkMode] = useState(false);
  const [tab, setTab] = useState(2);
  const [projectId, setProjectId] = useState(1);
  const { t, i18n } = useTranslation();
  const [language, setLanguage] = useState(i18n.language);
  const [incomeError, setIncomeError] = useState("");
  const [expenseError, setExpenseError] = useState("");
  const [memberError, setMemberError] = useState("");
  const handleError = (err, setter, key = 'errors.add') => {
    setter(err.message || t(key));
  };

  const theme = createTheme({
    palette: {
      mode: darkMode ? "dark" : "light",
      primary: { main: "#1976d2" },
      secondary: { main: "#9c27b0" },
    },
  });

  const fetchExpenses = useCallback(async () => {
    const list = await ListExpenses(projectId);
    setExpenses(list || []);
  }, [projectId]);
  const fetchIncomes = useCallback(async () => {
    const list = await ListIncomes(projectId);
    setIncomes(list || []);
  }, [projectId]);
  const fetchMembers = async () => {
    const list = await ListMembers();
    setMembers(list || []);
  };

  const handleLanguageChange = (e) => {
    const lng = e.target.value;
    i18n.changeLanguage(lng);
    setLanguage(lng);
  };

  useEffect(() => {
    fetchExpenses();
    fetchIncomes();
  }, [fetchExpenses, fetchIncomes]);
  useEffect(() => {
    fetchMembers();
  }, []);

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
      handleError(err, setIncomeError);
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
      handleError(err, setExpenseError);
    }
  };

  const submitMember = async (name, email, joinDate, setError) => {
    try {
      if (editMember) {
        await UpdateMember(editMember.id, name, email, joinDate);
        setEditMember(null);
      } else {
        await AddMember(name, email, joinDate);
      }
      setError("");
      fetchMembers();
    } catch (err) {
      handleError(err, setMemberError);
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
          <Tab label={t('tab.members')} />
          <Tab label={t('tab.incomes')} />
          <Tab label={t('tab.expenses')} />
          <Tab label={t('tab.forms')} />
          <Tab label={t('tab.taxes')} />
          <Tab label={t('tab.settings')} />
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
                {t('member.new')}
              </Typography>
              <MemberForm onSubmit={submitMember} editItem={editMember} />
            </Paper>
            <Paper>
              <MemberTable
                members={members}
                onEdit={(m) => setEditMember(m)}
                onDelete={async (id) => {
                  await DeleteMember(id);
                  fetchMembers();
                }}
              />
            </Paper>
          </>
        )}
        {tab === 2 && (
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
        {tab === 3 && (
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
        {tab === 4 && <FormsPanel projectId={projectId} />}
        {tab === 5 && (
          <Paper sx={{ p: 3 }}>
            <TaxPanel projectId={projectId} />
          </Paper>
        )}
        {tab === 6 && (
          <Paper sx={{ p: 3 }}>
            <SettingsPanel projectId={projectId} />
          </Paper>
        )}
      </Container>
      <Snackbar open={!!incomeError} autoHideDuration={6000} onClose={() => setIncomeError('')}>
        <Alert severity="error" onClose={() => setIncomeError('')}>
          {incomeError}
        </Alert>
      </Snackbar>
      <Snackbar open={!!expenseError} autoHideDuration={6000} onClose={() => setExpenseError('')}>
        <Alert severity="error" onClose={() => setExpenseError('')}>
          {expenseError}
        </Alert>
      </Snackbar>
      <Snackbar open={!!memberError} autoHideDuration={6000} onClose={() => setMemberError('')}>
        <Alert severity="error" onClose={() => setMemberError('')}>
          {memberError}
        </Alert>
      </Snackbar>
    </ThemeProvider>
  );
}
