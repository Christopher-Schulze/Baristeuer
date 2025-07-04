import { useState, useEffect } from "react";
import {
  Box,
  TextField,
  Button,
  Typography,
  List,
  ListItem,
  ListItemText,
} from "@mui/material";
import { useTranslation } from "react-i18next";
import {
  ListProjects,
  CreateProject,
  DeleteProject,
} from "../wailsjs/go/service/DataService";

export default function ProjectPanel({ activeId, onSelect }) {
  const [projects, setProjects] = useState([]);
  const [name, setName] = useState("");
  const [error, setError] = useState("");
  const { t } = useTranslation();

  const fetchProjects = async () => {
    const list = await ListProjects();
    setProjects(list || []);
  };

  useEffect(() => {
    fetchProjects();
  }, []);

  const handleCreate = async (e) => {
    e.preventDefault();
    if (!name) return;
    try {
      const p = await CreateProject(name);
      setName("");
      setError("");
      fetchProjects();
      onSelect && onSelect(p.id);
    } catch (err) {
      setError(err.message || t('project.error'));
    }
  };

  const handleDelete = async (id) => {
    await DeleteProject(id);
    fetchProjects();
  };

  return (
    <Box>
      <Box component="form" onSubmit={handleCreate} display="flex" gap={2} mb={2}>
        <TextField
          label={t('project.new')}
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <Button type="submit" variant="contained">
          {t('project.create')}
        </Button>
      </Box>
      {error && (
        <Typography color="error" sx={{ mb: 2 }}>
          {error}
        </Typography>
      )}
      <List>
        {projects.map((p) => (
          <ListItem
            key={p.id}
            selected={p.id === activeId}
            onClick={() => onSelect && onSelect(p.id)}
            secondaryAction={
              <Button color="error" onClick={() => handleDelete(p.id)}>
                {t('delete')}
              </Button>
            }
          >
            <ListItemText primary={p.name} />
          </ListItem>
        ))}
      </List>
    </Box>
  );
}
