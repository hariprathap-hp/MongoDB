FROM golang:1.16.4

ENV PORT=9000
WORKDIR /app/server

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o server
CMD ["./server"]