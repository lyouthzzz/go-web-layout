```
protoc  --proto_path=.  --proto_path=/Users/y.liu/go/src/github.com/go-kratos/kratos/third_party --go_out=paths=source_relative:.  --go-grpc_out=paths=source_relative:.  --go-http_out=paths=source_relative:.  --go-errors_out=paths=source_relative:.  --openapiv2_out=.  --openapiv2_opt=logtostderr=true  --openapiv2_opt=json_names_for_fields=false  --openapiv2_opt=enums_as_ints=true  api/order/v1/order.proto
```

```
protoc  --proto_path=.  --proto_path=/Users/y.liu/go/src/stash.weimob.com/DataIR/library/go-xupiter  --proto_path=/Users/y.liu/go/src/stash.weimob.com/DataIR/library/go-xupiter/third_party  --go_out=paths=source_relative:.  --go-grpc_out=paths=source_relative:.  --go-http_out=paths=source_relative:.  api/order/v1/order.proto
```