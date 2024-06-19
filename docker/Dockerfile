FROM node:current-alpine3.20

WORKDIR /app/

COPY . .

WORKDIR /app/frontend/

RUN yarn install

RUN yarn build

CMD [ "npx" "http-server" "/app/build/frontend/index.html" "-P" "http://localhost:8080?"]