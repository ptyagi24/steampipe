FROM ubuntu:20.04

MAINTAINER Steampipe www.steampipe.io

RUN apt-get update && \
    useradd -ms /bin/bash steampipe

USER steampipe
WORKDIR /home/steampipe

RUN tar -xf ~/artifacts/linux.tar.gz -C /usr/local/bin/steampipe

ENTRYPOINT [ "/usr/local/bin/steampipe" ]

CMD ["query"]