import { useState } from "react";
import { Box, TextField, Button, Typography } from "@mui/material";

export default function ExpenseForm({ onSubmit, editItem }) {
  const [description, setDescription] = useState(editItem ? editItem.description : "");
  const [amount, setAmount] = useState(editItem ? String(editItem.amount) : "");
  const [error, setError] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    const value = parseFloat(amount);
    if (!description || !amount) {
      setError("Beschreibung und Betrag erforderlich");
      return;
    }
    if (Number.isNaN(value) || value <= 0) {
      setError("Betrag muss eine positive Zahl sein");
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
      <TextField label="Beschreibung" value={description} onChange={(e) => setDescription(e.target.value)} fullWidth />
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
