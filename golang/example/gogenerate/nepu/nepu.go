package nepu

// +gen slice:"Where,Count,GroupBy[string]"
type Nepu struct {
	ID        int64
	Name      string
	GroupName string
}

// +gen set
type Nowa struct{}

// +gen list
type Buran struct{}

// +gen ring
type Beru struct{}
