# Вебинар "Использование gRPC в Go"

# Dummy server

Сгенерируем код grpc-интерфейса сервера:

```bash
protoc \
    --go_out=. \
    --go_opt=paths=import \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    internal/protos/*.proto
```

Назначение флагов:
- `go_out`: 1. Генерация файлов Go 2. Указание output-директории (относительно `go_opt=paths`)
- `go_opt=paths`: определение куда именно в output-директории помещаются файлы: `source_relative` или `import`

## Protobuf

### Encoding

Для разбора бинарного proto-сообщения можно использовать инструмент [protoscope](https://github.com/protocolbuffers/protoscope).

```bash
go install github.com/protocolbuffers/protoscope/cmd/protoscope...@latest
```