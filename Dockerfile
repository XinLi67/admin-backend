FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
ADD alumni /app/alumni
ADD .env /app/.env

CMD ["/app/alumni"]
EXPOSE 9000