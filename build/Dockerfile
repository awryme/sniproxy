FROM golang:1.22 AS build
COPY ./ /go/src/sniproxy
WORKDIR /go/src/sniproxy
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
RUN go build -v -o /bins/sniproxy ./cmd/sniproxy

FROM alpine
COPY --from=build /bins/sniproxy /bins/sniproxy

ENTRYPOINT [ "/bins/sniproxy" ]