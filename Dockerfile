FROM golang

WORKDIR /go/src/app
COPY . .

CMD ["make", "run"]
