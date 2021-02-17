package handler

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/morevi/GetItDone/pkg/tareas"
)

func getFromDB() tareas.Dashboard {
  // El plan es que se cargue el dashboard de usuario desde una db,
  // pero por ahora no está implementado, asi que se generan datos
  // de ejemplo
  var d tareas.Dashboard
  d.New(91)

  var t tareas.Task
  t.New(true, "Hito 5", "2020-02-17")

  var p tareas.Project
  p.New([]string{"iv","serverless","despliegue"},
  "Hito 5 de la asignatura IV",
  []tareas.Task{t})
  d.Add(p)

  return d
}

func Handler(w http.ResponseWriter, r *http.Request) {
  d := getFromDB()

  switch r.Method {
  // Get every or some projects
  case "GET":
    tags, ok := r.URL.Query()["tags"]
    if !ok {
      res, _ := json.Marshal(d.GetAll())
      fmt.Fprintf(w, string(res))
    } else {
      res, _ := json.Marshal(d.SearchByTags(tags))
      fmt.Fprintf(w, string(res))
    }

  case "POST":
    body, err := ioutil.ReadAll(r.body)
    if err != nil {
	     fmt.Fprint(w, err.Error(), 500)
		    return
	   }

    var p tareas.Project
    err := json.Unmarshal(b, &p)
    if err != nil {
      d.Add(p)
      res, _ := json.Marshal("201 Created")
      fmt.Fprintf(w, string(res))
    } else {
      res, _ := json.Marshal("500 Something happened")
      fmt.Fprintf(w, string(res))
    }
  }
}

