FROM node:10.15.3 as build-stage
WORKDIR /app
RUN ls
RUN echo $PWD
COPY ./frontend/package*.json ./                                                                                                                                                 
RUN ls
RUN npm install
RUN ls
COPY ./frontend .
RUN ls
RUN npm run build

FROM nginx:latest as production-stage
RUN mkdir /app
RUN ls
COPY --from=build-stage ./app/dist /app
EXPOSE 80
EXPOSE 443
COPY nginx.conf /etc/nginx/nginx.conf
COPY nginx.crt /etc/nginx/nginx.crt
COPY nginx.key /etc/nginx/nginx.key
