FROM golang:latest

LABEL maintainer="Josh <lisyaoran51@hotmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 5000

RUN go build

CMD ["./GoSimpleApi"]