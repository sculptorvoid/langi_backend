FROM golang:1.17-buster AS build

ENV GOPATH=/
WORKDIR /src/
COPY ./ /src/

# build go app
RUN go mod download; CGO_ENABLED=0 go build -o /langi ./cmd/main.go


FROM alpine:latest

# copy go app, config and wait-for-postgres.sh
COPY --from=build /langi /langi
COPY ./configs/ /configs/
COPY ./.env ./

# install psql and make wait-for-postgres.sh executable
RUN apk --no-cache add postgresql-client

CMD ["/langi"]