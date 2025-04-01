ARG NODE_VER="latest"
FROM node:${NODE_VER}

RUN npm install -g npm

USER node

WORKDIR /home/node/app

RUN mkdir next-app

WORKDIR /home/node/app/next-app

RUN mkdir node_modules
RUN chown node:node -R node_modules

COPY --chown=node:node ./next-app .

RUN npm install
