###
# Copyright (c) 2015-2019 Mainflux
#
# Mainflux is licensed under an Apache license, version 2.0 license.
# All rights not explicitly granted in the Apache license, version 2.0 are reserved.
# See the included LICENSE file for more details.
###

FROM node:lts-alpine

WORKDIR /app
RUN npm install --unsafe-perm=true --allow-root -g elm

COPY . /app
RUN elm make --optimize src/Main.elm

ENTRYPOINT ["elm", "reactor"]

