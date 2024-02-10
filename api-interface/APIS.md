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
    "uid": "2d05c44d110a7f09fea37c86a25dc9f9"
  },
  "datetime": 1710011577018,
  "info": "OK",
  "request": "/api/queue [POST]",
  "status": 200
}
```

#### 2. /api/queue/:uid [GET]

This API gets information about the queued file

Response:

```json
{
  "data": {
    "filename": "filename.pdf",
    "queueCount": 0,
    "status": "processed",
    "uid": "2d05c44d110a7f09fea37c86a25dc9f9"
  },
  "datetime": 1710011841999,
  "info": "OK",
  "request": "/api/queue/2d05c44d110a7f09fea37c86a25dc9f9 [GET]",
  "status": 200
}
```

#### 3. /api/download/:uid [GET]

This API downloads ZIP-archive with JPEG files
