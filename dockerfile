FROM golang:alpine
RUN mkdir /swallow
COPY ./ /swallow
WORKDIR /swallow
RUN go build -o swallow swallow/cmd/main.go
ENTRYPOINT ["/swallow/swallow"]