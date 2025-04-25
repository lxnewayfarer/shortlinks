It is a microservice that implements shortened links

## Features:

- Checking environment variables presence
- Graceful shutdown
- Health-check resource `/ping`
- Docker containerized

## Setup
- Fill `.env` file from `.env.example`
- Docker build and run (app only):
    - `source .env`
    - `docker build --build-arg PORT=$PORT -t shortlinks .`
    - `docker run -p $PORT:$PORT shortlinks`
- Docker comopse run (Redis and app):
    - `docker compose up`

## Usage
Health-check

```
GET /ping
200 OK
{
    "response": "pong"
}
```
