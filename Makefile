build:
	go build ./cmd/yapp

example-hello-world:
	@ go run ./cmd/yapp run examples/hello-world/yapp.yml