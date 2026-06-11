import io
import logging
import os

import qrcode
from telegram import Update
from telegram.ext import Application, CommandHandler, MessageHandler, filters

logging.basicConfig(
    format="%(asctime)s - %(name)s - %(levelname)s - %(message)s",
    level=logging.INFO,
)
logger = logging.getLogger(__name__)


async def start(update: Update, _context):
    await update.message.reply_text(
        "Send me any text or URL and I'll generate a QR code for it."
    )


async def generate_qr(update: Update, _context):
    text = update.message.text.strip()
    if not text:
        await update.message.reply_text("Text can't be empty.")
        return

    if len(text) > 4000:
        await update.message.reply_text("Text is too long (max 4000 characters).")
        return

    try:
        img = qrcode.make(text)
        buf = io.BytesIO()
        img.save(buf, format="PNG")
        buf.seek(0)
        await update.message.reply_photo(photo=buf)
    except Exception as e:
        logger.exception("Failed to generate QR code")
        await update.message.reply_text("Something went wrong. Try again.")


def main():
    token = os.getenv("BOT_TOKEN")
    if not token:
        raise RuntimeError("BOT_TOKEN environment variable is not set")

    app = Application.builder().token(token).build()
    app.add_handler(CommandHandler("start", start))
    app.add_handler(MessageHandler(filters.TEXT & ~filters.COMMAND, generate_qr))

    app.run_polling(allowed_updates=Update.ALL_TYPES)


if __name__ == "__main__":
    main()
