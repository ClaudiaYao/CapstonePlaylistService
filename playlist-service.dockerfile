FROM golang:latest as builder


RUN mkdir /app
WORKDIR /app

ADD . ./app
COPY . .

COPY Makefile go.mod go.sum .env ./
COPY .env ./cmd/

# Builds your app with optional configuration

RUN go build -o playlist-service ./cmd

# Tells Docker which network port your container listens on
EXPOSE 8080

# Specifies the executable command that runs when the container starts
#RUN [ “/playlist-service”]
CMD [ "/app/playlist-service"]