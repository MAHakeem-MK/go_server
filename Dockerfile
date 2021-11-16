FROM golang

WORKDIR go_server

RUN go get github.com/gofiber/fiber/v2
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/sqlite
RUN swag init
RUN go get -u github.com/arsmn/fiber-swagger/v2

COPY . .

EXPOSE 3000

CMD ["go","run","main.go"]