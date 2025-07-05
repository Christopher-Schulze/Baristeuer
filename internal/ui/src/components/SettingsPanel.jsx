import { useState } from "react";
import {
  Box,
  TextField,
  Button,
  Select,
  MenuItem,
  Typography,
  Alert,
} from "@mui/material";
import { useTranslation } from "react-i18next";
import {
  ExportDatabase,
  RestoreDatabase,
  SetLogLevel,
  SetLogFormat,
  ExportProjectCSV,
} from "../wailsjs/go/service/DataService";
import { SetTaxYear } from "../wailsjs/go/pdf/Generator";

export default function SettingsPanel({ projectId }) {
  const { t } = useTranslation();
  const [exportPath, setExportPath] = useState("");
  const [restorePath, setRestorePath] = useState("");
  const [csvPath, setCsvPath] = useState("");
  const [level, setLevel] = useState("info");
  const [format, setFormat] = useState("text");
  const [taxYear, setTaxYear] = useState(2025);
  const [feedback, setFeedback] = useState({ type: "", text: "" });

  const doExport = async () => {
    try {
      await ExportDatabase(exportPath);
      setFeedback({ type: "success", text: t("settings.exported") });
    } catch (e) {
      setFeedback({ type: "error", text: String(e) });
    }
  };

  const doRestore = async () => {
    try {
      await RestoreDatabase(restorePath);
      setFeedback({ type: "success", text: t("settings.restored") });
    } catch (e) {
      setFeedback({ type: "error", text: String(e) });
    }
  };

  const doExportCSV = async () => {
    try {
      await ExportProjectCSV(projectId, csvPath);
      setFeedback({ type: "success", text: t("settings.csv_exported") });
    } catch (e) {
      setFeedback({ type: "error", text: String(e) });
    }
  };

  const changeLevel = () => {
    SetLogLevel(level);
    setFeedback({ type: "success", text: t("settings.applied") });
  };

  const changeFormat = () => {
    SetLogFormat(format);
    setFeedback({ type: "success", text: t("settings.applied") });
  };

  const applyYear = () => {
    SetTaxYear(parseInt(taxYear));
    setFeedback({ type: "success", text: t("settings.applied") });
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
        <Select
          value={format}
          onChange={(e) => setFormat(e.target.value)}
          size="small"
        >
          <MenuItem value="text">text</MenuItem>
          <MenuItem value="json">json</MenuItem>
        </Select>
        <Button variant="outlined" onClick={changeFormat}>
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
      {feedback.text && (
        <Alert severity={feedback.type} sx={{ mt: 2 }}>
          {feedback.text}
        </Alert>
      )}
    </Box>
  );
}
