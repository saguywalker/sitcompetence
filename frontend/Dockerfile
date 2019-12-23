FROM node:10.15.3 as build-stage
WORKDIR /app
EXPOSE 8080
COPY ./ .
RUN yarn
RUN yarn build

FROM nginx:latest as production-stage
RUN mkdir /app
RUN ls
COPY --from=build-stage ./app/dist /app
EXPOSE 80
COPY nginx.conf /etc/nginx/nginx.conf
