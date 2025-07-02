import { useMemo, useState } from 'react'
import { createTheme, ThemeProvider, CssBaseline, Button } from '@mui/material'

function App() {
  const [mode, setMode] = useState('light')
  const theme = useMemo(() => createTheme({ palette: { mode } }), [mode])

  const toggleMode = () => {
    setMode(prev => (prev === 'light' ? 'dark' : 'light'))
  }

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <div style={{ padding: '1rem' }}>
        <Button variant="contained" onClick={toggleMode}>
          Switch to {mode === 'light' ? 'dark' : 'light'} mode
        </Button>
      </div>
    </ThemeProvider>
  )
}

export default App
