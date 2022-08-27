package main

import (
	"fmt"
	"log"
	"os"

	json "github.com/json-iterator/go"
	"github.com/silverswords/sand/server/web"
)

func main() {
	config := web.Config{
		Host: "localhost",
		Addr: "9090",
	}

	data, err := json.Marshal(config)
	if err != nil {
		fmt.Println("error:", err)
	}

	err = os.WriteFile("test", data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
