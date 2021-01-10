
package tareas

// PROJECT
type Project struct {
  Tags []string
  Description string
  Items []Task
}

func (p *Project) New(tags []string, desc, todos []Task){
  p.Tags = tags
  p.Description = desc
  p.Items = todos
}
