# Week4

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

