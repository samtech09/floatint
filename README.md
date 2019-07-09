# floatint
Packer for small float values to int16 - golang

### Why packing float to int ?
Memory is precious for full in-memory databases like [VoltDB](http://community.voltdb.com/). Float is taking 8 bytes for the small values like 2.37. To optimize memory usage, such small float values can be packed into 2 byte int16, thus saving 6 byte per row.

### Limits
This package allow to pack values 
- up to 8.189 into int16
- up to 536870.875 into int32

### Installation
```
go get -u github.com/samtech09/floatint
```

### Usage

```
package main

import (
	"fmt"
	"github.com/samtech09/floatint"
)

func main() {
	floatvalue := 6.21
	fint16 := floatint.Pack(floatvalue)
	fmt.Printf("\n%f packed into %d\n", floatvalue, fint16)
	fmt.Printf("\t unpacked to %f", floatint.Unpack(fint16))

	floatvalue = 378283.537
	fint32 := floatint.Pack32(floatvalue)
	fmt.Printf("\n%f packed into %d\n", floatvalue, fint32)
	fmt.Printf("\t unpacked to %f\n", floatint.Unpack32(fint32))
}

```
