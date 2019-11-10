FROM alpine:3.7
MAINTAINER Myriad-Dreamin <camiyoru@gmail.com>

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && \
    apk --no-cache add curl jq file

VOLUME /tendermint
WORKDIR /tendermint
EXPOSE 26656 26657
# ENTRYPOINT ["/bin/sh"]
ENTRYPOINT ["/usr/bin/wrapper.sh"]

STOPSIGNAL SIGTERM

# COPY tendermint /usr/bin
COPY NSB /usr/bin/
COPY wrapper.sh /usr/bin/wrapper.sh
RUN dos2unix /usr/bin/wrapper.sh
