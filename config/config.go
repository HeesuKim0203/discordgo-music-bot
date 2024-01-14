package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	discordToken  string
	youtubeToken  string
	queueSize     int
	musicDuration int
	botName       string
	searchListMax int
}

var c *Config

func init() {
	env, err := godotenv.Read("./.env")

	if err != nil {
		fmt.Println("Not found .env File!")
		fmt.Println("It fetches the specified environment variables, not from the .env file.")
		specifiedEnv()
	} else {
		envFile(env)
	}
}

func envFile(env map[string]string) {

	c = &Config{}

	discrodToken := env["DISCORD_PUBLIC_KEY"]

	if discrodToken != "" {
		c.discordToken = discrodToken
	} else {
		panic("Not found in .env file the 'DISCORD_PUBLIC_KEY'!")
	}

	youtubeToken := env["YOUTUBE_PUBLIC_KEY"]

	if youtubeToken != "" {
		c.youtubeToken = youtubeToken
	} else {
		panic("Not found in .env file the 'YOUTUBE_PUBLIC_KEY'!")
	}

	queueSize, err := strconv.Atoi(env["MUSIC_QUEUE_SIZE"])

	if err != nil {
		fmt.Println("Not correct format in .env file the 'MUSIC_QUEUE_SIZE'!")
		fmt.Println("Use the default for 'MUSIC_QUUE_SIZE'.")
		c.queueSize = 10
	} else {
		c.queueSize = queueSize
	}

	musicDuration, err := strconv.Atoi(env["MUSIC_DURATION"])

	if err != nil {
		fmt.Println("Not correct format in .env file the 'MUSIC_DURATION'!")
		fmt.Println("Use the default for 'MUSIC_DURATION'.")
		c.musicDuration = 480
	} else {
		c.musicDuration = musicDuration
	}

	botName := env["BOT_NAME"]

	if botName != "" {
		c.botName = botName
	} else {
		fmt.Println("Not found in .env file the 'BOT_NAME'!")
		fmt.Println("Use the default for 'BOT_NAME'.")
		c.botName = "!music"
	}

	searchListMax, err := strconv.Atoi(env["MAX_MUSIC_SEARCH_LIST"])

	if err != nil {
		fmt.Println("Not correct format in .env file the 'MAX_MUSIC_SEARCH_LIST'!")
		fmt.Println("Use the default for 'MAX_MUSIC_SEARCH_LIST'.")
		c.searchListMax = 5
	} else {
		c.searchListMax = searchListMax
	}

	fmt.Println("Config complete!")
}

func specifiedEnv() {

	c = &Config{}

	discrodToken := os.Getenv("DISCORD_PUBLIC_KEY")

	if discrodToken != "" {
		c.discordToken = discrodToken
	} else {
		panic("Not found the 'DISCORD_PUBLIC_KEY'!")
	}

	youtubeToken := os.Getenv("YOUTUBE_PUBLIC_KEY")

	if youtubeToken != "" {
		c.youtubeToken = youtubeToken
	} else {
		panic("Not found the 'YOUTUBE_PUBLIC_KEY'!")
	}

	queueSize, err := strconv.Atoi(os.Getenv("MUSIC_QUEUE_SIZE"))

	if err != nil {
		fmt.Println("Not correct format the 'MUSIC_QUEUE_SIZE'!")
		fmt.Println("Use the default for 'MUSIC_QUUE_SIZE'.")
		c.queueSize = 10
	} else {
		c.queueSize = queueSize
	}

	musicDuration, err := strconv.Atoi(os.Getenv("MUSIC_DURATION"))

	if err != nil {
		fmt.Println("Not correct format the 'MUSIC_DURATION'!")
		fmt.Println("Use the default for 'MUSIC_DURATION'.")
		c.musicDuration = 480
	} else {
		c.musicDuration = musicDuration
	}

	botName := os.Getenv("BOT_NAME")

	if botName != "" {
		c.botName = botName
	} else {
		fmt.Println("Not found the 'BOT_NAME'!")
		fmt.Println("Use the default for 'BOT_NAME'.")
		c.botName = "!music"
	}

	searchListMax, err := strconv.Atoi(os.Getenv("MAX_MUSIC_SEARCH_LIST"))

	if err != nil {
		fmt.Println("Not correct format the 'MAX_MUSIC_SEARCH_LIST'!")
		fmt.Println("Use the default for 'MAX_MUSIC_SEARCH_LIST'.")
		c.searchListMax = 5
	} else {
		c.searchListMax = searchListMax
	}

	fmt.Println("Config complete!")
}

func (c *Config) GetDiscordToken() string {
	return c.discordToken
}

func (c *Config) GetYoutubeToken() string {
	return c.youtubeToken
}

func (c *Config) GetQueueSize() int {
	return c.queueSize
}

func (c *Config) GetMusicDuration() int {
	return c.musicDuration
}

func (c *Config) GetBotName() string {
	return c.botName
}

func (c *Config) GetMusicSearchListMax() int {
	return c.searchListMax
}

func GetConfig() *Config {
	return c
}
