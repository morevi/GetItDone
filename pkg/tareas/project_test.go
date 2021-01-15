
package tareas

import (
  "testing"
  "time"
  "reflect"
)

func initializeProject() Project{
  var p Project
  tags := []string{"proyecto", "iv", "test"}
  desc := "test description"

  var t1 Task
  t1.New(false, "test task 1", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
  var t2 Task
  t2.New(false, "test task 2", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

  todos := []Task{t1, t2}

  p.Tags = tags
  p.Description = desc
  p.Items = todos

  return p
}

func eqStringSlice(a []string, b []string) bool {
  if (a == nil) != (b == nil) {
    return false
  }
  if len(a) != len(b) {
    return false
  }
  for i := range a {
    if a[i] != b[i] {
      return false
    }
  }
  return true
}

func eqTaskSlice(a []Task, b []Task) bool {
  if (a == nil) != (b == nil) {
    return false
  }
  if len(a) != len(b) {
    return false
  }
  for i := range a {
    if a[i] != b[i] {
      return false
    }
  }
  return true
}

func TestProjectNew(t *testing.T) {
  t.Log("Test Project contructor")
  want := initializeProject()

  var t1 Task
  t1.New(false, "test task 1", time.Now().Add(24*time.Hour))
  var t2 Task
  t2.New(false, "test task 2", time.Now().Add(48*time.Hour))

  var p Project
  p.New([]string{"proyecto", "iv", "test"}, "test description", []Task{t1, t2})

  if reflect.DeepEqual(p,want) {
    t.Errorf("Test for Project.New() did not pass")
  }
}

func TestSetDescription(t *testing.T) {
  t.Log("Test project set description")
  want := "new project description"

  p := initializeProject()
  p.SetDescription(want)

  if p.Description != want {
    t.Errorf("received %s", p.Description)
  }
}

func TetSetTags(t *testing.T) {
  t.Log("Test project set tags")
  p := initializeProject()

  p.SetTags("a b c")
  want := []string{"a", "b", "c"}

  if eqStringSlice(p.Tags, want) {
    t.Errorf("setTags failed ")
  }
}

func TestAddTask(t *testing.T) {
  t.Log("Test project add task")
  p := initializeProject()
  var task Task
  task.New(false, "test", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
  want := append(p.Items, task)

  p.AddTask(task)

  if len(p.Items) != len(want) {
    t.Log("Not same length")
  }

  if p.Items[2] != want[2] {
    t.Log("Not same task")
  }
}

func TestRemove(t *testing.T) {
  var task Task
  task.New(false, "test task 1", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

  want := []Task{task}

  p := initializeProject()
  p.Remove(1)
  got := p.Items

  if eqTaskSlice(got, want) {
    t.Errorf("\ngot : %+v\nwant: %+v", got, want)
  }

}

func TestGet(t *testing.T) {
  t.Log("Test Project.Get")
  var want Task
  want.New(false, "test task 1", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

  p := initializeProject()
  got := p.Get(0)

  if got != want {
    t.Errorf("\ngot : %+v\nwant: %+v", got, want)
  }
}

func TestGetAll(t *testing.T) {
  t.Log("Test get all tasks in project")

  var t1 Task
  t1.New(false, "test task 1", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
  var t2 Task
  t2.New(false, "test task 2", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

  p := initializeProject()

  want := []Task{t1, t2}
  got  := p.GetAll()

  if !eqTaskSlice(want, got) {
    t.Errorf("\ngot : %+v\nwant: %+v", got, want)
  }
}

func TestSearchByCompleted(t *testing.T) {
  var t1 Task
  t1.New(false, "test task 1", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
  var t2 Task
  t2.New(false, "test task 2", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

  want1 := []Task{t1, t2}
  p   := initializeProject()
  got1 := p.SearchByCompleted(false)

  if !eqTaskSlice(got1, want1) {
    t.Errorf("\ngot : %+v\nwant: %+v", got1, want1)
  }

  var want2 []Task
  got2  := p.SearchByCompleted(true)

  if !eqTaskSlice(got2, want2) {
    t.Errorf("\ngot : %+v\nwant: %+v", got2, want2)
  }
}
