import { useState } from "react";
import { Box, Button, Typography, MenuItem, TextField } from "@mui/material";
import { CalculateProjectTaxes } from "../wailsjs/go/service/DataService";

export default function TaxPanel({ projectId }) {
  const [taxes, setTaxes] = useState(null);
  const [error, setError] = useState("");
  const [year, setYear] = useState(2025);

  const handleCalculate = async () => {
    try {
      const result = await CalculateProjectTaxes(projectId, year);
      setTaxes(result);
      setError("");
    } catch (err) {
      setError(err.message || "Fehler bei Berechnung");
    }
  };

  return (
    <Box>
      <TextField
        select
        label="Steuerjahr"
        value={year}
        onChange={(e) => setYear(parseInt(e.target.value))}
        sx={{ mr: 2, minWidth: 120 }}
      >
        <MenuItem value={2025}>2025</MenuItem>
        <MenuItem value={2026}>2026</MenuItem>
      </TextField>
      <Button variant="contained" color="secondary" onClick={handleCalculate}>
        Steuern berechnen
      </Button>
      {taxes && (
        <Box sx={{ mt: 2 }}>
          <Typography>Einnahmen: {taxes.revenue.toFixed(2)} €</Typography>
          <Typography>Ausgaben: {taxes.expenses.toFixed(2)} €</Typography>
          <Typography>
            Steuerpflichtiges Einkommen: {taxes.taxableIncome.toFixed(2)} €
          </Typography>
          <Typography>Gesamtsteuer: {taxes.totalTax.toFixed(2)} €</Typography>
        </Box>
      )}
      {error && (
        <Typography color="error" sx={{ mt: 2 }}>
          {error}
        </Typography>
      )}
    </Box>
  );
}
