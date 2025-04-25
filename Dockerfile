FROM golang:alpine

ARG PORT=8080

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o shortlinks .

EXPOSE ${PORT}

CMD ["./shortlinks"]