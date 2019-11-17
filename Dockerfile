FROM golang:1.13-alpine3.10 AS build
# Support CGO and SSL
WORKDIR ./cmd
RUN apk update && apk add --no-cache gcc g++ make python3 py-pip postgresql-dev python3-dev musl-dev
RUN pip3 install --upgrade pip
RUN pip3 install --upgrade pip setuptools
RUN pip3 install psycopg2
RUN apk add git
RUN echo $PWD
COPY . .
RUN ls
#RUN python3 initdb.py
RUN go build -o sitcom .
#FROM alpine:3.10
#RUN apk --no-cache add ca-certificates
EXPOSE 3000
ENTRYPOINT ./sitcom serve

