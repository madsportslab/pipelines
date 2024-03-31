FROM golang:latest as builder
WORKDIR /sources
COPY . .
RUN go build

FROM ubuntu:latest
WORKDIR /usr/local/pipelines
COPY --from=builder /sources/pipelines .
EXPOSE 8686
CMD ["/usr/local/pipelines/pipelines"]
