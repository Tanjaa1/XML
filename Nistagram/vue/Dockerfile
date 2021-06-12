FROM node:12.18.1-alpine AS
RUN npm install -g https-server
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build
EXPOSE 3000
CMD ["https-server","dist"]