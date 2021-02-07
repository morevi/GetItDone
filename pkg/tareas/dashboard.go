
package tareas

// DASHBOARD
// Main object. Allows working with projects

type Dashboard struct {
  ID int
  Projects []Project
}

// Initialize a empty dashboard
func (d *Dashboard) New(id int) {}

// Get all projects
func (d *Dashboard) DashboardGetAll() []Project {
  var out []Project
  return out
}

// Get project searching for each tag separately
func (d *Dashboard) SearchByTags(tags string) []Project {
  var out []Project
  return out
}

// Tries to get projects that have coincidences in tasks
func (d *Dashboard) SearchByTasks(search string) []Project {
  var out []Project
  return out
}

// Add a given project the dashboard
func (d *Dashboard) Add(p Project) {}

// Delete a project from the dashboard
func (d *Dashboard) Remove() {}

// Replace a project from the dashboard
func (d *Dashboard) Update(p Project) {}
