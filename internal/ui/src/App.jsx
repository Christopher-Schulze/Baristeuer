import { useMemo, useState } from 'react'
import { createTheme, ThemeProvider, CssBaseline, Button } from '@mui/material'
import { GenerateReport } from './wailsjs/go/pdf/Generator'

function App() {
  const [mode, setMode] = useState('light')
  const theme = useMemo(() => createTheme({ palette: { mode } }), [mode])

  const toggleMode = () => {
    setMode(prev => (prev === 'light' ? 'dark' : 'light'))
  }

  const handleReport = async () => {
    try {
      const path = await GenerateReport('demo')
      console.log('PDF saved at', path)
    } catch (err) {
      console.error(err)
    }
  }

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <div style={{ padding: '1rem' }}>
        <Button variant="contained" onClick={toggleMode}>
          Switch to {mode === 'light' ? 'dark' : 'light'} mode
        </Button>
        <Button variant="outlined" onClick={handleReport} style={{ marginLeft: '1rem' }}>
          Generate PDF Report
        </Button>
      </div>
    </ThemeProvider>
  )
}

export default App
