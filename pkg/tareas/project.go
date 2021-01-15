
package tareas

import (
  "strings"
)

// PROJECT
type Project struct {
  Tags []string
  Description string
  Items []Task
}

func (p *Project) New(tags []string, desc string, todos []Task) {
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

func (p *Project) AddTask(t Task) {
  p.Items = append(p.Items, t)
}

func (p *Project) RemoveTask(tpos int){
  p.Items = append(p.Items[:tpos], p.Items[tpos:]...)
}

func (p *Project) GetAll() []Task {
  return p.Items
}

func (p *Project) Get(pos int) Task{
  return p.Items[pos]
}

func (p *Project) SearchByCompleted(done bool) []Task {
  var tasks []Task
  for _, t := range p.Items {
    if t.Done == done {
      tasks = append(tasks, t)
    }
  }
  return tasks
}
