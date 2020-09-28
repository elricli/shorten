FROM golang:latest

WORKDIR /app

ENV GO111MODULE=on
ENV GOPROXY=http://goproxy.cn,direct

EXPOSE 80

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -o shorten cmd/main.go

FROM scratch

WORKDIR /root/

COPY --from=0 /app/shorten .

CMD ["./shorten"]
