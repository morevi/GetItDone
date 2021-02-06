package tareas

//Class TASK
import (
  "time"
)

type Task struct {
  Done bool
  Content string
  Due time.Time
}

func (t *Task) New(done bool, content string, due time.Time) {
  t.Done = donetes
  t.Content = content
  t.Due = due
}

func (t *Task) ToogleCheck() {
  t.Done = !(t.Done)
}

func (t *Task) SetContent(content string) {
  t.Content = content
}

func (t *Task) SetDue(due time.Time) {
  t.Due = due
}
