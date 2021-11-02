build:
	go build \
		-ldflags "-X main.appVersion=${YAPP_VERSION} -X main.appCommit=${YAPP_COMMIT}" \
		./cmd/yapp

example-hello-world:
	go run ./cmd/yapp examples/hello-world --debug

example-broken:
	go run ./cmd/yapp examples/broken --debug

example-static:
	go run ./cmd/yapp examples/static --debug

example-echo:
	go run ./cmd/yapp examples/echo --debug

example-fancy:
	go run ./cmd/yapp examples/fancy --debug

example-response-headers:
	go run ./cmd/yapp examples/response-headers --debug