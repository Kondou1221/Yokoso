FROM golang:1.22-bullseye as base

ENV ROOT=/go/src/go_api
ENV CGO_ENABLED 0
ENV GOOS=linux
WORKDIR ${ROOT}

FROM base as dev

EXPOSE 8080

COPY src/go.mod src/go.sum ./
RUN go mod download

FROM base as builder

# RUN apk update && apk add git
COPY  ./src .

RUN go mod download

RUN go build -o ./yokoso_api


FROM scratch as prod

WORKDIR /app

COPY --from=builder /go/src/go_api/yokoso_api ./

EXPOSE 8080
CMD ["./yokoso_api"]
