# Reborn IRC #glazov bot
A Telegram version of the old known bot.

## Usage
Create *.env* file:
```ini
TELEGRAM_TOKEN=telegram_bot_token
TELEGRAM_CHATID=telegram_channel_id
DEBUG=false
```

Create *docker-compose.yaml* file:
```yml
version: "3"

services:
    app:
        image: salemgolem/glazov-bot:latest
        restart: always
        environment:
            TELEGRAM_TOKEN: ${TELEGRAM_TOKEN}
            TELEGRAM_CHATID: ${TELEGRAM_CHATID}
            DEBUG: ${DEBUG}
```

Run
```bash
docker compose up -d
```

## Supported commands
`/nick Nickname` - setup your old nickname from IRC
