package functool

func (l *List) Over(fns ...Transformer) *List {
    r := make(List, 0, len(fns))
    for _, fn := range fns {
        r = append(r, fn(l))
    }
    return &r
}
