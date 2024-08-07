# Notifier
## Build 
```sh
go build .
```

## Send email

```sh
go run . sendEmail
```
```sh
go run . sendMessage
```

`.env`
```conf
MAIL_HOST=localhost
MAIL_PORT=1025
MAIL_USERNAME=jonas@mailhog.local
MAIL_PASSWORD=password
MAIL_TO=test@mailhog.local
MAIL_BODY="Lorem Ipsum is simply dummy text of the printing and typesetting industry."
MAIL_SUBJECT="Greetings"
MAIL_FROM=jonas@mailhog.local
TG_TOKEN=
TG_CHAT_ID=
TG_MESSAGE="Lorem Ipsum is simply dummy text of the printing and typesetting industry."
```