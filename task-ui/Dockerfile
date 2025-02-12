FROM node:22-alpine as build


WORKDIR /opt/app

COPY package*.json .

RUN npm install

COPY . .

RUN npm run build

FROM nginx:alpine

COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=build /opt/app/dist/task-ui/browser/ /usr/share/nginx/html

EXPOSE 80
