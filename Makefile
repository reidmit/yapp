build:
	mkdir -p out
	go build -o out/yapp ./cli

example-hello-world:
	go run ./cli -f examples/hello-world/yapp.yml