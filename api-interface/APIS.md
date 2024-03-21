### Available APIs

These APIs are available via HTTP

#### 1. /api/queue [POST]

This API uploads the file and queues it for processing

Body type: `form-data`

Body content:

```json
{
  "pdf": "<selected PDF file>"
}
```

Response:

```json
{
  "data": {
    "fileName": "filename.pdf",
    "queueCount": 1,
    "status": "queued",
    "uid": "<uid>"
  },
  "datetime": 1710011577018,
  "info": "OK",
  "request": "/api/queue [POST]",
  "status": 200
}
```

#### 2. /api/queue [GET]

This API loads paginated queue data

Example request:

```text
/api/queue?limit=10&page=5
```

Both `limit` and `page` query parameters are optional

Minimal `limit` is 1, maximum `limit` is 100 

Response:

```json
{
  "datetime": 1710011577018,
  "info": "OK",
  "request": "/api/queue/<uid> [DELETE]",
  "status": 200
}
```

#### 3. /api/queue/:uid [DELETE]

This API deletes queue entry by UID

Response:

```json
{
  "datetime": 1710011577018,
  "info": "OK",
  "request": "/api/queue/<uid> [DELETE]",
  "status": 200
}
```

#### 4. /api/queue/:uid [GET]

This API gets information about the queued file

Response:

```json
{
  "data": {
    "filename": "filename.pdf",
    "queueCount": 0,
    "status": "processed",
    "uid": "<uid>"
  },
  "datetime": 1710011841999,
  "info": "OK",
  "request": "/api/queue/<uid> [GET]",
  "status": 200
}
```

#### 5. /api/download/:uid [GET]

This API downloads ZIP-archive with JPEG files, it returns a filestream with data
