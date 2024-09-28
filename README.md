# Вебинар "Использование gRPC в Go"

## Dummy server

Сгенерируем код grpc-интерфейса сервера:

```bash
protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    internal/protos/v1/server_old/*.proto
```

Назначение флагов:
- `go_out`: 1. Генерация файлов Go 2. Указание output-директории (относительно `go_opt=paths`)
- `go_opt=paths`: определение куда именно в output-директории помещаются файлы: `source_relative` или `import`. Также можно использовать флаг `module=$PREFIX`

Пример:

```bash
protoc \
    --go_out=. \
    --go_opt=paths=import \
    --go-grpc_out=. \server_old
    --go-grpc_opt=module=webinar-service \
    internal/protos/v1/*.proto
```

См. документацию по флагам здесь: https://protobuf.dev/reference/go/go-generated/#invocation

## Best-practices

См. рекомендации здесь: https://protobuf.dev/programming-guides/dos-donts/

См. ветку `best-practices-demo` для примеров.

Protobuf не дает гарантий по стабильности алгоритма сериализации сообщений при изменении версии.

## Protobuf

### Encoding

Для разбора бинарного proto-сообщения можно использовать инструмент [protoscope](https://github.com/protocolbuffers/protoscope).

```bash
go install github.com/protocolbuffers/protoscope/cmd/protoscope...@latest
```
