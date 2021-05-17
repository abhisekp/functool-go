# functool

## Install

```shell
$ go get github.com/abhisekp/functool-go
```

## Usage

```go
package main

import (
    "github.com/abhisekp/functool-go"
)

func main() {
    // Map //
    list := functool.List{1, 2, 3, 4}
    actual := list.Map(func(item interface{}) interface{} {
        num := uint8(item.(int))
        return num * num
    })
    //➔ [...]uint8{1, 4, 9, 16}
    
    
    // Filter //
    list := functool.List{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    actual := list.Filter(func(i interface{}) bool {
        num := i.(int)
        return num%2 == 0
    })
    //➔ [...]int{0, 2, 4, 6, 8, 10}

    
    // Every //
    list := functool.List{true, true, true}
    actual := list.Every(func (b interface{}) bool {
        return b.(bool)
    })
    //➔ true
    
    
    // Some //
    list := functool.List{false, false, true, false}
    actual := list.Some(func (b interface{}) bool {
        return b.(bool)
    })
    //➔ true

    
    // Over //
    actual := list.Over(max, min)
    //➔ [...]int{2, 1}
}
```