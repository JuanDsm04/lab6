FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY wait-for-it.sh /app/wait-for-it.sh

RUN chmod +x /app/wait-for-it.sh

RUN apt-get update && apt-get install -y netcat-openbsd

RUN go build -o main .

EXPOSE 8080

CMD ["/app/wait-for-it.sh", "db:5432", "--", "./main"]