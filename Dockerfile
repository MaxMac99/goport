# syntax = docker/dockerfile:1.3
FROM golang:1.17 AS build

ARG ACCESS_TOKEN_USR=$GITHUB_USER
ARG ACCESS_TOKEN_PWD=$ACCESS_TOKEN_PWD

WORKDIR /go/src
COPY . .

ENV CGO_ENABLED=0
RUN --mount=type=secret,id=netrc,dst=/root/.netrc go get -d -v ./...

RUN go build -a -installsuffix cgo -o main .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/main ./
EXPOSE 9212/tcp
CMD [ "./main" ]
