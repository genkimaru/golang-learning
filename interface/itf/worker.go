package itf

// Worker define "sub class"
type Worker struct {
	Age      int
	Workshop string
}

// GetName implement the method
func (worker *Worker) GetName() string {
	return "YYY"
}

func FetchWorker() Person {
	return &Worker{12, "software"}
}
