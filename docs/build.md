# Как сбилдить

Все команды выполняем из корня проекта

```bash
go build -o bin/test1 -ldflags "-s -w" \
./entrypoint/main.go
```