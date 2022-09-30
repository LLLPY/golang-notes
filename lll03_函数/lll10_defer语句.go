package main

func main() {
	print("start\n")
	defer print("step 1\n")
	defer print("step 2\n")
	defer print("step 3\n")
	print("end\n")
}
