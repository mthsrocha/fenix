FROM golang:1.16 as builder

RUN apk --update upgrade && \
    apk add build-base gcc sqlite && \
    rm -rf /var/cache/apk/*


RUN mkdir -p $GOPATH/src/github.com/mthsrocha/fenix

COPY . $GOPATH/src/github.com/mthsrocha/fenix/

RUN CC=gcc CGO_ENABLED=1 go build -o $GOPATH/bin/fenix $GOPATH/src/github.com/mthsrocha/fenix/fenix.go

ENV PORT=5001 DATABASE_DRIVE=mysql DATABASE_HOST=127.0.0.1

EXPOSE 5001

COPY --from=builder /go/bin/fenix /bin/

CMD ["fenix", "run"]


# RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
# ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
# WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
# RUN go build -o fenix .
## Our start command which kicks off
## our newly created binary executable
# CMD ["/app/fenix"]