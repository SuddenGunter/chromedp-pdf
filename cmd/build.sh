CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -o main .
docker build -t h2p:latest .
