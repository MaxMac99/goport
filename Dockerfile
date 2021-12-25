FROM golang:1.17 AS build

ARG ACCESS_TOKEN_USR=$GITHUB_USER
ARG ACCESS_TOKEN_PWD=$ACCESS_TOKEN_PWD

WORKDIR /go/src
COPY . .

RUN echo "machine gitlab.com\n\tlogin $ACCESS_TOKEN_USR\n\tpassword $ACCESS_TOKEN_PWD" >> ~/.netrc

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o main .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/main ./
EXPOSE 9212/tcp
CMD [ "./main" ]
