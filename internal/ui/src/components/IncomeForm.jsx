import { useState } from "react";
import { Box, TextField, Button, Typography } from "@mui/material";
import { useTranslation } from "react-i18next";

export default function IncomeForm({ onSubmit, editItem }) {
  const [source, setSource] = useState(editItem ? editItem.source : "");
  const [amount, setAmount] = useState(editItem ? String(editItem.amount) : "");
  const [error, setError] = useState("");
  const { t } = useTranslation();

  const handleSubmit = async (e) => {
    e.preventDefault();
    const value = parseFloat(amount);
    if (!source || !amount) {
      setError(t('income.required'));
      return;
    }
    if (Number.isNaN(value) || value <= 0) {
      setError(t('income.positive'));
      return;
    }
    await onSubmit(source, value, setError);
    if (!editItem) {
      setSource("");
      setAmount("");
    }
  };

  return (
    <Box component="form" onSubmit={handleSubmit} display="flex" gap={2} flexWrap="wrap">
      <TextField label={t('income.source')} value={source} onChange={(e) => setSource(e.target.value)} fullWidth />
      <TextField label={t('income.amount')} type="number" value={amount} onChange={(e) => setAmount(e.target.value)} />
      <Button type="submit" variant="contained">
        {t('income.add')}
      </Button>
      {error && (
        <Typography color="error" sx={{ mt: 2 }}>
          {error}
        </Typography>
      )}
    </Box>
  );
}
