FROM node:boron-alpine

COPY *.proto mqtt/* ./

RUN npm rebuild && npm install

EXPOSE 1883 8880

CMD ["node", "mqtt.js"]
