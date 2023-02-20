FROM golang:1.18.2-alpine AS build
# ENV CGO_ENABLED=0
RUN set -ex &&\
    apk add --no-progress --no-cache \
      gcc \
      musl-dev
RUN mkdir /build
ADD . /build/
WORKDIR /build

# ENV GOARCH=amd64
RUN go build -o main -a  -tags musl .

FROM alpine

COPY --from=build build/main /app/

WORKDIR /app/
CMD ["./main"]