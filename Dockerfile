FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
ADD adcenter /app/adcenter
ADD .env /app/.env

CMD ["/app/adcenter"]
EXPOSE 9000