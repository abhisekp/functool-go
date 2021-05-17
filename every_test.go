package functool_test

import (
    "github.com/abhisekp/functool-go"
    "testing"
)

func TestEveryMixed(t *testing.T) {
    list := functool.List{
        false,
        false,
        true,
    }

    actual := list.Every(func (b interface{}) bool {
        return b.(bool)
    })

    if actual {
        t.Fatalf("Should result False for mixed boolean")
    }
}

func TestEveryAllTruthy(t *testing.T) {
    list := functool.List{
        true,
        true,
        true,
    }
    actual := list.Every(func (b interface{}) bool {
        return b.(bool)
    })

    if !actual {
        t.Fatalf("Should result True for all Truthy")
    }
}
