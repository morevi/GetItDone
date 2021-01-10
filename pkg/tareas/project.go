
package tareas

// PROJECT
type Project struct {
  Tags []string
  Description string
  Items []Task
}

func (p *Project) New(tags []string, desc, todos []Task) {
  p.Tags = tags
  p.Description = desc
  p.Items = todos
}

func (p *Project) SetDescription(nd string) {
  p.Description = nd
}

func (p *Project) SetTags(newtags string) {
  p.Tags = strings.Fields(newtags)
}

func(p *Project) AddTask(t Task) {
  append(p.Items, t)
}

func(p *Project) RemoveTask(tpos int){
  p.Items = append(p.Items[:tpos], p.Items[tpos:]...)
}

func(p *Project) GetAll() {
  return p.Items
}

func(p *Project) GetDone(done bool) {
  todos := []Task
  for i, t := p.Items {
    if t.Done == done {
      append(todos, t)
    }
  }
  return todos
}
