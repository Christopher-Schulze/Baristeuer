package backend

import "baristeuer/internal/data"

type Backend struct{}

func NewBackend() *Backend { return &Backend{} }

func (b *Backend) LoadProjects() []data.Project {
    return data.LoadProjects()
}

func (b *Backend) SaveProject(p data.Project) data.Project {
    return data.SaveProject(p)
}

func (b *Backend) SaveProjectData(id int, d data.ProjectData) {
    d.ProjectID = id
    data.SaveProjectData(d)
}
