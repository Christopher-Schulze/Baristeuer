package data

// ProjectData represents placeholder project information.
type ProjectData struct {
	ID   string
	Name string
	Info string
}

// GetProjectData returns dummy data for a project.
func GetProjectData(projectID string) ProjectData {
	return ProjectData{
		ID:   projectID,
		Name: "Example Project" + " " + projectID,
		Info: "This is placeholder project information used for PDF generation.",
	}
}
