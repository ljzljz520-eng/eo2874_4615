# eo2874_4615

Go project for module `traininganalysis`.

## Standard commands

```bash
go build ./...
go test -count=1 ./...
```

## Run

```bash
go run ./cmd/coach
```

## Docker validation

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh my-go-task linux/arm64
./build_benzhi_docker.sh my-go-task linux/amd64
docker run -it my-go-task:latest
```
