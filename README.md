# Easypool

Workerpools made easy!

Easiest way to create customized worker bot pools in golang

## Install
Fist, use go get to install the latest version of the library:
```
go get -u github.com/Bittu96/easypool
```

Next, include this package in your application:
```go
import "github.com/Bittu96/easypool"
```

## Example
```Go
// set your task function in the pool
ep := easypool.New(mockTaskFunc)

// deploy number of concurrent bots you need
ep.Deploy(10)
```

Note: task function must be in the below format.

```go
func(interface{}) interface{}
```

 Make your function underlying like this function start using..!

 ```go
func(in interface{}) (out interface{}){
    out := yourFunction(in)
    return out
}
 ```

Open to suggestions! 
Please feel free to contact me incase you need any help or help me improve this package :)
