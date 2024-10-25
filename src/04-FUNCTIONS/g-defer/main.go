package main

import (
	"log"
	"os"
)

/*
	-------------------------------------------------------------

	Defer is executed right before the exit

	In the example, it closes the file with "f.Cloes()" before

	Returning anything

	--------------------------------------------------------------
*/

func createTempFile() {
	f, err := os.Create("/tmp/foo.txt") // dosyayı oluştur
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close() // fonksiyondan çıkarken dosyayı kapat!
}

func main() {
	createTempFile()
}
