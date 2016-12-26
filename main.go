package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/heroku/ir_bot/github"
	"github.com/joho/godotenv"
	"github.com/yelinaung/go-haikunator"
)

func rand_name() string {
	time := time.Now().UTC().UnixNano()
	haikunator := haikunator.New(time)
	rand.Seed(time)
	name := haikunator.Haikunate()
	number := strconv.Itoa(rand.Intn(8999) + 1000)
	name += "-"
	name += number
	return (name)

}
func main() {
	// grab environment variables from .env if it exists but throw no error
	// if it does not exist.
	godotenv.Load()

	newCommand := flag.NewFlagSet("new", flag.ExitOnError)
	nameFlag := newCommand.String("name", rand_name(), "A name for the incident")
	descFlag := newCommand.String("desc", "", "A short description for the incident")

	if len(os.Args) == 1 {
		fmt.Println("Shit ain't right. Use the new command")
		return
	}

	switch os.Args[1] {
	case "new":
		newCommand.Parse(os.Args[2:])
	default:
		fmt.Println("bitches be trippin")
	}

	if newCommand.Parsed() {
		repo := github.New(*nameFlag, *descFlag)
		fmt.Println("Created new repo:", repo)
	}
}
