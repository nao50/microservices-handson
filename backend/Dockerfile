# FROM golang:1.14.4-alpine3.12 AS build

# ENV GO111MODULE=on
# ENV CGO_ENABLED=1

# WORKDIR /

# COPY . /go/src/github.com/nao50/gen-stub-api-server

# RUN apk update && apk add --no-cache git musl-dev gcc make g++ file
# RUN cd /go/src/github.com/nao50/gen-stub-api-server/c　md && go build -o bin/gen-stub-api-server main.go

# FROM alpine:3.12

# COPY --from=build /go/src/github.com/nao50/gen-stub-api-server/cmd/bin/gen-stub-api-server /usr/local/bin/

# CMD ["gen-stub-api-server"]

FROM golang:1.14.4-alpine3.12 AS build

ENV GO111MODULE=on
ENV CGO_ENABLED=1

WORKDIR /

COPY . /go/src/backend-server

RUN apk update && apk add --no-cache git musl-dev gcc make g++ file
RUN cd /go/src/backend-server && go build -o bin/backend-server main.go

FROM alpine:3.12

COPY --from=build /go/src/backend-server/bin/backend-server /usr/local/bin/

CMD ["backend-server"]
