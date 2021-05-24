```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/user/v1/user.proto 
```

```
protoc-go-inject-tag -input=api/user/v1/user.pb.go
```