FROM golang:latest as build-env

RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
RUN go mod init && go build


FROM gcr.io/distroless/base

COPY --from=build-env /go/src/app /
CMD ["/app"]