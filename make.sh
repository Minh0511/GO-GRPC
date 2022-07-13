protoc -I . -I ./googleapis \
    --go_out ./person \
    --go_opt paths=source_relative \
    --go-grpc_out ./person \
    --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./person \
    --grpc-gateway_opt paths=source_relative \
    personService.proto