FROM centos:7

RUN curl -OL https://golang.google.cn/dl/go1.15.2.linux-amd64.tar.gz && \
    tar -C /usr/local -zxf go1.15.2.linux-amd64.tar.gz && \
    mkdir -p /go/src /go/bin /go/pkg && \
    chmod -R 777 /go && \
    rm -fr go1.15.2.linux-amd64.tar.gz

RUN yum -y update && \
    yum install -y epel-release && \
    yum install -y chromium

RUN yum install -y wget git mkfontscale fontconfig&& \
    cd /usr/share/fonts && \
    git clone --progress --verbose https://github.com/lampnick/free-fonts.git && \
    mv free-fonts/* ./ && \
    mkfontscale && \
    mkfontdir && \
    fc-cache && \
    yum clean all && \
    rm -fr /var/cache/yum/* && \
    fc-list :lang=zh && \
    chromium-browser --version

ENV GOLANG_VERSION 1.15.2
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on

#ADD DOCTRON
RUN mkdir -p /doctron
COPY . /doctron
RUN cd /doctron && \
    go build && \
    cp -fr doctron /usr/local/bin && \
    chmod +x /usr/local/bin/doctron && \
    rm -fr /doctron/*

CMD doctron