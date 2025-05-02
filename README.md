It is a microservice that implements shortened links

## Features:

- Checking environment variables presence
- Graceful shutdown
- Health-check resource `/ping`
- Docker containerized
- 2,18×10^14 variants of short link of 8 symbols like `fjdJ5sSf`
- Caching of shorten links
- Fully transactional

## Setup
- Fill `.env` file from `.env.example`
- Docker build and run (app only):
    - `source .env`
    - `docker build --build-arg PORT=$PORT -t shortlinks .`
    - `docker run -p $PORT:$PORT shortlinks`
- Docker comopse run (Redis and app):
    - `docker compose up`
- Tests with coverage: `go test -cover -v`

## Usage
Health-check

`GET /ping`

```
response:
200 OK
{
    "response": "pong"
}
```
___
Create short link

`POST /shorten`

```
body:
{
    "link": "https://example.com"
}
response:
200 OK
{
    "response": http://localhost:PORT/l/sweDwSLK
}
```
___
Use short link

`GET /l/{path}`

```
GET /l/{path}
response:
301 REDIRECT
```
