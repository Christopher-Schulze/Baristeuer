import { useState } from "react";
import { Box, TextField, Button, Typography } from "@mui/material";
import { useTranslation } from "react-i18next";
import { Login, Register } from "../wailsjs/go/service/DataService";

export default function LoginPanel({ onLoggedIn }) {
  const [user, setUser] = useState("");
  const [pass, setPass] = useState("");
  const [error, setError] = useState("");
  const { t } = useTranslation();

  const handle = async (fn) => {
    try {
      await fn(user, pass);
      setError("");
      onLoggedIn && onLoggedIn();
    } catch (e) {
      setError(e.message || t("login.error"));
    }
  };

  return (
    <Box sx={{ mt: 4 }}>
      <Box display="flex" gap={2} mb={2}>
        <TextField
          label={t("login.username")}
          value={user}
          onChange={(e) => setUser(e.target.value)}
        />
        <TextField
          label={t("login.password")}
          type="password"
          value={pass}
          onChange={(e) => setPass(e.target.value)}
        />
      </Box>
      <Button variant="contained" onClick={() => handle(Login)} sx={{ mr: 2 }}>
        {t("login.login")}
      </Button>
      <Button variant="outlined" onClick={() => handle(Register)}>
        {t("login.register")}
      </Button>
      {error && (
        <Typography color="error" sx={{ mt: 2 }}>
          {error}
        </Typography>
      )}
    </Box>
  );
}
