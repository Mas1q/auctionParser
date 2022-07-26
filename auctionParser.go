package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func print(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("START!")
	print(FetchRealmInfo("19555"))
}
