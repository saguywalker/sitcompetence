FROM node:10.15.3 as build-stage
WORKDIR /app
COPY ./frontend/package*.json ./
RUN yarn
COPY ./frontend/ ./
RUN yarn build

FROM nginx:latest
RUN mkdir /app
COPY --from=build-stage /app/dist /app
RUN ls -l /app 
RUN rm -rf /etc/nginx/conf.d/default.conf
COPY ./frontend/default.conf /etc/nginx/conf.d
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
