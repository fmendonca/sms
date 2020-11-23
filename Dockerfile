FROM golang:1.15.5-alpine3.12

COPY . $GOPATH/app

WORKDIR $GOPATH/app

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN chmod +x $GOPATH/app/deps.sh
RUN sh /$GOPATH/app/deps.sh

CMD ["go", "run", "main.go"]