# http-app

Simple HTTP server.

## Run

```bash
cd http/http-app

go run ./cmd/server
```

Test:

```bash
curl -i http://127.0.0.1:8080/
curl -i http://127.0.0.1:8080/healthz
```
