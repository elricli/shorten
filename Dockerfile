FROM golang:latest

WORKDIR /app

ENV GO111MODULE=on
ENV GOPROXY=http://goproxy.cn,direct

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -o shorten cmd/main.go

FROM scratch

WORKDIR /root/

COPY --from=0 /app/shorten .
COPY --from=0 /app/data .
COPY --from=0 /app/content ./content
COPY --from=0 /app/config.yml .

CMD ["./shorten"]
