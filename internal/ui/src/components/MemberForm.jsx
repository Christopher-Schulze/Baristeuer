import { useState } from "react";
import { Box, TextField, Button, Typography } from "@mui/material";
import { useTranslation } from "react-i18next";

export default function MemberForm({ onSubmit }) {
  const { t } = useTranslation();
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [joinDate, setJoinDate] = useState("2024-01-01");
  const [error, setError] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!name || !email) {
      setError(t('member.required'));
      return;
    }
    await onSubmit(name, email, joinDate, setError);
    setName("");
    setEmail("");
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
      {error && (
        <Typography color="error" sx={{ mt: 2 }}>
          {error}
        </Typography>
      )}
    </Box>
  );
}
