// Generated by: main
// TypeWriter: slice
// Directive: +gen on MySliceType

package generate

// MySliceTypeSlice is a slice of type MySliceType. Use it where you would use []MySliceType.
type MySliceTypeSlice []MySliceType

// Where returns a new MySliceTypeSlice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv MySliceTypeSlice) Where(fn func(MySliceType) bool) (result MySliceTypeSlice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Count gives the number elements of MySliceTypeSlice that return true for the passed func. See: http://clipperhouse.github.io/gen/#Count
func (rcv MySliceTypeSlice) Count(fn func(MySliceType) bool) (result int) {
	for _, v := range rcv {
		if fn(v) {
			result++
		}
	}
	return
}

// GroupByString groups elements into a map keyed by string. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv MySliceTypeSlice) GroupByString(fn func(MySliceType) string) map[string]MySliceTypeSlice {
	result := make(map[string]MySliceTypeSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}