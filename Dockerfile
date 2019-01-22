FROM golang:latest 
COPY . /app
ENV GOPATH=/app
WORKDIR /app
RUN CGO_ENABLED=1 GOOS=linux go build -o app main.go
EXPOSE 8080
CMD ["./app"]