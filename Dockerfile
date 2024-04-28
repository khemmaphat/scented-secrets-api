
FROM  golang:1.21.4

WORKDIR /go/src/app
RUN apt update && apt install -y nocache --force-yes --no-install-recommends gcc g++ git make git libc-dev binutils-gold
COPY . . 

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download

RUN go build main.go

EXPOSE 8080

CMD ["./main"]