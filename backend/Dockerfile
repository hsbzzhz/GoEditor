FROM golang:alpine
LABEL authors="hogan"

WORKDIR /app
COPY . .

RUN go build main.go

ENTRYPOINT ["top", "-b"]