package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	//"app"
)

func main() {

	discord := NewDiscord()

	err = discord.Open()
	if err != nil {
		fmt.Println("discord Open Error : ")
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
}
