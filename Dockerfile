# Build the go application into a binary
FROM golang:alpine as builder
WORKDIR /app
ADD . ./
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -a -installsuffix cgo -o bin/discordgo-music-bot .

FROM alpine:3.16
ENV DISCORD_PUBLIC_KEY=""
ENV YOUTUBE_PUBLIC_KEY=""
ENV MUSIC_QUEUE_SIZE=""
ENV MUSIC_DURATION=""
ENV BOT_NAME=""
ENV MAX_MUSIC_SEARCH_LIST=""
ENV APP_HOME=/app
WORKDIR ${APP_HOME}
RUN apk --update add --no-cache ca-certificates ffmpeg opus
COPY --from=builder /app/bin/discordgo-music-bot ./bin/discordgo-music-bot
RUN yt-dlp --version
ENTRYPOINT ["/app/bin/discordgo-music-bot"]