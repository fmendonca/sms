FROM golang:1.15.5-alpine3.12

COPY . .

RUN apk update && \
    apk add --no-cache bash=5.0.17-r0 git=2.26.2-r0 openssh=8.3_p1-r0 curl=7.69.1-r1 && \
    chmod +x deps.sh && \
    sh deps.sh
    
EXPOSE 80


ENTRYPOINT [ "go","run","main.go" ]