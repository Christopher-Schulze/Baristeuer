import React from 'react';
import { Box, Typography, Button } from '@mui/material';
import { saveProjectData } from '../services/backend';

export default function SummaryView({ data, project, onBack, onReset }) {
  const save = async () => {
    await saveProjectData(project.id, data);
    onReset();
  };

  return (
    <Box>
      <Typography variant="h6" gutterBottom>Summary for {project.name}</Typography>
      <Typography>Income: {data.income}</Typography>
      <Typography>Expense: {data.expense}</Typography>
      <Box sx={{ mt: 2 }}>
        <Button onClick={onBack} sx={{ mr: 1 }}>Back</Button>
        <Button variant="contained" onClick={save}>Save</Button>
      </Box>
    </Box>
  );
}
