package main

//set analog
func main() {
	seen := make(map[string]struct{})
	//...
	if _, ok := seen["w"]; !ok {
		seen["w"] = struct{}{}
	}
}
