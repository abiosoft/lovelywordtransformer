FROM alpine:3.4

COPY app /app

EXPOSE 80

ENTRYPOINT ["/app"]