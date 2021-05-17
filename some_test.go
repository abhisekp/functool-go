package functool_test

import (
    "github.com/abhisekp/functool-go"
    "testing"
)

func TestSomeTruthy(t *testing.T) {
    list := functool.List{
        false,
        false,
        true,
        false,
    }

    actual := list.Some(func (b interface{}) bool {
        return b.(bool)
    })

    if !actual {
        t.Fatalf("Should result True for mixed boolean")
    }
}


func TestSomeFalsy(t *testing.T) {
    list := functool.List{
        false,
        false,
        false,
    }

    actual := list.Some(func (b interface{}) bool {
        return b.(bool)
    })

    if actual {
        t.Fatalf("Should result False for all false")
    }
}
