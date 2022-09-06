package main

import (
	"avocet/internal/app"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	// ====Route log to file====
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	var wg sync.WaitGroup
	//
	app.Run(&wg)

	wg.Wait()
	fmt.Println("Done")
}
