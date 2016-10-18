FROM golang:1.7.1

WORKDIR /app
COPY . /app

EXPOSE 3000

ENTRYPOINT ["./main"]