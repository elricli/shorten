FROM golang:latest

WORKDIR /app

COPY . .

RUN cd cmd && GOPROXY=http://goproxy.cn,direct CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -o shorten

FROM scratch

WORKDIR /root/

COPY --from=0 /app/cmd/shorten .
COPY --from=0 /app/public ./public

CMD ["./shorten"]
