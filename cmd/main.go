package main

func main() {
	// if err := repository.InitIndexMap("./data/"); err != nil {
	// 	panic(err)
	// }
	r := NewRouter()
	
	r.Run()
}