#
# Build and bundle the Vue.js SPA
#
FROM node:10-alpine as vue-build
WORKDIR /build

COPY ui/admin-sc  ./

RUN yarn install && yarn build

#
# Build Golang Gorrila server
#

FROM golang:1.12-alpine as go-build
WORKDIR /build/src/server

RUN apk update && apk add git gcc musl-dev

COPY api/*.go ./
COPY app/*.go ./
COPY cmd/*.go ./
COPY db/*.go ./
COPY model/*.go ./
COPY main.go ./
COPY config.yaml ./

FROM scratch
WORKDIR /app
LABEL org.label-schema.name="sitcom-demoapp" \
      org.label-schema.description="Demonstration Vue.js and Go web app" \
      org.label-schema.version="1.5.2"

COPY --from=vue-build /build/dist .
COPY --from=go-build /build/src/server/main.go .

ENV PORT 3000
EXPOSE 3000
CMD ["/app/main", "serve"]
