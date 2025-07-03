import { useState, useEffect } from "react";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import {
  CssBaseline,
  Container,
  Typography,
  Button,
  TextField,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Paper,
  Switch,
  FormControlLabel,
  Box,
  AppBar,
  Toolbar,
  Tabs,
  Tab,
  Grid,
  Card,
  CardContent,
} from "@mui/material";
import { AddExpense, ListExpenses } from "./wailsjs/go/service/DataService";
import {
  GenerateReport,
  GenerateKSt1,
  GenerateAnlageGem,
  GenerateAnlageGK,
  GenerateKSt1F,
  GenerateAnlageSport,
  GenerateAllForms,
} from "./wailsjs/go/pdf/Generator";

export default function App() {
  const [expenses, setExpenses] = useState([]);
  const [description, setDescription] = useState("");
  const [amount, setAmount] = useState("");
  const [error, setError] = useState("");
  const [darkMode, setDarkMode] = useState(false);
  const [tab, setTab] = useState(0);

  const theme = createTheme({
    palette: {
      mode: darkMode ? "dark" : "light",
      primary: {
        main: "#1976d2",
      },
      secondary: {
        main: "#9c27b0",
      },
    },
  });

  const fetchExpenses = async () => {
    try {
      const list = await ListExpenses(1);
      setExpenses(list || []);
    } catch (err) {
      setError(err.message || "Fehler beim Abrufen der Ausgaben");
    }
  };

  useEffect(() => {
    fetchExpenses();
  }, []);

  const handleAdd = async (e) => {
    e.preventDefault();
    if (!description || !amount) {
      setError("Beschreibung und Betrag erforderlich");
      return;
    }
    try {
      await AddExpense(1, description, parseFloat(amount));
      setDescription("");
      setAmount("");
      setError("");
      await fetchExpenses();
    } catch (err) {
      setError(err.message || "Fehler beim Hinzufügen");
    }
  };

  const handleGenerate = async (fn) => {
    try {
      await fn(1);
      setError("");
    } catch (err) {
      setError(err.message || "Fehler beim Erzeugen");
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
            control={
              <Switch
                checked={darkMode}
                onChange={() => setDarkMode(!darkMode)}
                color="default"
              />
            }
            label={darkMode ? "Dunkel" : "Hell"}
          />
        </Toolbar>
        <Tabs
          value={tab}
          onChange={(_, v) => setTab(v)}
          textColor="inherit"
          indicatorColor="secondary"
          centered
        >
          <Tab label="Ausgaben" />
          <Tab label="Formulare" />
        </Tabs>
      </AppBar>
      <Container maxWidth="md" sx={{ py: 4 }}>
        {tab === 0 && (
          <>
            <Paper sx={{ p: 3, mb: 4 }}>
              <Typography variant="h6" component="h2" gutterBottom>
                Neue Ausgabe
              </Typography>
              <Box
                component="form"
                onSubmit={handleAdd}
                display="flex"
                gap={2}
                flexWrap="wrap"
              >
                <TextField
                  label="Beschreibung"
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                  fullWidth
                />
                <TextField
                  label="Betrag (€)"
                  type="number"
                  value={amount}
                  onChange={(e) => setAmount(e.target.value)}
                />
                <Button type="submit" variant="contained">
                  Hinzufügen
                </Button>
              </Box>
              {error && (
                <Typography color="error" sx={{ mt: 2 }}>
                  {error}
                </Typography>
              )}
            </Paper>
            <Paper>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Beschreibung</TableCell>
                    <TableCell align="right">Betrag (€)</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {expenses.length > 0 ? (
                    expenses.map((e, idx) => (
                      <TableRow key={idx} hover>
                        <TableCell>{e.description}</TableCell>
                        <TableCell align="right">
                          {e.amount.toFixed(2)}
                        </TableCell>
                      </TableRow>
                    ))
                  ) : (
                    <TableRow>
                      <TableCell colSpan={2} align="center">
                        Keine Ausgaben vorhanden
                      </TableCell>
                    </TableRow>
                  )}
                </TableBody>
              </Table>
            </Paper>
          </>
        )}
        {tab === 1 && (
          <Grid container spacing={2}>
            <Grid item xs={12}>
              <Button
                fullWidth
                variant="contained"
                color="secondary"
                onClick={() => handleGenerate(GenerateAllForms)}
              >
                Alle Formulare erstellen
              </Button>
            </Grid>
            <Grid item xs={12} sm={6}>
              <Card>
                <CardContent>
                  <Typography gutterBottom>KSt 1</Typography>
                  <Button
                    variant="outlined"
                    onClick={() => handleGenerate(GenerateKSt1)}
                  >
                    Erstellen
                  </Button>
                </CardContent>
              </Card>
            </Grid>
            <Grid item xs={12} sm={6}>
              <Card>
                <CardContent>
                  <Typography gutterBottom>Anlage Gem</Typography>
                  <Button
                    variant="outlined"
                    onClick={() => handleGenerate(GenerateAnlageGem)}
                  >
                    Erstellen
                  </Button>
                </CardContent>
              </Card>
            </Grid>
            <Grid item xs={12} sm={6}>
              <Card>
                <CardContent>
                  <Typography gutterBottom>Anlage GK</Typography>
                  <Button
                    variant="outlined"
                    onClick={() => handleGenerate(GenerateAnlageGK)}
                  >
                    Erstellen
                  </Button>
                </CardContent>
              </Card>
            </Grid>
            <Grid item xs={12} sm={6}>
              <Card>
                <CardContent>
                  <Typography gutterBottom>KSt 1F</Typography>
                  <Button
                    variant="outlined"
                    onClick={() => handleGenerate(GenerateKSt1F)}
                  >
                    Erstellen
                  </Button>
                </CardContent>
              </Card>
            </Grid>
            <Grid item xs={12} sm={6}>
              <Card>
                <CardContent>
                  <Typography gutterBottom>Anlage Sport</Typography>
                  <Button
                    variant="outlined"
                    onClick={() => handleGenerate(GenerateAnlageSport)}
                  >
                    Erstellen
                  </Button>
                </CardContent>
              </Card>
            </Grid>
          </Grid>
        )}
      </Container>
    </ThemeProvider>
  );
}
