FROM ubuntu:14.04

WORKDIR /app
COPY main /app/

EXPOSE 8000

ENTRYPOINT ["./main"]