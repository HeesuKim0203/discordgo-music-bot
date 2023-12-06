package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscordSession(t *testing.T) {
	assert := assert.New(t)
	discord := NewDiscord()

	err := discord.Open()
	assert.NoError(err)

	discord.Close()
}
