FROM golang:1.21.1

# Enviroment
ENV GO111MODULE=auto
ENV APP_ENV docker

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Biuld
RUN go build -o ./bin/main .

COPY config-docker.json /app/bin

WORKDIR /app/bin

# Run
CMD ["/app/bin/main"]