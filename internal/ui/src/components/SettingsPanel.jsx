import { useState } from "react";
import {
  Box,
  TextField,
  Button,
  Select,
  MenuItem,
  Typography,
} from "@mui/material";
import { useTranslation } from "react-i18next";
import {
  ExportDatabase,
  RestoreDatabase,
  SetLogLevel,
  ExportProjectCSV,
} from "../wailsjs/go/service/DataService";

export default function SettingsPanel({ projectId }) {
  const { t } = useTranslation();
  const [exportPath, setExportPath] = useState("");
  const [restorePath, setRestorePath] = useState("");
  const [csvPath, setCsvPath] = useState("");
  const [level, setLevel] = useState("info");
  const [msg, setMsg] = useState("");

  const doExport = async () => {
    try {
      await ExportDatabase(exportPath);
      setMsg(t("settings.exported"));
    } catch (e) {
      setMsg(String(e));
    }
  };

  const doRestore = async () => {
    try {
      await RestoreDatabase(restorePath);
      setMsg(t("settings.restored"));
    } catch (e) {
      setMsg(String(e));
    }
  };

  const doExportCSV = async () => {
    try {
      await ExportProjectCSV(projectId, csvPath);
      setMsg(t("settings.csv_exported"));
    } catch (e) {
      setMsg(String(e));
    }
  };

  const changeLevel = () => {
    SetLogLevel(level);
    setMsg(t("settings.applied"));
  };

  return (
    <Box sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
      <Typography variant="h6" component="h2" gutterBottom>
        {t("settings.title")}
      </Typography>
      <Box sx={{ display: "flex", gap: 2 }}>
        <TextField
          label={t("settings.export")}
          value={exportPath}
          onChange={(e) => setExportPath(e.target.value)}
          fullWidth
        />
        <Button variant="outlined" onClick={doExport}>
          {t("settings.export")}
        </Button>
      </Box>
      <Box sx={{ display: "flex", gap: 2 }}>
        <TextField
          label={t("settings.csv")}
          value={csvPath}
          onChange={(e) => setCsvPath(e.target.value)}
          fullWidth
        />
        <Button variant="outlined" onClick={doExportCSV}>
          {t("settings.csv")}
        </Button>
      </Box>
      <Box sx={{ display: "flex", gap: 2 }}>
        <TextField
          label={t("settings.restore")}
          value={restorePath}
          onChange={(e) => setRestorePath(e.target.value)}
          fullWidth
        />
        <Button variant="outlined" onClick={doRestore}>
          {t("settings.restore")}
        </Button>
      </Box>
      <Box sx={{ display: "flex", gap: 2, alignItems: "center" }}>
        <Select
          value={level}
          onChange={(e) => setLevel(e.target.value)}
          size="small"
        >
          <MenuItem value="debug">debug</MenuItem>
          <MenuItem value="info">info</MenuItem>
          <MenuItem value="warn">warn</MenuItem>
          <MenuItem value="error">error</MenuItem>
        </Select>
        <Button variant="outlined" onClick={changeLevel}>
          {t("settings.apply")}
        </Button>
      </Box>
      {msg && (
        <Typography color="primary" sx={{ mt: 2 }}>
          {msg}
        </Typography>
      )}
    </Box>
  );
}
