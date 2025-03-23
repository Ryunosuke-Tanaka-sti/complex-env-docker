ARG NODE_VER
FROM node:${NODE_VER}

RUN npm install -g npm@11.0.0
RUN npm install -g @azure/static-web-apps-cli

USER node

WORKDIR /home/node/app

RUN mkdir next-app

WORKDIR /home/node/app/next-app

RUN mkdir node_modules
RUN chown node:node -R node_modules

COPY --chown=node:node ./next-app .

RUN npm install
