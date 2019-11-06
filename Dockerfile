FROM golang:1.13-alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN apk add --update git
RUN go get -v .
COPY ./config.yaml ./
RUN go build -o sitcom
CMD ["/app/sitcom"]
