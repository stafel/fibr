###
### Build Container
###

FROM golang:1.20-alpine as build
COPY . /build
WORKDIR /build
RUN go build

###
### Runtime Container
###

FROM alpine:latest
COPY --from=build /build/fibr /usr/bin/fibr
