# gocopy-benchmark

比較 Go 在 Clone 一個 slice 時，使用 slice 與 make 的效能差異。

## 執行效能測試

1. 執行效能測試

    ```sh
    go test -bench=. -v
    ```

    > 參數 `-v` 可以顯示更加詳細的效能測試過程。

2. 執行效能測試（每項測試僅跑 3 秒）

    ```sh
    go test -bench=. -benchtime=3s
    ```

3. 取得 CPU, MEMORY, Block 剖析檔

    ```sh
    go test -v -bench=. -cpuprofile=cpu.out -memprofile=mem.out -blockprofile=block.out .
    ```

4. 分析 CPU, MEMORY, Block 剖析檔

    ```sh
    go tool pprof -text gocopy-benchmark.test.exe cpu.out
    go tool pprof -text gocopy-benchmark.test.exe mem.out
    go tool pprof -text gocopy-benchmark.test.exe block.out
    ```
