
package tareas

// DASHBOARD
// Main object. Allows working with projects

import (
  "strings"
  "reflect"
)

type Dashboard struct {
  ID int
  Projects []Project
}

// Initialize a empty dashboard
func (d *Dashboard) New(id int) {}

func (d *Dashboard) Get(pos int) Project {
  return d.Projects[pos]
}

// Get all projects
func (d *Dashboard) GetAll() []Project {
  var out []Project
  return out
}

// Get project searching for each tag separately
func (d *Dashboard) SearchByTags(tags []string) []Project {
  var out = make([]Project, 0)
  for _, p := range d.Projects {
    for _, t := range tags {
      _, found := FindStringInSlice(p.Tags, t)
      if found {
        out = append(out, p)
        break
      }
    }
  }
  return out
}

// Tries to get projects that have coincidences in tasks
func (d *Dashboard) SearchByTasks(search string) []Project {
  var out []Project
  for _, pro := range(d.Projects) {
    for _, task := range(pro.Items) {
      if strings.Contains(search, task.Content) {
        is_in := false
        for _, p := range(out) {
          if reflect.DeepEqual(p, pro) {
            is_in = true
            break
          }
        }
        if !is_in {
          out = append(out, pro)
        }
      }
    }
  }
  return out
}

// Add a given project the dashboard
func (d *Dashboard) Add(p Project) {
  d.Projects = append(d.Projects, p)
}

// Delete a project from the dashboard
func (d *Dashboard) Remove(i int) {
   d.Projects[i] = d.Projects[len(d.Projects)-1]
   d.Projects = d.Projects[:len(d.Projects)-1]
}

// Replace a project from the dashboard
func (d *Dashboard) Update(i int, p Project) {
  d.Projects[i] = p
}

func FindStringInSlice(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}
