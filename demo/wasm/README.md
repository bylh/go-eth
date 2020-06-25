# wasm start [官网链接](https://github.com/golang/go/wiki/WebAssembly#getting-started)
 * 构建主文件
 * 进行编译：GOOS=js GOARCH=wasm go build -o main.wasm
 * 复制环境文件：cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
 * 新建html文件引用相关文件 go get -u github.com/shurcooL/goexec
 * 启动一个服务， 可以使用goexec命令行执行启动（未安装需要先安装：go get -u github.com/shurcooL/goexec）
 * goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(``)))'
