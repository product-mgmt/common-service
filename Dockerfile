FROM golang:1.21-alpine

WORKDIR /app

RUN apk add --no-cache make
COPY . .
RUN go mod download
CMD [ "make", "db-setup" ]

