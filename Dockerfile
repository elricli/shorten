FROM alpine

WORKDIR /app

COPY cmd/shorten .

# Default environment
# Use -e SHORTEN_DSN=xx to override default
ENV SHORTEN_DSN="user='postgres' password='' host='postgres' port=5432 dbname='shorten' sslmode=disable options='-c statement_timeout=60000'"
ENV SHORTEN_ADDR="shorten:80"

CMD ["/app/shorten"]
