package main

import "fmt"

// Server doc
type Server int

// Add doc
func (s *Server) add(pattern string, obj interface{}, methods ...string) []int {

	return nil
}

// Add doc
func (s *Server) add2(pattern string, methods ...string) []int {

	return nil
}

// Arr doc
type Arr []int

// Sort big -> small
func (arr *Arr) Sort() {
	v := *arr
	var len = len(v)
	for i := 0; i < len; i++ {
		for j := 0; j < i; j++ {
			if v[i] > v[j] {
				var temp int
				temp = v[i]
				v[i] = v[j]
				v[j] = temp
			}
		}
	}

}

func main() {
	var s Server = 1
	s.add("", "", "")
	s.add2("")
	var arr Arr = []int{2, 5, 3, 10, 7}
	arr.Sort()
	fmt.Println(arr)
}
