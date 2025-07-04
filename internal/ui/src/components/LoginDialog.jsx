import { useState } from "react";
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Button,
  Typography,
} from "@mui/material";
import { AuthenticateUser, CreateUser } from "../wailsjs/go/service/DataService";

export default function LoginDialog({ open, onSuccess }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [register, setRegister] = useState(false);
  const [error, setError] = useState("");

  const handleSubmit = async () => {
    try {
      if (register) {
        await CreateUser(username, password);
      }
      await AuthenticateUser(username, password);
      setError("");
      onSuccess && onSuccess();
    } catch (err) {
      setError(err.message || "Fehler");
    }
  };

  return (
    <Dialog open={open}>
      <DialogTitle>{register ? "Registrieren" : "Anmelden"}</DialogTitle>
      <DialogContent sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
        <TextField
          label="Benutzername"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <TextField
          label="Passwort"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        {error && <Typography color="error">{error}</Typography>}
      </DialogContent>
      <DialogActions>
        <Button onClick={() => setRegister(!register)}>
          {register ? "Login" : "Registrieren"}
        </Button>
        <Button onClick={handleSubmit}>OK</Button>
      </DialogActions>
    </Dialog>
  );
}
