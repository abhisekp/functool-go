package functool

type MapOptions struct {
    Concurrency uint8 /* MAX: 255 */
}

func (l *List) Map(fn Transformer, options ...*MapOptions) *List {
    var option *MapOptions
    for _, option = range options {
        break
    }
    if option != nil && option.Concurrency < 1 {
        option.Concurrency = 1
    }
    r := make(List, 0, len(*l))
    for _, item := range *l {
        r = append(r, fn(item))
    }
    return &r
}
