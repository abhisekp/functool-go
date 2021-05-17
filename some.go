package functool

func (l *List) Some(fn Predicate) (returnValue bool) {
    for _, item := range *l {
        if fn(item) {
            return true
        }
    }
    return false
}
