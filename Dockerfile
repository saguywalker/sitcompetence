FROM node:10
COPY ./ui/admin-sc /app
WORKDIR /app
RUN yarn install && yarn build

FROM nginx
RUN mkdir /admin
COPY /ui/admin-sc/dist/ /admin
COPY nginx.conf /etc/nginx/nginx.conf