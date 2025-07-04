import { useState } from "react";
import { Box, Button, Typography } from "@mui/material";
import { useTranslation } from "react-i18next";
import { CalculateProjectTaxes } from "../wailsjs/go/service/DataService";

export default function TaxPanel({ projectId }) {
  const [taxes, setTaxes] = useState(null);
  const [error, setError] = useState("");
  const { t } = useTranslation();

  const handleCalculate = async () => {
    try {
      const result = await CalculateProjectTaxes(projectId);
      setTaxes(result);
      setError("");
    } catch (err) {
      setError(err.message || t('tax.error'));
    }
  };

  return (
    <Box>
      <Button variant="contained" color="secondary" onClick={handleCalculate}>
        {t('tax.calculate')}
      </Button>
      {taxes && (
        <Box sx={{ mt: 2 }}>
          <Typography>{t('tax.revenue', { value: taxes.revenue.toFixed(2) })}</Typography>
          <Typography>{t('tax.expenses', { value: taxes.expenses.toFixed(2) })}</Typography>
          <Typography>{t('tax.taxableIncome', { value: taxes.taxableIncome.toFixed(2) })}</Typography>
          <Typography>{t('tax.totalTax', { value: taxes.totalTax.toFixed(2) })}</Typography>
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
