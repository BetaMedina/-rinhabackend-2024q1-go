FROM golang:1.22-bullseye AS build-stage

WORKDIR /app

COPY go.* ./
COPY .env ./
RUN go mod download

COPY . ./

RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-buildvcs=false go build -v -o server cmd/main.go 

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /rinha
COPY .env ./
COPY --from=build-stage /app/server ./server

EXPOSE 8000

CMD ["/rinha/server"]