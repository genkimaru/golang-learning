package itf

// Person define the interface
type Person interface {
	GetName() string
}

// Student define "sub class"
type Student struct {
	Age   int
	Class int
}

// GetName implement the method
func (student Student) GetName() string {
	return "XXX"
}

func FetchStudent() Person {
	return Student{12, 12}
}
