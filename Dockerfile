FROM golang:1.20.0

WORKDIR /usr/src/app

RUN apt-get update && apt-get install -y bash

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy

CMD ["bash"]