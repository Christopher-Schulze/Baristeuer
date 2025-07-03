import { Grid, Card, CardContent, Typography, Button } from "@mui/material";
import {
  GenerateAllForms,
  GenerateAnlageGem,
  GenerateAnlageGK,
  GenerateAnlageSport,
  GenerateKSt1,
  GenerateKSt1F,
} from "../wailsjs/go/pdf/Generator";

export default function FormsPanel() {
  const handleGenerate = async (fn) => {
    try {
      await fn(1);
    } catch (err) {
      // TODO: error handling could be improved
    }
  };

  return (
    <Grid container spacing={2}>
      <Grid item xs={12}>
        <Button fullWidth variant="contained" color="secondary" onClick={() => handleGenerate(GenerateAllForms)}>
          Alle Formulare erstellen
        </Button>
      </Grid>
      <Grid item xs={12} sm={6}>
        <Card>
          <CardContent>
            <Typography gutterBottom>KSt 1</Typography>
            <Button variant="outlined" onClick={() => handleGenerate(GenerateKSt1)}>
              Erstellen
            </Button>
          </CardContent>
        </Card>
      </Grid>
      <Grid item xs={12} sm={6}>
        <Card>
          <CardContent>
            <Typography gutterBottom>Anlage Gem</Typography>
            <Button variant="outlined" onClick={() => handleGenerate(GenerateAnlageGem)}>
              Erstellen
            </Button>
          </CardContent>
        </Card>
      </Grid>
      <Grid item xs={12} sm={6}>
        <Card>
          <CardContent>
            <Typography gutterBottom>Anlage GK</Typography>
            <Button variant="outlined" onClick={() => handleGenerate(GenerateAnlageGK)}>
              Erstellen
            </Button>
          </CardContent>
        </Card>
      </Grid>
      <Grid item xs={12} sm={6}>
        <Card>
          <CardContent>
            <Typography gutterBottom>KSt 1F</Typography>
            <Button variant="outlined" onClick={() => handleGenerate(GenerateKSt1F)}>
              Erstellen
            </Button>
          </CardContent>
        </Card>
      </Grid>
      <Grid item xs={12} sm={6}>
        <Card>
          <CardContent>
            <Typography gutterBottom>Anlage Sport</Typography>
            <Button variant="outlined" onClick={() => handleGenerate(GenerateAnlageSport)}>
              Erstellen
            </Button>
          </CardContent>
        </Card>
      </Grid>
    </Grid>
  );
}
