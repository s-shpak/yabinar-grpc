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

## V2

Зарефакторим dummy-сервис. Новые proto-определения находятся в `internal/protos/v2/dummy`.

Обратите внимание на изменения вызова `protoc`.

Теперь определение сервиса `dummy` состоит из нескольких proto-файлов, передать определения в Postman как раньше не получится.

## gRPC reflection

См. `internal/api/api.go` для примера подключения серверной рефлексии.

Postman позволяет получить информацию по gRPC-эндпоинтам при помощи рефлекии. Но сейчас этот функционал сломан. Что же произошло ?

Используетс `grpcurl` для "отладки":

```bash
# получим информацию по всем сервисам
grpcurl -plaintext localhost:8081 list

# посмотрим на интересующий нас сервис
grpcurl -plaintext localhost:8081 describe practicum.go.grpc_webinar.v2.dummy.Dummy
```

В чем дело ?

Здесь несколько устаревшее, но полезное обсуждение: https://github.com/fullstorydev/grpcurl/issues/22

Исправим проблему и убедимся, что теперь Postman может использовать рефлексию.

Также для дебаггинга могут пригодиться переменные окружения `GRPC_GO_LOG_SEVERITY_LEVEL` и `GRPC_GO_LOG_VERBOSITY_LEVEL`:

```bash
GRPC_GO_LOG_SEVERITY_LEVEL=info GRPC_GO_LOG_VERBOSITY_LEVEL=2 ./cmd/server/server
```

## Api best practices

См. рекомендации здесь: https://protobuf.dev/programming-guides/api/




## Protobuf

### Encoding

Для разбора бинарного proto-сообщения можно использовать инструмент [protoscope](https://github.com/protocolbuffers/protoscope).

```bash
go install github.com/protocolbuffers/protoscope/cmd/protoscope...@latest
```
