package data

import (
    "encoding/json"
    "os"
    "path/filepath"
    "sync"
)

type Project struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type ProjectData struct {
    ProjectID int     `json:"projectId"`
    Income    float64 `json:"income"`
    Expense   float64 `json:"expense"`
}

var (
    mu       sync.Mutex
    projects []Project
    nextID   int
    dataMap  = map[int]ProjectData{}
    storeDir = filepath.Join(".", "data")
)

func init() {
    os.MkdirAll(storeDir, 0755)
    load()
}

func load() {
    f, err := os.Open(filepath.Join(storeDir, "projects.json"))
    if err == nil {
        defer f.Close()
        json.NewDecoder(f).Decode(&projects)
        for _, p := range projects {
            if p.ID >= nextID {
                nextID = p.ID + 1
            }
        }
    }
    f2, err := os.Open(filepath.Join(storeDir, "data.json"))
    if err == nil {
        defer f2.Close()
        json.NewDecoder(f2).Decode(&dataMap)
    }
}

func persist() {
    f, _ := os.Create(filepath.Join(storeDir, "projects.json"))
    json.NewEncoder(f).Encode(projects)
    f.Close()
    f2, _ := os.Create(filepath.Join(storeDir, "data.json"))
    json.NewEncoder(f2).Encode(dataMap)
    f2.Close()
}

func LoadProjects() []Project {
    mu.Lock()
    defer mu.Unlock()
    return append([]Project(nil), projects...)
}

func SaveProject(p Project) Project {
    mu.Lock()
    defer mu.Unlock()
    p.ID = nextID
    nextID++
    projects = append(projects, p)
    persist()
    return p
}

func SaveProjectData(d ProjectData) {
    mu.Lock()
    defer mu.Unlock()
    dataMap[d.ProjectID] = d
    persist()
}
