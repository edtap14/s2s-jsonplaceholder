# S2S JsonPlaceholder

## Descripción
Servidor http con mux

## Ejecutar

Para correr el proyecto en modo local
```sh
$ go run main.go
```
Para probar el servicio, enviar la siguiente petición

```sh
$ curl -v -X POST http://localhost:3000/ -H "Content-Type: application/json" -d '{"name": "Edgar"}'
```
