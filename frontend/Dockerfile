FROM node:alpine as build-stage
WORKDIR /app
EXPOSE 8080
COPY ./frontend .
RUN ls -a | grep .env
RUN yarn
RUN yarn build

FROM abiosoft/caddy:latest as production-stage
WORKDIR /app
COPY ./frontend/Caddyfile /etc/Caddyfile
COPY --from=build-stage ./app/dist /app
RUN ls && pwd
EXPOSE 80
# EXPOSE 443
# COPY nginx.conf /etc/nginx/nginx.conf
# COPY nginx.crt /etc/nginx/nginx.crt
# COPY nginx.key /etc/nginx/nginx.key
