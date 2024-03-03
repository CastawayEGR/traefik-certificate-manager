FROM golang:1.22-alpine
WORKDIR /go/src/tcm
ENV CGO_ENABLED=0 \
    BIN_NAME='tcm'
COPY . .
RUN apk add --update gcc musl-dev
RUN go mod init \
    && go mod tidy \
    && go test -v -cover ./utils/... \
    && GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.appVersion=devel' -X 'main.buildTime=$(date +'%Y-%m-%d %T%z' -u)' -X 'main.gitCommit=local'" -trimpath -o /${BIN_NAME}-amd64-linux
