import { useState } from "react";
import {
  Box,
  TextField,
  Button,
  CircularProgress,
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
import { SetTaxYear } from "../wailsjs/go/pdf/Generator";

export default function SettingsPanel({ projectId }) {
  const { t } = useTranslation();
  const [exportPath, setExportPath] = useState("");
  const [restorePath, setRestorePath] = useState("");
  const [csvPath, setCsvPath] = useState("");
  const [level, setLevel] = useState("info");
  const [taxYear, setTaxYear] = useState(2025);
  const [msg, setMsg] = useState("");
  const [loading, setLoading] = useState(false);

  const doExport = async () => {
    try {
      setLoading(true);
      setMsg(t("settings.processing"));
      await ExportDatabase(exportPath);
      setMsg(t("settings.exported"));
    } catch (e) {
      setMsg(String(e));
    }
    setLoading(false);
  };

  const doRestore = async () => {
    try {
      setLoading(true);
      setMsg(t("settings.processing"));
      await RestoreDatabase(restorePath);
      setMsg(t("settings.restored"));
    } catch (e) {
      setMsg(String(e));
    }
    setLoading(false);
  };

  const doExportCSV = async () => {
    try {
      setLoading(true);
      setMsg(t("settings.processing"));
      await ExportProjectCSV(projectId, csvPath);
      setMsg(t("settings.csv_exported"));
    } catch (e) {
      setMsg(String(e));
    }
    setLoading(false);
  };

  const changeLevel = () => {
    SetLogLevel(level);
    setMsg(t("settings.applied"));
  };

  const applyYear = () => {
    SetTaxYear(parseInt(taxYear));
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
      <Box sx={{ display: "flex", gap: 2, alignItems: "center" }}>
        <TextField
          label={t("settings.tax_year") || "Tax Year"}
          type="number"
          value={taxYear}
          onChange={(e) => setTaxYear(e.target.value)}
          size="small"
        />
        <Button variant="outlined" onClick={applyYear}>
          {t("settings.apply")}
        </Button>
      </Box>
      {(msg || loading) && (
        <Box sx={{ display: "flex", alignItems: "center", gap: 1, mt: 2 }}>
          {loading && <CircularProgress size={20} />}
          <Typography color="primary">{msg}</Typography>
        </Box>
      )}
    </Box>
  );
}
