package tareas

//Class TASK

type Task struct {
  Done bool
  Content string
  Due string
}

func (t *Task) New(done bool, content string, due string) {
  t.Done = done
  t.Content = content
  t.Due = due
}

func (t *Task) ToogleCheck() {
  t.Done = !(t.Done)
}

func (t *Task) SetContent(content string) {
  t.Content = content
}

func (t *Task) SetDue(due string) {
  t.Due = due
}
