
package tareas

//Objeto TaskList -------------------------
type TaskList struct {
  Name string
  Label []string
  Items []Task
}

//Obtener numero de tareas por realizar
func (l TaskList) nChecks() int{
  n := 0
  for i := 0; i < len(l.Items); i++ {
    if l.Items[i].Done == true {
      n++;
    }
  }
  return n
}

//Obtener solo las tareas no completadas
//TO DO
func (l TaskList) getUnchecked() []Task {
  t := []Task{}
  return t
}

//TRABAJAR CON LISTAS ------------------

//Buscar listas por etiqueta  
func searchByLabel(l []TaskList) []TaskList {
  newl := []TaskList{}
  return newl
}



