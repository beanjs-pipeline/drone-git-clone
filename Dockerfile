FROM golang:1.16-alpine as builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOPROXY="https://goproxy.cn"

WORKDIR /go/src
COPY . .

RUN go build -a -o /go/bin/drone-git-clone ./cli

FROM alpine:3.14
RUN apk add --no-cache git

WORKDIR /drone/src
COPY --from=builder /go/bin/drone-git-clone .
ENTRYPOINT [ "/go/bin/drone-git-clone" ]