## Build
FROM docker.io/golang:1.19 AS build

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

COPY *.go ./

RUN go build -o /tcp-test-linux-amd64

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /tcp-test-linux-amd64 /tcp-test-linux-amd64

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/tcp-test-linux-amd64"]
