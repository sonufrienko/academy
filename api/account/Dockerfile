FROM node:14-alpine as build

# adds deps for node-gyp: add if native modules are used
# RUN apk update && apk upgrade \
#   && apk --no-cache add --virtual builds-deps build-base python

# set app basepath
ENV HOME=/app

# add app dependencies
RUN mkdir -p $HOME
COPY package.json $HOME
COPY package-lock.json $HOME

# change workgin dir and install deps in quiet mode
WORKDIR $HOME
RUN npm ci -q

# copy all app files
COPY . $HOME

# compile typescript and build all production stuff
RUN npm run build

# remove dev dependencies and files that are not needed in production
RUN npm prune --production

# start new image for lower size
FROM node:14-alpine

# create use with no permissions
RUN addgroup -g 101 -S app && adduser -u 100 -S -G app -s /bin/false app

# set app basepath
ENV HOME=/app

# copy production complied node app to the new image
RUN mkdir -p $HOME
COPY --from=build $HOME $HOME
RUN chown -R app:app $HOME/*

# run app with low permissions level user
USER app
WORKDIR $HOME

EXPOSE 3000

ENV NODE_ENV=production

CMD ["node", "./build"]