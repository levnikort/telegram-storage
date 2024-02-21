# Telegram-storage

## Получение файлов

```bash
GET /:telegram_file_id
```

## Отправка файлов на сервер

```bash
POST /upload/:file_type
```

```ts
file_type: "photo" | "document"
```
photo - отправляет бинарный файл как фото в telegram
document - отправляет бинарный файл как документ в telegram

Надо указать заголовок **File-Name**

```bash
FIle-Name: good-photo.jpeg
```

### Переменные окружения

**TELEGRAM_BOT_TOKEN** - токен бота в telegram
**TELEGRAM_CHAT_ID** - идентификатор чата telegram в который будут приходить сообщения
**CACHE_EXPIRATION_DATE** - время хранения кэша в минутах
**CACHE_ELEMENTS** - лимит элементов в кэше
**HTTP_SERVER_PORT** - порт http сервера