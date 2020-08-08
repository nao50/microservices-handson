# microservices-handson

Hands-on for microservices written in go and angular

# Run server

```sh
$ go run backend/*.go
```

# API usage

```sh
$ # create todo
$ curl -X POST -H "Content-Type: application/json" -d '{"id": "hoge", "name": "write backend server", "uid": "nao", "done": false}' localhost:5050/todos
$
$ # get todo
$ curl localhost:5050/todos/hoge
$
$ # list todos
$ curl localhost:5050/todos
$
$ # update todo
$ curl -X PUT -H "Content-Type: application/json" -d '{"id": "hoge", "name": "write frontend server", "uid": "nao", "done": false}' localhost:5050/todos/gaga
$
$ # delete todo
$ curl -X DELETE localhost:5050/todos/hoge
```
