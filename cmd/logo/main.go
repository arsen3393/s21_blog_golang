package main

import (
	"github.com/fogleman/gg"
	"log"
)

func main() {
	const W = 300
	const H = 300
	dc := gg.NewContext(W, H)

	dc.SetRGB(0.2, 0.6, 0.8)
	dc.Clear()

	dc.SetRGB(40, 40, 40)

	err := dc.LoadFontFace("cmd/logo/DejaVuSans-Bold.ttf", 48) // Путь к шрифту на вашей системе
	if err != nil {
		log.Fatal(err)
	}

	dc.DrawStringAnchored("Lashawnh", W/2, H/2, 0.5, 0.5)

	err = dc.SavePNG("./web/amazing_logo.png")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Логотип успешно сохранен в amazing_logo.png")
}
