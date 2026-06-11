# irekqr

Minimalistic Telegram bot that generates QR codes from text or URLs.

## Usage

1. Create a bot via [@BotFather](https://t.me/botfather) and get a token
2. Set `BOT_TOKEN` environment variable
3. Build and run:

```bash
export BOT_TOKEN=your_token_here
go build -o irekqr .
./irekqr
```

Or with `.env` file (auto-loaded via shell):

```bash
set -a; source .env; set +a
go build -o irekqr .
./irekqr
```

Send any text or URL to the bot — it replies with a QR code image.
