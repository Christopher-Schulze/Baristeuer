import React, { useState } from 'react';
import { Stepper, Step, StepLabel, Button, Box } from '@mui/material';
import ProjectSelector from './ProjectSelector';
import IncomeExpenseForm from './IncomeExpenseForm';
import SummaryView from './SummaryView';

const steps = ['Project', 'Income & Expenses', 'Summary'];

export default function StepperApp() {
  const [activeStep, setActiveStep] = useState(0);
  const [project, setProject] = useState(null);
  const [data, setData] = useState({ income: '', expense: '' });

  const next = () => setActiveStep((s) => s + 1);
  const back = () => setActiveStep((s) => s - 1);

  const handleProjectSelect = (p) => {
    setProject(p);
    next();
  };

  const handleDataSubmit = (values) => {
    setData(values);
    next();
  };

  const handleReset = () => {
    setActiveStep(0);
    setProject(null);
    setData({ income: '', expense: '' });
  };

  return (
    <Box sx={{ width: '100%', p: 2 }}>
      <Stepper activeStep={activeStep} sx={{ mb: 2 }}>
        {steps.map((label) => (
          <Step key={label}>
            <StepLabel>{label}</StepLabel>
          </Step>
        ))}
      </Stepper>
      {activeStep === 0 && (
        <ProjectSelector onSelect={handleProjectSelect} />
      )}
      {activeStep === 1 && (
        <IncomeExpenseForm onSubmit={handleDataSubmit} onBack={back} />
      )}
      {activeStep === 2 && (
        <SummaryView data={data} project={project} onBack={back} onReset={handleReset} />
      )}
    </Box>
  );
}
