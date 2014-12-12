package generate

// +gen slice:"Where,Count,GroupBy[string]"
type MySliceType struct {
	ID        int64
	Name      string
	GroupName string
}

// +gen set
type MySetType struct{}

// +gen list
type MyListType struct{}

// +gen ring
type MyRingType struct{}
