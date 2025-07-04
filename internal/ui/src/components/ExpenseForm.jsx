import { useState } from "react";
import { Box, TextField, Button, Typography } from "@mui/material";
import { useTranslation } from "react-i18next";

export default function ExpenseForm({ onSubmit, editItem }) {
  const [description, setDescription] = useState(editItem ? editItem.description : "");
  const [amount, setAmount] = useState(editItem ? String(editItem.amount) : "");
  const [error, setError] = useState("");
  const { t } = useTranslation();

  const handleSubmit = async (e) => {
    e.preventDefault();
    const value = parseFloat(amount);
    if (!description || !amount) {
      setError(t('expense.required'));
      return;
    }
    if (Number.isNaN(value) || value <= 0) {
      setError(t('expense.positive'));
      return;
    }
    await onSubmit(description, value, setError);
    if (!editItem) {
      setDescription("");
      setAmount("");
    }
  };

  return (
    <Box component="form" onSubmit={handleSubmit} display="flex" gap={2} flexWrap="wrap">
      <TextField label={t('expense.description')} value={description} onChange={(e) => setDescription(e.target.value)} fullWidth />
      <TextField label={t('expense.amount')} type="number" value={amount} onChange={(e) => setAmount(e.target.value)} />
      <Button type="submit" variant="contained">
        {t('expense.add')}
      </Button>
      {error && (
        <Typography color="error" sx={{ mt: 2 }}>
          {error}
        </Typography>
      )}
    </Box>
  );
}
