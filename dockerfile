FROM ubuntu:latest

MAINTAINER Steampipe www.steampipe.io

RUN apt-get update && \
    apt-get install -y curl && \
    useradd -ms /bin/bash steampipe

COPY $GITHUB_WORKSPACE / 

RUN /bin/bash /install.sh

USER steampipe

WORKDIR /home/steampipe

RUN steampipe plugin install steampipe

ENTRYPOINT [ "/usr/local/bin/steampipe" ]