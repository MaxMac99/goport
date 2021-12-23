FROM golang:1.10 AS build
WORKDIR /go/src
COPY . .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/openapi ./
EXPOSE 9212/tcp
ENTRYPOINT ["./openapi"]