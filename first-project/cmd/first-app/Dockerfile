FROM golang:alpine AS build-env
RUN apk add g++
COPY . /src
RUN apk update && apk add --no-cache git openssh
RUN cd /src && go build ./cmd/first-app

FROM alpine
WORKDIR /app
COPY --from=build-env /src/first-app /app/
ENTRYPOINT ./first-app
