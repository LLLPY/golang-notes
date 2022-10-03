package main

type Fly interface {
	fly()
}

type Swim interface {
	swim()
}

type FlyFish interface {
	Fly
	Swim
}

type Fish struct {
}

func (f Fish) fly() {
	print("flying...\n")
}

func (f Fish) swim() {
	print("swimming...\n")
}

func main() {

	f := Fish{}
	f.fly()

	var ff FlyFish
	ff = Fish{}
	ff.fly()

}
