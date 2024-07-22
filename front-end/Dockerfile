FROM node:lts-buster
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install -g serve
RUN npm install
COPY . .
EXPOSE 3000
ENV REACT_APP_NODE_PROXY_URL=https://node-proxy-4jjvxz6spq-uc.a.run.app
RUN npm run build
CMD ["serve", "-s", "-l", "3000", "./build"]