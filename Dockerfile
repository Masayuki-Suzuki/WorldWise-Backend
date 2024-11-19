FROM golang:1.23.2-alpine

ENV TZ /usr/shara/zoneinfo/America/Vancouver

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

CMD ["pwd"]

COPY ./app .
EXPOSE 4000

RUN apk upgrade --update && \
    apk --no-cache add git

RUN go mod download && go mod tidy
RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]
