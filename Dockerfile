FROM golang:latest

WORKDIR /app

COPY . .

RUN GOPROXY=http://goproxy.cn,direct CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -o shorten cmd/main.go cmd/wire_gen.go

FROM scratch

WORKDIR /root/

COPY --from=0 /app/shorten .
COPY --from=0 /app/public ./public

CMD ["./shorten"]
