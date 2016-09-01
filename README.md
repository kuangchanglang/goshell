# goshell
run go as shell  

## Why 
Try some basic idea without creating a file, compiling, running.  
Just like what python does.  

## Installation  
### Requirements:  
1. golang 1.6 (if you are using golang 1.6-, ```go get github.com/chzyer/readline``` is required)
2. goimports(strongly recommend which help you import package you are using)

### Installation:  
go get github.com/kuangchanglang/goshell

## Basic Usage 
```shell
$ goshell 
>>> sum := 0   
>>> for i:=0; i<=100; i++{
>>> 	sum += i
>>> }
>>> import "fmt"
>>> fmt.Println(sum)
5050
>>> 
```

```shell
$ goshell
>>> func fi(n int) int{
>>> 	if n < 0 {
>>> 		return 0
>>> 	}
>>> 	if n <= 2{
>>> 		return 1
>>> 	}
>>> 	return fi(n-1) + fi(n-2)
>>> }
>>> fmt.Println(10)
10
>>> fmt.Println(3)
3
>>> for i:=0;i<10;i++{
>>> 	fmt.Println(fi(i))
>>> }
1
1
1
2
3
5
8
13
21
34
```

### hint
* type "quit" to exit  
* run ```stty erase ^h``` to enable backspace, run ```stty erase ^?``` to recover  

## Reference
[gosh](https://github.com/mkouhei/gosh)
