package functool

func (l *List) Every(fn Predicate) (returnValue bool) {
    for _, item := range *l {
        if !fn(item) {
            return false
        }
    }
    return true
}
