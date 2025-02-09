package main

import (
	"image/png"
	"log"
	"os"

	"go.linka.cloud/letteravatar"
)

var names = []string{
	"Alice Johnson",
	"Bob Smith",
	"Carol Davis",
	"Dave Brown",
	"Eve Wilson",
	"Frank Miller",
	"Gloria Taylor",
	"Henry Anderson",
	"Isabella Martinez",
	"James Thomas",
	"Жозефина Иванова",
	"Ярослав Петров",
}

func main() {
	for _, name := range names {
		img, err := letteravatar.DrawInitials(75, name, &letteravatar.Options{})
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create(name + ".png")
		if err != nil {
			log.Fatal(err)
		}

		err = png.Encode(file, img)
		if err != nil {
			log.Fatal(err)
		}
	}
}
