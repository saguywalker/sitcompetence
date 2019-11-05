#FROM node:latest as build-stage
#WORKDIR /app
#RUN ls
#RUN echo $PWD
#COPY ../package*.json ./
#RUN ls
#RUN yarn
#COPY ../ ./
#RUN ls
#RUN yarn build

#FROM nginx:latest as production-stage
#RUN mkdir /app
#RUN ls 
#COPY --from=build-stage ./frontend/dist /app
#EXPOSE 80
#COPY nginx.conf /etc/nginx/nginx.conf

