FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN make build

EXPOSE 5050
CMD ["./bin/one_liner"]
