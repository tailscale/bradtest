package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	env, _ := json.MarshalIndent(os.Environ(), "", "  ")
	fmt.Printf("Hello from Go. Env is: %s", env)
}
