API_NAME := hideki
APP_KEY=
MODE=Development
# Email
MAIL_MAILER=smtp
MAIL_HOST=smtp.mailtrap.io
MAIL_PORT=2525
MAIL_USERNAME=
MAIL_PASSWORD=
MAIL_ENCRYPTION=tls
# Database
DATABASE_URL=postgresql://postgres:123456@192.168.100.47/pollos_hermanos
DATABASE_MAX_OPEN_CONNECTIONS=25
DATABASE_MAX_IDDLE_CONNECTIONS=25
DATABASE_MAX_IDDLE_TIME=15m
# HTTP
HTTP_SERVER_IDLE_TIMEOUT=60s
PORT=8080
HTTP_SERVER_READ_TIMEOUT=1s
HTTP_SERVER_WRITE_TIMEOUT=2s


build:
	env CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -ldflags '-w -s' -a -installsuffix cgo -o bin/$(API_NAME) ./cmd/main.go

run_local:
	env APP_KEY=$(APP_KEY) MODE=$(MODE) DATABASE_URL=$(DATABASE_URL) PORT=$(PORT) ./bin/$(API_NAME)

clean:
	rm -fr ./bin

db_up:
	migrate -path=./migrations -database=$(DATABASE_URL) up

db_down:
	migrate -path=./migrations -database=$(DATABASE_URL) down