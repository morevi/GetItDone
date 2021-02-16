package handler

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/morevi/GetItDone/pkg/tareas"
)

func Handler("/create"w http.ResponseWriter, r *http.Request) {
    var d tareas.Dashboard
    d.New(0)
    res, _ := json.Marshal(d)
    fmt.Fprintf(w, string(res))
}

