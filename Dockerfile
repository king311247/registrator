FROM  golang:1.15.3-alpine3.12 AS builder
WORKDIR /go/src/github.com/king311247/registrator/
COPY . .

RUN apk add --no-cache curl git \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN export CGO_ENABLED=0 \
    && export GO111MODULE=on \
    && go build -a -installsuffix cgo -ldflags "-X main.Version=$(cat VERSION)" -o bin/registrator

FROM alpine:3.12
RUN apk add --no-cache ca-certificates \
        && echo "hosts: files dns" > /etc/nsswitch.conf
COPY --from=builder /go/src/github.com/king311247/registrator/conf /root/conf
COPY --from=builder /go/src/github.com/king311247/registrator/bin/registrator /bin/registrator

ENTRYPOINT ["/bin/registrator"]