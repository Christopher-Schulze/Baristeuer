// Placeholder backend service using Wails calls via window.backend
export async function loadProjects() {
  try {
    return await window.backend.loadProjects();
  } catch (e) {
    console.error(e);
    return [];
  }
}

export async function saveProject(project) {
  try {
    return await window.backend.saveProject(project);
  } catch (e) {
    console.error(e);
    return project;
  }
}

export async function saveProjectData(projectId, data) {
  try {
    return await window.backend.saveProjectData(projectId, data);
  } catch (e) {
    console.error(e);
  }
}
