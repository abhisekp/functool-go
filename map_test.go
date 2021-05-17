package functool_test

import (
    "github.com/abhisekp/functool-go"
    "testing"
)

func TestMapSquared(t *testing.T) {
    list := functool.List{1, 2, 3, 4}
    actual := list.Map(func(item interface{}) interface{} {
        num := uint8(item.(int))
        return num * num
    })
    expected := [...]uint8{1, 4, 9, 16}

    for i, squareNum := range *actual {
        if expected[i] != squareNum {
            t.Fatalf(
                "Expected map of %v should be %v but got %v",
                list,
                expected,
                actual,
            )
        }
    }

}

func TestMapNewStruct(t *testing.T) {
    type Person struct {
        Name string
        Age  uint8
    }
    list := functool.List{
        Person{
            Name: "Jack",
            Age:  15,
        },
        Person{
            Name: "Tommy",
            Age:  30,
        },
        Person{
            Name: "David",
            Age:  25,
        },
    }
    actual := list.Map(func(item interface{}) interface{} {
        return item.(Person).Age
    })
    expected := [...]uint8{15, 30, 25}

    for i, age := range *actual {
        if expected[i] != age {
            t.Fatalf(
                "Expected map of %v should be %v but got %v",
                list,
                expected,
                actual,
            )
        }
    }

}
