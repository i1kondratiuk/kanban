FROM golang:1.14.4-alpine3.12

WORKDIR $GOPATH/src/github.com/kanban/

# Populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY config.yml /app/

# Build the Go app
RUN go build -o /app/kanban .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go build`
CMD /app/kanban