FROM node:20-alpine3.17 as builder

RUN npm install -g pm2

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm ci --only=production

COPY . . 

RUN npm run build 

EXPOSE 5000

CMD ["pm2-runtime", "start", "ecosystem.config.js", "--env", "production"]
