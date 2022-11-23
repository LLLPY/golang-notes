package main

import "github.com/kataras/iris/v12"

// import "github.com/kataras/iris/v12/mvc"

//Book example
type Book struct {
	Title string
}

func list(ctx iris.Context) {
	books := []Book{
		{"《Python封神之路》"},
		{"《Golang全集》"},
		{"《高性能的Go》"},
	}
	ctx.JSON(books)
}

func create(ctx iris.Context) {
	var b Book
	err := ctx.ReadJSON(&b)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("Book creation fialure").DetailErr(err))
		return
	}
	println("Received Book:", b.Title)
	ctx.StatusCode(iris.StatusCreated)
}

func main() {
	app := iris.New()

	booksAPI := app.Party("/books")
	booksAPI.Use(iris.Compression)
	booksAPI.Get("/", list)
	booksAPI.Post("/", create)

	app.Listen(":8080")
}
