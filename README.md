# Music bot

Discord Bot Using Go

## ENV
The [godotenv](godotenv) library is in use.  

### .env
|Name|Description|type|Default|
|:---|:---|:---:|:---:|
|DISCORD_PUBLIC_KEY|Please look at the [discord bot guide](https://discord.com/developers/docs/getting-started#step-1-creating-an-app) and create a token and put it in.|string|X|
|YOUTUBE_PUBLIC_KEY|Please create the key of [YouTube Data API v3](https://developers.google.com/youtube/v3/docs?hl=en) and put it in.|string|X|
|MUSIC_QUEUE_SIZE|It refers to the maximum number of music you can have as a playlist.|int|"10"|
|MUSIC_DURATION|It means the maximum music run time.|int|"480"|
|BOT_NAME|Indicates the name of the bot and the first command.(Case insensitive)|string|"!music"|
|MAX_MUSIC_SEARCH_LIST|Indicates the number of result values in the "sarch" command.|int|"5"|

## Command

You must use it with the commands set in the ```BOT_NAME``` mentioned above.

```
!music [Command] [Text or No Input]
```

#### Example

```
!music search lil nas x
```

|Command|Description|
|:---|:---|
|search|Searches for a song. You must enter the song title as text.|
|add|Adds a song to the playlist. You need to enter the song's id. The id of the song can be found in the search results.|
|delete|Deletes a song from the playlist. You must enter the id of the song you want to delete.|
|view|Displays the current playlist and the songs playing in the streaming playlist.|
|play|Streams the current playlist. The streaming playlist cannot be changed.|
|exit|Deletes all songs in the currently streaming playlist.|
|skip|Skips one song.|
|help|Provides a manual for the commands.|

## Docker Hub

The image file is provided on [Docker hub](https://hub.docker.com/repository/docker/heesu998/discordgo-music-bot/general)!  
After pulling it, you can use it by entering the commands in order.

#### Pull docker image
```
docker pull heesu998/discordgo-music-bot
```

#### Running a container
```
sudo docker run -e DISCORD_PUBLIC_KEY=[your_public_key] -e YOUTUBE_PUBLIC_KEY=[your_public_key] --name [container name] heesu998/discordgo-music-bot
```
