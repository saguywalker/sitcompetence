FROM node:10.15.3 as build-stage
WORKDIR /app
COPY ./frontend/package*.json ./
RUN yarn
COPY ./frontend/ ./
RUN yarn build

FROM abiosoft/caddy:latest as production-stage
WORKDIR /app
COPY --from=build-stage ./app/dist /app
COPY ./frontend/Caddyfile /etc/Caddyfile
RUN ls && pwd
EXPOSE 80
# EXPOSE 443
# COPY nginx.conf /etc/nginx/nginx.conf
# COPY nginx.crt /etc/nginx/nginx.crt
# COPY nginx.key /etc/nginx/nginx.key
