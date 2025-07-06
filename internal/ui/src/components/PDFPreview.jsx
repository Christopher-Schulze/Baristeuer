import { Box, Typography, Alert, Button } from "@mui/material";
import { useTranslation } from "react-i18next";
import { useState, useEffect } from "react";

export default function PDFPreview({ filePath }) {
  const { t } = useTranslation();
  const [loadError, setLoadError] = useState(false);

  useEffect(() => {
    setLoadError(false);
  }, [filePath]);

  if (!filePath) {
    return <Typography>{t('preview.none') || 'No PDF generated yet'}</Typography>;
  }

  return (
    <Box sx={{ mt: 2 }}>
      <Typography variant="h6" gutterBottom>
        {t('preview.title') || 'PDF Preview'}
      </Typography>
      {loadError && (
        <Alert severity="error" sx={{ mb: 2 }}>
          {t('preview.error') || 'Unable to load PDF'}
        </Alert>
      )}
      <Box sx={{ mb: 1 }}>
        <Button
          variant="outlined"
          onClick={() => window.open(filePath, '_blank')}
        >
          {t('preview.open') || 'Open'}
        </Button>
        <Button
          variant="outlined"
          sx={{ ml: 1 }}
          component="a"
          href={filePath}
          download
        >
          {t('preview.download') || 'Download'}
        </Button>
      </Box>
      <iframe
        src={filePath}
        title="PDF Preview"
        style={{ width: '100%', height: '600px', border: 'none' }}
        onError={() => setLoadError(true)}
      />
    </Box>
  );
}
