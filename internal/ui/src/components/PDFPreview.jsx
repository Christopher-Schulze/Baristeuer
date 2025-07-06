import { Box, Typography } from "@mui/material";
import { useTranslation } from "react-i18next";

export default function PDFPreview({ filePath }) {
  const { t } = useTranslation();
  if (!filePath) {
    return <Typography>{t('preview.none') || 'No PDF generated yet'}</Typography>;
  }
  return (
    <Box sx={{ mt: 2 }}>
      <Typography variant="h6" gutterBottom>
        {t('preview.title') || 'PDF Preview'}
      </Typography>
      <iframe
        src={filePath}
        title="PDF Preview"
        style={{ width: '100%', height: '600px', border: 'none' }}
      />
    </Box>
  );
}
