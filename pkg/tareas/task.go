package tareas

//Class TASK
import "time"

type Task struct {
  Done bool
  Content string
  Due time.Time
}

func (t Task) New(done bool, content string, due time.Time) {
  t.Done = done
  t.Content = content
  t.Due = due
}

func (t Task) toogleCheck() {
  t.Done = !(t.Done)
}
