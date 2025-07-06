import { useEffect, useState } from "react";
import { Box } from "@mui/material";
import { Line } from "react-chartjs-2";
import {
  Chart,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from "chart.js";
import { GenerateStatistics } from "../wailsjs/go/service/DataService";
import { useTranslation } from "react-i18next";

Chart.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
);

export default function StatisticsChart({ projectId }) {
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

  const trendLabel = t("reports.trend", { value: "" }).replace(/[:\s]*$/, "");

  const data = {
    labels: [t("reports.avgIncome"), t("reports.avgExpense")],
    datasets: [
      {
        label: t("reports.average"),
        data: [stats.averageIncome, stats.averageExpense],
        borderColor: "#1976d2",
        backgroundColor: "rgba(25,118,210,0.5)",
        tension: 0.3,
      },
      {
        label: t("reports.median"),
        data: [stats.medianIncome, stats.medianExpense],
        borderColor: "#9c27b0",
        backgroundColor: "rgba(156,39,176,0.5)",
        tension: 0.3,
      },
      {
        label: trendLabel,
        data: [stats.trend, stats.trend],
        borderColor: "#ff9800",
        backgroundColor: "rgba(255,152,0,0.5)",
        tension: 0.3,
      },
    ],
  };

  return (
    <Box sx={{ mt: 3 }}>
      <Line data={data} />
    </Box>
  );
}
