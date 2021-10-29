build:
	mkdir -p out
	go build -o out/yapp ./cmd/yapp

example-hello-world:
	go run ./cmd/yapp -f examples/hello-world/yapp.yml