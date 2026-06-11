# irekqr

Minimalistic Telegram bot that generates QR codes from text or URLs.

Available in two implementations:

- **Go** — `main.go` (single binary, faster cold start)
- **Python** — `bot.py` (simpler to modify)

## Usage

1. Create a bot via [@BotFather](https://t.me/botfather) and get a token
2. Set the `BOT_TOKEN` environment variable

### Go

```bash
go build -o irekqr .
./irekqr
```

### Python

```bash
pip install -r requirements.txt
python bot.py
```

Send any text or URL to the bot — it replies with a QR code image.

## Deploy

Works anywhere: Railway, Render, Fly.io, PythonAnywhere, or a VPS.
