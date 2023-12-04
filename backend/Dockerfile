FROM golang:alpine
LABEL authors="hogan"

WORKDIR /app
COPY . .

RUN go build -o code-editor

EXPOSE 8080

CMD ["./code-editor"]
