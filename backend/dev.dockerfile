FROM golang:1.22-bullseye as dev

ENV ROOT=/go/src/go_api
ENV CGO_ENABLED 0
ENV GOOS=linux
WORKDIR ${ROOT}

COPY src/go.mod src/go.sum ./
RUN go mod download
EXPOSE 8080

CMD ["go", "run", "main.go"]

# FROM golang:1.22-bullseye as builder

# ENV ROOT=/go/src/app
# WORKDIR ${ROOT}

# # RUN apk update && apk add git
# COPY go.mod go.sum ./
# RUN go mod download

# COPY . ${ROOT}
# RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary


# FROM scratch as prod

# ENV ROOT=/go/src/app
# WORKDIR ${ROOT}
# COPY --from=builder ${ROOT}/binary ${ROOT}

# EXPOSE 8080
# CMD ["/go/src/app/binary"]