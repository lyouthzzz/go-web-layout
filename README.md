# go-web-layout

### generate grpc proto
```
// step one
cd api
// step two
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    v1/user.proto
// step three
protoc-go-inject-tag -input=v1/user.pb.go
```

### todo

- grpc server
- api code design
- migration
- middleware
- log
- opentracing


### 参考

- https://github.com/go-kratos/kratos-layout
- https://github.com/bxcodec/go-clean-arch
- https://github.com/eminetto/clean-architecture-go