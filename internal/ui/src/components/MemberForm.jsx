import { useState } from "react";
import { Box, TextField, Button, Snackbar, Alert } from "@mui/material";
import { useTranslation } from "react-i18next";

export default function MemberForm({ onSubmit, editItem }) {
  const { t } = useTranslation();
  const [name, setName] = useState(editItem ? editItem.name : "");
  const [email, setEmail] = useState(editItem ? editItem.email : "");
  const [joinDate, setJoinDate] = useState(editItem ? editItem.joinDate : "2024-01-01");
  const [error, setError] = useState("");
  const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  const datePattern = /^\d{4}-\d{2}-\d{2}$/;

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!name || !email) {
      setError(t('errors.member_required'));
      return;
    }
    if (!emailPattern.test(email)) {
      setError(t('errors.member_invalid_email'));
      return;
    }
    if (!datePattern.test(joinDate)) {
      setError(t('errors.member_invalid_date'));
      return;
    }
    await onSubmit(name, email, joinDate, setError);
    if (!editItem) {
      setName("");
      setEmail("");
    }
  };

  return (
    <Box component="form" onSubmit={handleSubmit} display="flex" gap={2} flexWrap="wrap">
      <TextField label={t('member.name')} value={name} onChange={(e) => setName(e.target.value)} fullWidth />
      <TextField label={t('member.email')} value={email} onChange={(e) => setEmail(e.target.value)} fullWidth />
      <TextField
        label={t('member.joinDate')}
        type="date"
        value={joinDate}
        onChange={(e) => setJoinDate(e.target.value)}
        InputLabelProps={{ shrink: true }}
      />
      <Button type="submit" variant="contained">
        {t('member.add')}
      </Button>
      <Snackbar open={!!error} autoHideDuration={6000} onClose={() => setError('')}>
        <Alert severity="error" onClose={() => setError('')}>
          {error}
        </Alert>
      </Snackbar>
    </Box>
  );
}
