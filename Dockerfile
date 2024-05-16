FROM golang

WORKDIR /go/src/app
COPY . .

RUN go mod tidy

RUN go build -o /go/bin/app

CMD ["/go/bin/app"]
