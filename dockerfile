FROM ubuntu:20.04

MAINTAINER Steampipe www.steampipe.io

COPY $GITHUB_WORKSPACE / 

RUN tar -xf /linux.tar.gz -C /usr/local/bin/steampipe

ENTRYPOINT [ "/usr/local/bin/steampipe" ]

CMD ["query"]