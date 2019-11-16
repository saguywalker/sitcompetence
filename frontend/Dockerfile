FROM node:10.15.3 as build-stage
WORKDIR /app
RUN ls
RUN echo $PWD
COPY ./ ./                                                                                                                                                 
RUN ls
EXPOSE 8080
COPY ./ .
RUN ls
RUN yarn
RUN yarn build 

FROM nginx:latest as production-stage
RUN mkdir /app
RUN ls
COPY --from=build-stage ./app/dist /app
EXPOSE 80
EXPOSE 443
COPY nginx.conf /etc/nginx/nginx.conf
COPY nginx.crt /etc/nginx/nginx.crt
COPY nginx.key /etc/nginx/nginx.key
