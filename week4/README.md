# Week4 Create todo service by kratos

## start server
```shell
cd ./todo
go generate ./...
go build -o ./bin/ ./...
./bin/todo -conf ./configs
```

## how to test
After starting server, open terminal
`post`:
```shell
curl -X POST -H "Content-Type: application/json" -d '{"todo":"test"}' http://127.0.0.1:8000/todos
```

`get`:
```shell
url http://127.0.0.1:8000/todos
```

`delete`:
```shell
curl -X DELETE http://127.0.0.1:8000/todos/{id}
```

