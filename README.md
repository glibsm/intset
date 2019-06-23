# package intset

Provides an easy way for human-readable int ranges. 

This was put together primarily to help manage the year range in the license
headers when a file is touched. It may be useful in other applications, but
keep in mind this code is not optimized.

## Install

```
go get github.com/glibsm/intset
```

## Example

```go
package main

import (
	"fmt"

	"github.com/glibsm/intset"
)

func main() {
	s := intset.Must(intset.Parse("2012-2014,2017,2018,2019"))
	s.Add(2020)
	fmt.Println(s) // 2012-2014,2017-2020
}
```
