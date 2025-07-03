import { useState } from "react";
import { Box, TextField, Button, Typography } from "@mui/material";

export default function IncomeForm({ onSubmit, editItem }) {
  const [source, setSource] = useState(editItem ? editItem.source : "");
  const [amount, setAmount] = useState(editItem ? String(editItem.amount) : "");
  const [error, setError] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    const value = parseFloat(amount);
    if (!source || !amount) {
      setError("Quelle und Betrag erforderlich");
      return;
    }
    if (Number.isNaN(value) || value <= 0) {
      setError("Betrag muss eine positive Zahl sein");
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
      <TextField label="Quelle" value={source} onChange={(e) => setSource(e.target.value)} fullWidth />
      <TextField label="Betrag (€)" type="number" value={amount} onChange={(e) => setAmount(e.target.value)} />
      <Button type="submit" variant="contained">
        Hinzufügen
      </Button>
      {error && (
        <Typography color="error" sx={{ mt: 2 }}>
          {error}
        </Typography>
      )}
    </Box>
  );
}
