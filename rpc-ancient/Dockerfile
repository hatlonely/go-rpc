FROM centos:centos7

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" >> /etc/timezone

COPY docker/ /var/docker/go-rpc-ancient
RUN mkdir -p /var/docker/go-rpc-ancient/log

WORKDIR /var/docker/go-rpc-ancient
CMD [ "bin/ancient", "-c", "config/go-rpc-ancient.json" ]