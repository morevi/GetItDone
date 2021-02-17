package handler

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/morevi/GetItDone/pkg/tareas"
)

func getFromDB() tareas.Dashboard {
  // El plan es que se cargue el dashboard desde una db,
  // por ahora no está implementado, asi que se generan datos
  // de ejemplo
  var d tareas.Dashboard
  d.New(91)

  var t tareas.Task
  t.New(true, "Hito 5", "2020-02-17")

  var p tareas.Projecto
  p.New([]string{"iv","serverless","despliegue"},
  "Hito 5 de la asignatura IV",
  []tareas.Task{t})
  t.Add(p)

  return d
}

func Handler(w http.ResponseWriter, r *http.Request) {
  d := getFromDB()

  switch r.Method{
  // Get every or some projects
  case "GET":
    tags, ok := r.URL.Query()["tags"]
    if !ok || len(tags[0]) < 1{
      res, _ := json.Marshal(d.GetAll())
      fmt.Fprintf(w, string(res))
    } else {
      key := keys[0]
      log.Println("Url Param 'key' is: " + string(key))
      res, _ := json.Marshal(d.SearchByTags(tags))
      fmt.Fprintf(w, string(res))
    }
  }
}

