GOARCH=wasm GOOS=js go build -o main.wasm main.go
#GOROOT=/usr/local/go
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .

