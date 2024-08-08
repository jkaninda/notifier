# Notifier
## Build 
```sh
go build .
```

## Send email

```sh
go run . sendMail --message "Your message" --subject "Your subject"
```

```sh
notifier sendMail --message "Your message" --subject "Your subject"
```

```sh
go run . sendMessage
```

`.env`
```conf
MAIL_HOST=localhost
MAIL_PORT=1025
MAIL_USERNAME=jkaninda@mailhog.local
MAIL_PASSWORD=password
MAIL_TO=test@mailhog.local
MAIL_BODY="Lorem Ipsum is simply dummy text of the printing and typesetting industry."
MAIL_SUBJECT="Greetings"
MAIL_FROM=jkaninda@mailhog.local
## Telegram
TG_TOKEN=
TG_CHAT_ID=
TG_MESSAGE="Lorem Ipsum is simply dummy text of the printing and typesetting industry."
```