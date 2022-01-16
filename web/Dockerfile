FROM node:alpine
WORKDIR /docker-compose-basics/web
COPY . .
RUN npm install --prod
CMD ["node", "index.js"]