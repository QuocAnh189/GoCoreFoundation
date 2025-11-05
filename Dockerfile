# Step 1: Modules caching
FROM golang:1.24.4 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.24.4 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN mkdir -p /bin && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd/server

ENTRYPOINT ["go", "run", "./cmd/server"]
