import { useEffect, useState } from "react";
import { Box, Typography } from "@mui/material";
import { Bar } from "react-chartjs-2";
import {
  Chart,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from "chart.js";
import { GenerateStatistics } from "../wailsjs/go/service/DataService";
import { useTranslation } from "react-i18next";

Chart.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

export default function ReportPanel({ projectId }) {
  const [stats, setStats] = useState(null);
  const { t } = useTranslation();

  useEffect(() => {
    const load = async () => {
      try {
        const s = await GenerateStatistics(projectId, 2025);
        setStats(s);
      } catch (err) {
        console.error(err);
      }
    };
    load();
  }, [projectId]);

  if (!stats) return null;

  const data = {
    labels: [t("reports.avgIncome"), t("reports.avgExpense")],
    datasets: [
      {
        label: t("reports.average"),
        data: [stats.averageIncome, stats.averageExpense],
        backgroundColor: "#1976d2",
      },
      {
        label: t("reports.median"),
        data: [stats.medianIncome, stats.medianExpense],
        backgroundColor: "#9c27b0",
      },
      {
        label: t("reports.stdDev"),
        data: [stats.stdDevIncome, stats.stdDevExpense],
        backgroundColor: "#ff9800",
      },
    ],
  };

  return (
    <Box>
      <Bar data={data} />
      <Typography sx={{ mt: 2 }}>
        {t("reports.trend", { value: stats.trend.toFixed(2) })}
      </Typography>
    </Box>
  );
}
