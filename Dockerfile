FROM alpine

WORKDIR /app

COPY cmd/shorten .

CMD ["/app/shorten"]
