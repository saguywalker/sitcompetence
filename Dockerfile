FROM node:latest as build-stage
WORKDIR /app
COPY ./frontend/package*.json ./
RUN yarn
COPY ./frontend .
RUN yarn build

FROM nginx as production-stage
RUN mkdir ./app
EXPOSE 80
EXPOSE 3000
EXPOSE 8080
COPY --from=build-stage /app/dist /app
COPY nginx.conf /etc/nginx/nginx.conf
