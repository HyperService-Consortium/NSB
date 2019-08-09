FROM alpine:3.7
MAINTAINER Myriad-Dreamin <camiyoru@gmail.com>

RUN apk update && \
    apk upgrade && \
    apk --no-cache add curl jq file

VOLUME [ /tendermint ]
WORKDIR /tendermint
EXPOSE 26656 26657
ENTRYPOINT ["/usr/bin/wrapper.sh"]

STOPSIGNAL SIGTERM

COPY tendermint /usr/bin
COPY NSB /usr/bin/
COPY wrapper.sh /usr/bin/wrapper.sh

