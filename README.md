# simple-go

# test

```
# 运行当前目录下所有 x_test.go
go test 
# 运行指定测试文件 （基本没法这么用，因为目标代码不会写在test.go中）
go test xxx_test.go -v 
# 运行指定目录下所有 x_test.go
go test ./command -v
# 运行当前项目下所有的 x_test.go 覆盖率
go test ./...  -cover -v 


# 运行所有的基准测试
go  test -benchmem  ./... -bench .
# 运行 ./command 目录下的所有基准测试
go  test -benchmem  ./command -bench . 
# 运行 simle-go/command 中的所有基准测试
go  test -benchmem  simple-go/command  command -bench . 
```
