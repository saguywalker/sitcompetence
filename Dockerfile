FROM golang:1.13-alpine3.10 AS build
# Support CGO and SSL
WORKDIR ./cmd
RUN apk --no-cache add gcc g++ make
RUN apk add git
RUN echo $PWD
COPY . .
RUN ls
RUN go build -o sitcom .
#FROM alpine:3.10
#RUN apk --no-cache add ca-certificates
EXPOSE 3000
ENTRYPOINT ./sitcom serve

FROM node:latest as build-stage
WORKDIR /app
RUN ls
RUN echo $PWD
COPY ./frontend/package*.json ./                                                                                                                                                 
RUN ls
RUN yarn
RUN ls
COPY ./frontend .
RUN ls
RUN yarn build

FROM nginx:latest as production-stage
RUN mkdir /app
RUN ls
COPY --from=build-stage ./dist /app
EXPOSE 80
COPY nginx.conf /etc/nginx/nginx.conf
