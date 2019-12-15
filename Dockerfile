FROM golang:1.13-alpine3.10 AS build
# Support CGO and SSL
WORKDIR ./cmd
RUN apk update && apk add --no-cache gcc g++ make postgresql-dev musl-dev git
COPY . .
RUN ls
#RUN python3 initdb.py
RUN go build -o sitcom .
#FROM alpine:3.10
#RUN apk --no-cache add ca-certificates
EXPOSE 3000
ENTRYPOINT ./sitcom dev

