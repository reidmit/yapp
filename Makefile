build:
	go build ./cmd/yapp

example-hello-world:
	@ go run ./cmd/yapp run examples/hello-world

example-broken:
	@ go run ./cmd/yapp run examples/broken

example-static:
	@ go run ./cmd/yapp run examples/static

example-echo:
	@ go run ./cmd/yapp run examples/echo