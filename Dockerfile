FROM node:10
COPY ./ui/admin-sc /app
WORKDIR /app
RUN yarn install && yarn build

FROM nginx
RUN mkdir /app
COPY /ui/admin-sc/dist /app
COPY nginx.conf /etc/nginx/nginx.conf