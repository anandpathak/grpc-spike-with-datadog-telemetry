
```
 protoc -I .  --proto_path=demogrpc/   --go_out=./ demogrpc/internal/proto-files/services/ping-service.proto
protoc -I .  --proto_path=demogrpc/internal/proto-files/domain/ping.proto   --go_out=./ demogrpc/internal/proto-files/services/ping-service.proto
```