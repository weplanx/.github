FROM alpine:edge

WORKDIR /app

RUN apk --no-cache add tzdata

ADD server /app/
ADD model/*.json /app/model/
COPY templates /app/templates

EXPOSE 3000

CMD ["./server"]
