import { useState } from "react";
import { Box, Button, Typography, MenuItem, TextField } from "@mui/material";
import { useTranslation } from "react-i18next";
import { CalculateProjectTaxes } from "../wailsjs/go/service/DataService";

export default function TaxPanel({ projectId }) {
  const [taxes, setTaxes] = useState(null);
  const [error, setError] = useState("");
  const [year, setYear] = useState(2025);
  const { t } = useTranslation();

  const handleCalculate = async () => {
    try {
      const result = await CalculateProjectTaxes(projectId, year);
      setTaxes(result);
      setError("");
    } catch (err) {
      setError(err.message || t('tax.error'));
    }
  };

  return (
    <Box>
      <TextField
        select
        label={t('tax.year') || "Steuerjahr"}
        value={year}
        onChange={(e) => setYear(parseInt(e.target.value))}
        sx={{ mr: 2, minWidth: 120 }}
      >
        <MenuItem value={2025}>2025</MenuItem>
        <MenuItem value={2026}>2026</MenuItem>
      </TextField>
      <Button variant="contained" color="secondary" onClick={handleCalculate}>
        {t('tax.calculate')}
      </Button>
      {taxes && (
        <Box sx={{ mt: 2 }}>
          <Typography>
            {t('tax.revenue', { value: taxes.revenue.toFixed(2) }) || `Einnahmen: ${taxes.revenue.toFixed(2)} €`}
          </Typography>
          <Typography>
            {t('tax.expenses', { value: taxes.expenses.toFixed(2) }) || `Ausgaben: ${taxes.expenses.toFixed(2)} €`}
          </Typography>
          <Typography>
            {t('tax.taxableIncome', { value: taxes.taxableIncome.toFixed(2) }) || `Steuerpflichtiges Einkommen: ${taxes.taxableIncome.toFixed(2)} €`}
          </Typography>
          <Typography>
            {t('tax.totalTax', { value: taxes.totalTax.toFixed(2) }) || `Gesamtsteuer: ${taxes.totalTax.toFixed(2)} €`}
          </Typography>
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