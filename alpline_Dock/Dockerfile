FROM golang:alpine
ENV GO111MODULE=on
ENV PORT=9000
WORKDIR /app/server

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o mongodb

#FROM scratch AS build
#COPY --from=build /go/bin/mongodb /go/bin/mongodb
CMD ["./mongodb"]