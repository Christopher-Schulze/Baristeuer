import React, { useState } from 'react';
import { Box, TextField, Button } from '@mui/material';

export default function IncomeExpenseForm({ onSubmit, onBack }) {
  const [income, setIncome] = useState('');
  const [expense, setExpense] = useState('');
  const [errors, setErrors] = useState({});

  const validate = () => {
    const errs = {};
    if (!income) errs.income = 'Required';
    if (!expense) errs.expense = 'Required';
    if (income && isNaN(Number(income))) errs.income = 'Must be a number';
    if (expense && isNaN(Number(expense))) errs.expense = 'Must be a number';
    setErrors(errs);
    return Object.keys(errs).length === 0;
  };

  const submit = () => {
    if (!validate()) return;
    onSubmit({ income: Number(income), expense: Number(expense) });
  };

  return (
    <Box>
      <TextField
        label="Income"
        value={income}
        onChange={(e) => setIncome(e.target.value)}
        required
        error={!!errors.income}
        helperText={errors.income}
        sx={{ display: 'block', mb: 2 }}
      />
      <TextField
        label="Expense"
        value={expense}
        onChange={(e) => setExpense(e.target.value)}
        required
        error={!!errors.expense}
        helperText={errors.expense}
        sx={{ display: 'block', mb: 2 }}
      />
      <Box sx={{ mt: 2 }}>
        <Button onClick={onBack} sx={{ mr: 1 }}>Back</Button>
        <Button variant="contained" onClick={submit}>Next</Button>
      </Box>
    </Box>
  );
}
