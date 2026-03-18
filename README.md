# standalone-phasher

A small web server that computes **perceptual hashes (phash)** for images using [goimagehash](https://github.com/corona10/goimagehash). Accepts one image per request and returns the phash as JSON.

## Run locally

```bash
go run .
```

Server listens on `http://localhost:8080` by default. Set `PORT` to use another port.

## Run with Docker

```bash
docker build -t phasher .
docker run -p 8080:8080 phasher
```

## API

### POST `/phash`

Upload a single image as multipart form data.

**Request:** `multipart/form-data` with one file under the field name `image`.

**Response (200):**

```json
{
  "phash": "p:0a1b2c3d4e5f6789",
  "hash": 123456789012345678
}
```

**Error (4xx/5xx):** `{"error": "message"}`

### curl example

Send an image file the same way the UI does:

```bash
curl -F "image=@/path/to/photo.jpg" http://localhost:8080/phash
```

Example with a local file:

```bash
curl -F "image=@./image.png" http://localhost:8080/phash
```

## UI

Open `http://localhost:8080` in a browser. Use the form to pick an image and click **Compute phash**. The request is the same as the curl example (multipart form, field name `image`); the response is shown on the page.
