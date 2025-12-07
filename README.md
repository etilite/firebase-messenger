# firebase-messenger

[![docker pulls](https://img.shields.io/docker/pulls/etilite/firebase-messenger)](https://hub.docker.com/r/etilite/firebase-messenger)
[![docker push](https://github.com/etilite/firebase-messenger/actions/workflows/docker.yml/badge.svg)](https://github.com/etilite/firebase-messenger/actions/workflows/docker.yml)
[![go build](https://github.com/etilite/firebase-messenger/actions/workflows/go.yml/badge.svg)](https://github.com/etilite/firebase-messenger/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/etilite/firebase-messenger/graph/badge.svg?token=PYVPKWSEP1)](https://codecov.io/gh/etilite/firebase-messenger)

`firebase-messenger` is a demo service to proxy sending firebase-admin multicast messages

## Features

## Usage

### Quick Start with Docker

```sh
docker run --rm -p 8080:8080 -e HTTP_HOST_PORT=:8080 etilite/firebase-messenger:latest
```

This will start the service and expose its API on port 8080.

### Build from source

```sh
git clone https://github.com/etilite/firebase-messenger.git
cd firebase-messenger
make run
```

This will build and run app at `http://localhost:8080`.

## Configuration

Set ENV variables

| ENV                              | Description                                | Mandatory |
|----------------------------------|--------------------------------------------|-----------|
| `HTTP_HOST_PORT`                 | host:port(default :3000)                   | No        |
| `WEBHOOK_AUTH_TOKEN`             | Bearer token api key for auth              | Yes       |
| `GOOGLE_APPLICATION_CREDENTIALS` | Path to JSON file Firebase service account | Yes       |

## API

### GET /health

Health check endpoint.

**Response:**

```json
{
  "status": "ok"
}
```

### POST /webhook/send

Send push-notifications.

**Headers:**

```
Authorization: Bearer <WEBHOOK_AUTH_TOKEN>
Content-Type: application/json
```

**Request Body:**

```json
{
  "tokens": [
    "device_token_1",
    "device_token_2"
  ],
  "notification": {
    "title": "title",
    "body": "Notification text"
  },
  "data": {
    "key": "value"
  }
}
```

**Response (200):**

```json
{
  "successCount": 2,
  "failureCount": 0,
  "responses": [
    {
      "success": true,
      "messageId": "..."
    },
    {
      "success": true,
      "messageId": "..."
    }
  ]
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

If you'd like to contribute to the project, please open an issue or submit a pull request on GitHub.
