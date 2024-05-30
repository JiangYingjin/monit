FROM node:16

WORKDIR /gva_web/
COPY . .

ENV NODE_ENV=dev
RUN npm config set registry http://mirrors.cloud.tencent.com/npm/
RUN npm install
#RUN npm install \
#    && npm install -g serve \
#    && npm run build \
#    && rm -fr node_modules

EXPOSE 8080

#CMD [ "serve", "-s", "build" ]
#CMD [ "npm", "run", "serve", "--", "--host", "0.0.0.0" ]
CMD npm run serve

