# kvhttpstore (key-value in-memory http storage)

you can store any data in this storage (e.g files)
service runs on :8080

## API

### GET /:key & POST /:key

#### POST: saves the raw body in a in-memory storage stored at key
#### GET: recieves the saved content at key in the in-memory storage

## run

```
$ go run cmd/server/main.go
```