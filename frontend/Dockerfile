FROM node:10.15.3 as build-stage
WORKDIR /app
RUN ls
RUN echo $PWD
EXPOSE 8080
COPY ./frontend .
RUN ls
RUN yarn
RUN yarn build

FROM abiosoft/caddy:latest as production-stage
WORKDIR /app
COPY --from=build-stage ./app/dist /app
EXPOSE 80
# EXPOSE 443
# COPY nginx.conf /etc/nginx/nginx.conf
# COPY nginx.crt /etc/nginx/nginx.crt
# COPY nginx.key /etc/nginx/nginx.key
