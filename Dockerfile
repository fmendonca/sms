FROM golang:1.15.5-alpine3.12

COPY . "$GOPATH"/app

WORKDIR "$GOPATH"/app

RUN apk update && \
    apk add --no-cache bash=5.0.17-r0 git=2.26.2-r0 openssh=8.3_p1-r0 curl=7.69.1-r1 && \
    chmod +x "$GOPATH"/app/deps.sh && \
    sh /"$GOPATH"/app/deps.sh && \ 
    go build "$GOPATH"/app/main.go

CMD [""$GOPATH"/app/sms"]