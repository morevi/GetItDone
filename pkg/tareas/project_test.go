
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

func TestProjectNew(t *testing.T) {
  t.Log("Builds a project object by giving it data")
  want := initializeProject()

  var t1 Task
  t1.New(false, "test task 1", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
  var t2 Task
  t2.New(false, "test task 2", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

  var got Project
  got.New([]string{"proyecto", "iv", "test"}, "test description", []Task{t1, t2})

  if !reflect.DeepEqual(got,want) {
    t.Errorf("\ngot : %+v\nwant: %+v",got,want)
  }
}

func TestSetDescription(t *testing.T) {
  t.Log("Sets a new description for a project")
  want := "new project description"

  p := initializeProject()
  p.SetDescription(want)
  got := p.Description

  if got != want {
    t.Errorf("\ngot : %+v\nwant: %+v",got,want)
  }
}

func TetSetTags(t *testing.T) {
  t.Log("Sets a new group of tags for a project")
  p := initializeProject()

  p.SetTags("a b c")
  got := p.Tags
  want := []string{"a", "b", "c"}

  if !reflect.DeepEqual(got, want) {
    t.Errorf("\ngot : %+v\nwant: %+v",got,want)
  }
}

func TestAddTask(t *testing.T) {
  t.Log("Adds a new task to the project")
  p := initializeProject()

  var task Task
  task.New(false, "test", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
  want := append(p.Items, task)

  p.AddTask(task)
  got := p.Items

  if !reflect.DeepEqual(got, want) {
    t.Errorf("\ngot : %+v\nwant: %+v", got, want)
  }
}

func TestRemove(t *testing.T) {
  t.Log("Removes a task from a project")
  var task Task
  task.New(false, "test task 1", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

  want := []Task{task}

  p := initializeProject()
  p.Remove(1)
  got := p.Items

  if !reflect.DeepEqual(got, want) {
    t.Errorf("\ngot : %+v\nwant: %+v", got, want)
  }
}

func TestGet(t *testing.T) {
  t.Log("Get a task from a project by pos in array")
  var want Task
  want.New(false, "test task 1", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

  p := initializeProject()
  got := p.Get(0)

  if got != want {
    t.Errorf("\ngot : %+v\nwant: %+v", got, want)
  }
}

func TestGetAll(t *testing.T) {
  t.Log("Gets all tasks from the project")

  var t1 Task
  t1.New(false, "test task 1", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
  var t2 Task
  t2.New(false, "test task 2", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

  p := initializeProject()

  want := []Task{t1, t2}
  got  := p.GetAll()

  if !reflect.DeepEqual(want, got) {
    t.Errorf("\ngot : %+v\nwant: %+v", got, want)
  }
}

func TestSearchByCompleted(t *testing.T) {
  t.Log("Get the task whose property Done is as given")

  var t1 Task
  t1.New(false, "test task 1", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
  var t2 Task
  t2.New(false, "test task 2", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

  want1 := []Task{t1, t2}
  p   := initializeProject()
  got1 := p.SearchByCompleted(false)

  if !reflect.DeepEqual(got1, want1) {
    t.Errorf("\ngot : %+v\nwant: %+v", got1, want1)
  }

  var want2 []Task
  got2  := p.SearchByCompleted(true)

  if !reflect.DeepEqual(got2, want2) {
    t.Errorf("\ngot : %+v\nwant: %+v", got2, want2)
  }
}
