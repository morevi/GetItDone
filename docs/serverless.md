# Uso de sistemas serverless

## Conexión con Vercel, Github, y despliegue
Se ha enlazado la cuenta y el repositorio de Github con Vercel, de forma que Vercel comprueba si se han realizado _push_ en el repositorio, antes de hacer un despliegue.

Para lograr que se realicen los tests antes que el despliegue, se ha creado desde Vercel un webhook, y se ha añadido como _secreto_ al repositorio en Github. Luego, se ha creado un nuevo _workflow_: [deploy-vercel.yml](../.github/workflows/deploy-vercel.yml). Este _workflow_ espera a que acaben los tests ([dbt](../.github/workflows/dbt.yml)), y luego realiza una petición _POST_ con _curl_ a la URL de la API de Vercel, activando así el _build_ de Vercel.

El despliegue se puede comprobar [aquí](https://getitdone.vercel.app/).
En si se abre con el navegador llevará a una página directorio del proyecto. Pero podrán llevarse a cabo las peticiones _GET_ desde la propia URL.

## Funciones serverless
Se ha creado la función [api/project.go](api/project.go) que realiza las siguientes operaciones:

Acción        | Tipo de petición | Parámetros  | Respuesta | HU
------------- |:----------------:| ----------- | --------- | --- 
Listar todos los proyectos | GET | ninguno | Todos los proyectos     | [HU3](https://github.com/morevi/GetItDone/issues/3)
Buscar proyectos por tag   | GET | tags    | Proyectos con dicho tag | [HU3](https://github.com/morevi/GetItDone/issues/3)
Crear un nuevo proyecto sin _Tasks_   | POST | json del Proyecto   | Código _201_ | [HU2](https://github.com/morevi/GetItDone/issues/2)

Puede probarlos con:
```bash
curl -X GET https://getitdone.vercel.app/api/project
curl -X GET https://getitdone.vercel.app/api/project?tags=iv
curl -X POST https://getitdone.vercel.app/api/project -d '{"tags":["a","b"],"description":"desc","items":[]}'
```
