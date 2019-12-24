FROM golang:alpine as builder
LABEL maintainer="saguywalker"
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cho -o sitcom .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /app/sitcom .
COPY --from=builder /app/config.yaml .
EXPOSE 3000
CMD ["./sitcom", "dev"]