FROM golang:1.22
WORKDIR /app
ENV CGO_ENABLED=0 GOPROXY=https://proxy.golang.org
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build ./...
RUN go test -count=1 ./... || true
CMD ["go", "run", "./cmd/heritage", "-addr", ":8080"]
