import { useState } from "react";
import { Box, Button, Typography } from "@mui/material";
import { CalculateProjectTaxes } from "../wailsjs/go/service/DataService";

export default function TaxPanel() {
  const [taxes, setTaxes] = useState(null);
  const [error, setError] = useState("");

  const handleCalculate = async () => {
    try {
      const result = await CalculateProjectTaxes(1);
      setTaxes(result);
      setError("");
    } catch (err) {
      setError(err.message || "Fehler bei Berechnung");
    }
  };

  return (
    <Box>
      <Button variant="contained" color="secondary" onClick={handleCalculate}>
        Steuern berechnen
      </Button>
      {taxes && (
        <Box sx={{ mt: 2 }}>
          <Typography>Einnahmen: {taxes.revenue.toFixed(2)} €</Typography>
          <Typography>Ausgaben: {taxes.expenses.toFixed(2)} €</Typography>
          <Typography>Steuerpflichtiges Einkommen: {taxes.taxableIncome.toFixed(2)} €</Typography>
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
