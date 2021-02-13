
package tareas

import (
  "testing"
  "reflect"
)

func TestDashboardNew(t *testing.T) {
  t.Log("Test Dashboard.New()")

  var got Dashboard
  got.New(0)

  if got.ID != 0 {
    t.Errorf("Dashboard.ID\n want: 0\n got: %d", got.ID)
  }
  if len(got.Projects) != 0 {
    t.Errorf("len(Dashboard.Projects)\n want: 0\n got: %d", len(got.Projects))
  }
}

func TestDashboardGet(t *testing.T) {
  t.Log("Test Get All")
  var d Dashboard
  p1 := Project{[]string{"a", "b"}, "Description", []Task{}}
  p2 := Project{[]string{"c", "d"}, "Description 2", []Task{}}
  d.Add(p1)
  d.Add(p2)

  got := d.Get(0)
  if !reflect.DeepEqual(got, p1) {
    t.Errorf("Dashboard.Get\n want: %+v\n got: %+v", p1, got)
  }

  got = d.Get(1)
  if !reflect.DeepEqual(got, p2) {
    t.Errorf("Dashboard.Get\n want: %+v\n got: %+v", p2, got)
  }
}

func TestDashboardGetAll(t *testing.T) {
  t.Log("Test Get All")

  var d Dashboard
  d.New(0)

  var got []Project
  got = d.GetAll()

  if len(got) != 0 {
    t.Errorf("Dashboard.GetAll\n want: 0\n got: %d", len(got))
  }
}

func TestDashboardSearchByTags(t *testing.T) {
  t.Log("Test Seach project by tags")

  var want = Project{[]string{"a"}, "Desc1", []Task{}}
  var projects = []Project{ want, Project{[]string{"a","b","c"}, "Desc2", []Task{}} }
  var d = Dashboard{0, projects}
  var tags = []string{"a", "b"}

  got := d.SearchByTags(tags)
  if len(got) == 0 {
    t.Errorf("Dashboard.SearchByTags\n want: array size 1\n got: array size %+v", len(got))
  } else if !reflect.DeepEqual(got[0], want) {
    t.Errorf("Dashboard.SearchByTags\n want: %+v\n  got: %+v", want, got[0])
  }
}

func TestDashboardSearchByTasks(t *testing.T) {
  t.Log("Test Search by task description")

  var primera = Task{false, "blah", ""}
  var segunda = Task{false, "blah blah", ""}
  var pro2 = Project{[]string{"a","b","c"}, "Desc2", []Task{segunda}}
  var projects = []Project{
                            Project{[]string{"a"}, "Desc1", []Task{primera}},
                            pro2,
                          }
  var d = Dashboard{0, projects}

  got := d.SearchByTasks("blah blah")
  if len(got) == 0 {
    t.Errorf("len \n want: 1\n got: %d", len(got))
  } else if reflect.DeepEqual(got[0], pro2) {
    t.Errorf("\n want: {%+v}\n  got: %+v", pro2, got)
  }
}

func TestDashboardAdd(t *testing.T) {
  t.Log("Test Add project to dashboard")
  var pro = Project{[]string{"a","b","c"}, "Desc2", []Task{}}
  var want = []Project{pro}

  var d = Dashboard{ID:0}
  d.Add(pro)

  got := d.Projects

  if len(got) != 1 {
    t.Errorf("length\n want: 1\n  got: %+v", len(got))
  } else if !reflect.DeepEqual(got, want) {
    t.Errorf("project\n want: {%+v}\n  got: %+v", want, got)
  }
}


func TestDashboardRemove(t *testing.T) {
  t.Log("Test Remove project from dashboard")
  var pro = Project{[]string{"a","b","c"}, "Desc2", []Task{}}

  var d = Dashboard{ID:0}
  d.Add(pro)

  d.Remove(0)
  got := d.Projects

  if len(got) != 0 {
    t.Errorf("length\n want: 0\n  got: %+v", len(got))
  }
}

func TestDashboardUpdate(t *testing.T) {
  t.Log("Test Update project from dashboard")
  var pro = Project{[]string{"a","b","c"}, "Desc1", []Task{}}
  var pro_updated = Project{[]string{"a","b","c"}, "Desc2", []Task{}}

  var d = Dashboard{ID:0}
  d.Add(pro)

  d.Update(0, pro_updated)
  got := d.Projects

  if len(got) != 1 {
    t.Errorf("length\n want: 1\n  got: %+v", len(got))
  } else if !reflect.DeepEqual(got[0], pro_updated) {
    t.Errorf("update\n want: %+v\n  got: %+v", pro_updated, got[0])
  }
}

func TestFindStringInSlice(t *testing.T) {
  t.Log("Test String in slice of strings")
  items := []string{"a", "b", "c"}

  k, found := FindStringInSlice(items, "b")
  if !found {
    t.Errorf("FindStringInSlice found\n want: true\n got: %t", found)
  }
  if k != 1 {
    t.Errorf("FindStringInSlice pos\n want: 1\n got: %d", k)
  }
}
