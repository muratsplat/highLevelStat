highLevelStat
=============
[![Build Status](https://travis-ci.org/muratsplat/highLevelStat.svg)](https://travis-ci.org/muratsplat/highLevelStat)


A basic Go Package that gets system the percent of cpu(s) usage on only Gnu/Linux Os.

I'm new in Go Language and this library is my first library in go. I want to improve my go skils by writing go code more and more.

The package is high level to access  the information of system status. 

Probably I will not support the package at long term support.

It likes this:

```go
package main

import (
	"fmt"
	stat "github.com/muratsplat/highLevelStat"
)

func main() {

	var duration  int64 = 50 // millisecond

	// setting time range value in the package..
	// This value only effects the range time duration of cpu status samples..	
	stat.SetTimeOfRangeForCpuStat(duration)

	go func() {

		for {
			var test stat.SystemStatus
			// getting information the structer of memory

			var memInfo *stat.MemInfo = stat.GetMemInfo()

			fmt.Printf("Cpu(s): %.f%% UsedMem: %.f%%\n", test.GetCpuUsage().CpuUsage, memInfo.GetUsedMemForHuman())
		}
	}()

	// To block main method by working fmt.Scanln() method

	var in string
	fmt.Scanln(&in)

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
$ go get github.com/muratsplat/highLevelStat
```
Include in your source code:

    import "github.com/muratsplat/highLevelStat"


