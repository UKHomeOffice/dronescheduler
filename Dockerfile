FROM alpine:3.3
MAINTAINER Ivan Pedrazas <ipedrazas@gmail.com>

RUN apk add --update bash curl \
    && rm -rf /var/cache/apk/*

COPY run.sh /usr/bin/run
RUN chmod +x /usr/bin/run
COPY dronescheduler /dronescheduler
RUN chmod +x /dronescheduler


CMD ["run"]
