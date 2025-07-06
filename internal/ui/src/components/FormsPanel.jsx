import {
  Grid,
  Card,
  CardContent,
  Typography,
  Button,
  Snackbar,
  Alert,
} from "@mui/material";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import {
  GenerateAllForms,
  GenerateDetailedReport,
  GenerateAnlageGem,
  GenerateAnlageGK,
  GenerateAnlageSport,
  GenerateKSt1,
  GenerateKSt1F,
} from "../wailsjs/go/pdf/Generator";


export default function FormsPanel({ projectId, onGenerated }) {
  const [error, setError] = useState("");
  const { t } = useTranslation();

  const handleGenerate = async (fn) => {
    try {
      const result = await fn(projectId);
      const path = Array.isArray(result) ? result[result.length - 1] : result;
      onGenerated && onGenerated(path);
    } catch (err) {
      setError(err.message || t('forms.error'));
    }
  };

  return (
    <>
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <Button
            fullWidth
            variant="contained"
            color="secondary"
            onClick={() => handleGenerate(GenerateAllForms)}
          >
            {t('forms.generate_all')}
          </Button>
        </Grid>
        <Grid item xs={12} sm={6}>
          <Card>
            <CardContent>
              <Typography gutterBottom>{t('form.detailedReport')}</Typography>
              <Button
                variant="outlined"
                onClick={() => handleGenerate(GenerateDetailedReport)}
              >
                {t('form.generate')}
              </Button>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} sm={6}>
          <Card>
            <CardContent>
              <Typography gutterBottom>{t('form.kst1')}</Typography>
              <Button
                variant="outlined"
                onClick={() => handleGenerate(GenerateKSt1)}
              >
                {t('form.generate')}
              </Button>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} sm={6}>
          <Card>
            <CardContent>
              <Typography gutterBottom>{t('form.anlageGem')}</Typography>
              <Button
                variant="outlined"
                onClick={() => handleGenerate(GenerateAnlageGem)}
              >
                {t('form.generate')}
              </Button>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} sm={6}>
          <Card>
            <CardContent>
              <Typography gutterBottom>{t('form.anlageGK')}</Typography>
              <Button
                variant="outlined"
                onClick={() => handleGenerate(GenerateAnlageGK)}
              >
                {t('form.generate')}
              </Button>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} sm={6}>
          <Card>
            <CardContent>
              <Typography gutterBottom>{t('form.kst1f')}</Typography>
              <Button
                variant="outlined"
                onClick={() => handleGenerate(GenerateKSt1F)}
              >
                {t('form.generate')}
              </Button>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} sm={6}>
          <Card>
            <CardContent>
              <Typography gutterBottom>{t('form.anlageSport')}</Typography>
              <Button
                variant="outlined"
                onClick={() => handleGenerate(GenerateAnlageSport)}
              >
                {t('form.generate')}
              </Button>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
      <Snackbar
        open={!!error}
        autoHideDuration={6000}
        onClose={() => setError("")}
      >
        <Alert severity="error" onClose={() => setError("")}>
          {error}
        </Alert>
      </Snackbar>
    </>
  );
}
