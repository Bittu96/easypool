# Easypool

Easiest way to create customized worker bot pools in golang

## Install
Fist, use go get to install the latest version of the library:
```
go get -u github.com/Bittu96/easypool
```

Next, include this package in your application:
```
import "github.com/Bittu96/easypool"
```

## Example
```
// set your task function in the pool
ep := easypool.New(mockTaskFunc)

// deploy number of concurrent bots you need
ep.Deploy(10)
```

Open to suggestions! 
Please feel free to contact me incase you need any help or help me improve this package :)
