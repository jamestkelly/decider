FROM golang:latest
WORKDIR /app
COPY . /app
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]

# Metadata
LABEL author="Jim Kelly"