FROM alpine:3.9

# Installs latest Chromium  package.
RUN apk update && apk upgrade && \
    echo @edge http://nl.alpinelinux.org/alpine/edge/community >> /etc/apk/repositories && \
    echo @edge http://nl.alpinelinux.org/alpine/edge/main >> /etc/apk/repositories && \
    apk add --no-cache \
      chromium@edge=~73.0.3683.103 \
      nss@edge \
      freetype@edge \
      harfbuzz@edge \
      ttf-freefont@edge \
      dumb-init # hepls to fight chrome zombie processes

WORKDIR /app
ADD bin/main /app
RUN mkdir /store

ENTRYPOINT ["dumb-init", "--"]
CMD ["./main"]