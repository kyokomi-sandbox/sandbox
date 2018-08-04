package extpoints

type MyPlugin interface {
	Before(name string)
	Run()
	After()
	String() string
}
