FROM golang:alpine AS build-env
RUN apk add g++
COPY . /src
RUN apk update && apk add --no-cache git openssh
RUN cd /src && go build ./cmd/second-app

FROM alpine
WORKDIR /app
COPY --from=build-env /src/second-app /app/
ENTRYPOINT ./second-app
