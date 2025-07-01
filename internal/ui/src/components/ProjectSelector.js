import React, { useState, useEffect } from 'react';
import { Box, TextField, Button, List, ListItem, ListItemButton } from '@mui/material';
import { loadProjects, saveProject } from '../services/backend';

export default function ProjectSelector({ onSelect }) {
  const [projects, setProjects] = useState([]);
  const [name, setName] = useState('');

  useEffect(() => {
    loadProjects().then(setProjects);
  }, []);

  const create = async () => {
    if (!name) return;
    const p = await saveProject({ name });
    setProjects([...projects, p]);
    setName('');
  };

  return (
    <Box>
      <TextField
        label="Project name"
        value={name}
        onChange={(e) => setName(e.target.value)}
        required
        sx={{ mr: 1 }}
      />
      <Button variant="contained" onClick={create} disabled={!name}>Create</Button>
      <List>
        {projects.map((p) => (
          <ListItem key={p.id} disablePadding>
            <ListItemButton onClick={() => onSelect(p)}>{p.name}</ListItemButton>
          </ListItem>
        ))}
      </List>
    </Box>
  );
}
