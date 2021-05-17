package functool_test

import (
    "github.com/abhisekp/functool-go"
    "testing"
)

func TestFilterEven(t *testing.T) {
    list := functool.List{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    actual := list.Filter(func(i interface{}) bool {
        num := i.(int)
        return num%2 == 0
    })

    expected := [...]int{0, 2, 4, 6, 8, 10}

    for i, item := range *actual {
        num := item.(int)

        if expected[i] != num {
            t.Fatalf("expected even nums %v. Got %v", expected, *actual)
        }
    }

}
