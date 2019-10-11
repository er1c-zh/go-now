# go now

a tools set for golang

## log

一个简易好看的命令行日志

## assert

一个简单的测试用断言库

## go_boot

单体应用的脚手架，提供了如下功能：

### 优雅退出

```go
func main() {
    RegisterExitHandlers(func() {
    	// do some thing to exit gracefully
    })
    RegisterExitHandlers(func() {
    	// do other thing to exit gracefully
    })
    
    // do your work in another goroutine
    go func() {
    	// do your work
    	
    	// quit
    	Exit()
    }()
    
	WaitExit(0) // block until receive exit signal
	
	// or execute ExitHandlers with timeout
	// WaitExit(10 * time.Second)
}
```