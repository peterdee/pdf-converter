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
    "queuePosition": 1,
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
  "data": {
    "items": [
      {
        "createdAt": 1709641275824,
        "downloadedAt": 0,
        "id": "65e70e3b43c77e4c31a0aa7f",
        "originalFileName": "resume.pdf",
        "status": "processed",
        "totalDownloads": 0,
        "uid": "6aeff35982fbe2752312bb67284b5a34",
        "updatedAt": 1709641321114
      },
      {
        "createdAt": 1709641285065,
        "downloadedAt": 0,
        "id": "65e70e4543c77e4c31a0aa80",
        "originalFileName": "resume.pdf",
        "status": "processed",
        "totalDownloads": 0,
        "uid": "1deaa3907d8e361440e062c631b8e36d",
        "updatedAt": 1709641381214
      },
      {
        "createdAt": 1709641329083,
        "downloadedAt": 0,
        "id": "65e70e7143c77e4c31a0aa81",
        "originalFileName": "resume.pdf",
        "status": "processed",
        "totalDownloads": 0,
        "uid": "38fe7cbef957b249a0c9f36e49b70f2b",
        "updatedAt": 1709641441089
      },
    ],
    "limit": 10,
    "page": 1,
    "totalItems": 3,
    "totalPages": 1
  },
  "datetime": 1710011577018,
  "info": "OK",
  "request": "/api/queue?limit=10&page=1 [GET]",
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
			"createdAt": 1710011841999,
			"downloadedAt": 1710011841999,
			"id": "65e70e3b43c77e4c31a0aa7f",
			"originalFileName": "filename.pdf",
			"queuePosition": 1,
			"status": "queued",
			"totalDownloads": 0,
			"uid": "<uid>",
			"updatedAt": 1710011841999,
  },
  "datetime": 1710011841999,
  "info": "OK",
  "request": "/api/queue/<uid> [GET]",
  "status": 200
}
```

#### 5. /api/download/:uid [GET]

This API downloads ZIP-archive with JPEG files, it returns a filestream with data
