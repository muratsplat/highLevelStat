highLevelStat
=============
[![Build Status](https://travis-ci.org/MURATSPLAT/highLevelStat.svg)](https://travis-ci.org/MURATSPLAT/highLevelStat)


A basic library that gets system the percent of cpu(s) usage 

I'm new in Go Language and this library is my first library in go. I want to improve my go skils by writing go code more and more.

The package is high level to access  the information of system status.

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

			fmt.Printf("Cpu(s): %.f%%\n", test.GetCpuUsage().CpuUsage)

		}

	}()

	var input string

	fmt.Scanln(&input)

}

```

output:
```sh
Cpu(s): 0%
Cpu(s): 7%
Cpu(s): 0%
Cpu(s): 3%
Cpu(s): 0%
Cpu(s): 0%
Cpu(s): 3%
Cpu(s): 3%
Cpu(s): 0%
```
Installing
----------
Run

    go get github.com/MURATSPLAT/highLevelStat

Include in your source code:

    import "github.com/MURATSPLAT/highLevelStat"


