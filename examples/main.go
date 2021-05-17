package main

import (
    "fmt"
    "functool"
    "time"
)

func main() {
    const (
        YYYY, MM, DD = "2006", "01", "02"
    )

    DoBLayout := struct {
        Half    string
        HalfRev string
        Full    string
        FullRev string
    }{
        Half:    fmt.Sprintf("%s/%s", DD, MM),
        HalfRev: fmt.Sprintf("%s/%s", MM, DD),
        Full:    fmt.Sprintf("%s/%s/%s", DD, MM, YYYY),
        FullRev: fmt.Sprintf("%s/%s/%s", YYYY, MM, DD),
    }

    layouts := functool.List{
        DoBLayout.Half,
        DoBLayout.HalfRev,
        DoBLayout.Full,
        DoBLayout.FullRev,
    }

    const invalidDateStr = "29/32/2001"
    const invalidDateStrLeap = "29/02/2001"
    const invalidDateStrHalf = "29/32"
    const validDateStrHalf = "29/02"
    const validDateStrFull = "28/02/2001"
    const validDateStrFullRev = "2001/02/28"

    times := functool.List{
        invalidDateStr,
        invalidDateStrLeap,
        invalidDateStrHalf,
        validDateStrHalf,
        validDateStrFull,
        validDateStrFullRev,
    }

    values := times.Map(func(_ts interface{}) interface{} {
        ts := _ts.(string)
        var r interface{}
        for _, _layout := range layouts {
            layout := _layout.(string)
            if t, err := time.Parse(layout, ts); err != nil {
                r = err
            } else {
                r = fmt.Sprintf("%s <===> %s", ts, t.Format(layout))
                return r
            }
        }
        return r
    })
    for _, value := range *values {
        switch value.(type) {
        case error:
            fmt.Println(value)
        default:
            fmt.Printf("Value: %s\n", value)
        }
    }
}

/** OUTPUT

parsing time "29/32/2001" as "2006/01/02": cannot parse "2/2001" as "2006"
parsing time "29/02/2001" as "2006/01/02": cannot parse "2/2001" as "2006"
parsing time "29/32" as "2006/01/02": cannot parse "2" as "2006"
Value: 29/02 <===> 29/02
Value: 28/02/2001 <===> 28/02/2001
Value: 2001/02/28 <===> 2001/02/28

**/
