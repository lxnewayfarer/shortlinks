It is a microservice that implements shortened links

## Features:

- Checking environment variables presence
- Graceful shutdown
- Health-check resource `/ping`
- Docker containerized

## Setup
- Fill `.env` file from `.env.example`
- Docker build and run:
    - `source .env`
    - `docker build --build-arg PORT=$PORT -t shortlinks .`
    - `docker run -p $PORT:$PORT shortlinks`
