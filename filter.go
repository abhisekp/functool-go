package functool

func (l *List) Filter(fn Predicate) *List {
    r := make(List, 0, len(*l))
    for _, item := range *l {
        if fn(item) {
            r = append(r, item)
        }
    }
    return &r
}
