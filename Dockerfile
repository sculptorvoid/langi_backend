FROM golang:1.18-buster AS build

ENV GOPATH=/
WORKDIR /src/
COPY ./ /src/

RUN go mod download; CGO_ENABLED=0 go build -o /langi ./cmd/main.go

FROM alpine:latest

COPY --from=build /langi /langi
COPY ./configs/ /configs/
COPY ./migrations/ /migrations/
COPY ./.env ./

RUN apk --no-cache add postgresql-client

CMD ["/langi"]