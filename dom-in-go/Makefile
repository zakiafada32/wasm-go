.PHONY: build serve

build: wasm_exec.js
	GOOS=js GOARCH=wasm go build -o factorial.wasm factorial.go

serve:
	go run ../web-server/server.go
