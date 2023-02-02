package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	env, _ := json.MarshalIndent(os.Environ(), "", "  ")
	envs := strings.ReplaceAll(string(env), "TOKEN", "NOT_A_THING")
	fmt.Printf("Hello from Go. Env is: %s\n", envs)
	fmt.Printf("Argv is %q\n", os.Args)

	fmt.Printf("length of the not thing: %v\n", len(os.Getenv("GITHUB_TOKEN")))
}
