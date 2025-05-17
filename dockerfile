FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o /assistant .

EXPOSE 3000

CMD ["/assistant"]