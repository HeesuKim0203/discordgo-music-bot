package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/lets-go-bot/app"
)

func main() {

	discord := app.NewDiscord()

	err := discord.Open()
	if err != nil {
		fmt.Println("discord Open Error : ")
		panic(err)
	}

	registeredCommands := app.RegisterCommends(discord)

	log.Println(registeredCommands)

	defer app.UnRegister(discord, registeredCommands)
	defer discord.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
