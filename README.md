highLevelStat
=============
[![Build Status](https://travis-ci.org/muratsplat/highLevelStat.svg)](https://travis-ci.org/MURATSPLAT/highLevelStat)


A basic Go Package that gets system the percent of cpu(s) usage on only Gnu/Linux Os.

I'm new in Go Language and this library is my first library in go. I want to improve my go skils by writing go code more and more.

The package is high level to access  the information of system status. 

Probably I will not support the package at long term support.

It likes this:

```go
package main

import (
	"fmt"
	"github.com/MURATSPLAT/highLevelStat"
	
)

func main() {

	go func() {

		for {

			var test highlevelstat.SystemStatus

			fmt.Printf("Cpu(s): %.f%% UsedMem: %.f%%\n", test.GetCpuUsage().CpuUsage, GetMemInfo().GetUsedMemForHuman())
		}

	}()

	var input string

	fmt.Scanln(&input)

}

```

output:
```sh
Cpu(s): 11% UsedMem: 32%
Cpu(s): 13% UsedMem: 32%
Cpu(s): 12% UsedMem: 32%
Cpu(s): 13% UsedMem: 32%
Cpu(s): 10% UsedMem: 32%
Cpu(s): 9% UsedMem: 32%
Cpu(s): 13% UsedMem: 32%
Cpu(s): 11% UsedMem: 32%
Cpu(s): 8% UsedMem: 32%
Cpu(s): 12% UsedMem: 32%
```
Installing
----------
 to download the package by using "get" parameter via go such as..
```sh
$ go get github.com/MURATSPLAT/highLevelStat
```
Include in your source code:

    import "github.com/MURATSPLAT/highLevelStat"


