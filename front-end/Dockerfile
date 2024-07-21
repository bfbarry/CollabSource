FROM node:lts-buster
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install -g serve
RUN npm install
COPY . .
EXPOSE 3000
RUN npm run build
CMD ["serve", "-s", "-l", "3000", "./build"]