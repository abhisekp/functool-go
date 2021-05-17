package functool_test

import (
    "github.com/abhisekp/functool-go"
    "math"
    "testing"
)

func max(l *functool.List) interface{} {
    a := (*l)[0]
    b := (*l)[1]
    return math.Max(float64(a.(int)), float64(b.(int)))
}

func min(l *functool.List) interface{} {
    a := (*l)[0]
    b := (*l)[1]
    return math.Min(float64(a.(int)), float64(b.(int)))
}

func TestOver(t *testing.T) {
    list := functool.List{1, 2}

    actual := list.Over(max, min)

    expected := [...]int{2, 1}

    if (
        expected[0] != int((*actual)[0].(float64))) ||
        (expected[1] != int((*actual)[1].(float64))) {
        t.Fatalf("Expected is %v. Got %v", expected, *actual)
    }
}
