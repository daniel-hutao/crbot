FROM golang:alpine AS build-env
WORKDIR $GOPATH/src/github.com/daniel-hutao/crbot
COPY . .
RUN apk add git
RUN go get ./... && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/crbot/main.go

FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/daniel-hutao/crbot/app /app/
CMD ["./app"]
USER 1000
EXPOSE 8080/tcp
