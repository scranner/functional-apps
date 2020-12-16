FROM mhart/alpine-node:13

WORKDIR /opt/subtract/
COPY package*.json ./
RUN npm install --production
COPY . ./
EXPOSE 80
ENTRYPOINT node index.js