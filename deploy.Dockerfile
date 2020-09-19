# Probably won't be required.
FROM node:latest

WORKDIR /app

COPY ./ /app

RUN npm i -g serverless serverless-go-plugin && \
    sls deploy
