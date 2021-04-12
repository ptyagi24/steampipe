FROM ubuntu:20.04

MAINTAINER Steampipe www.steampipe.io

RUN apt-get update && \
    apt-get install -y curl

COPY $GITHUB_WORKSPACE / 

RUN /bin/bash /install.sh

ENTRYPOINT [ "/usr/local/bin/steampipe" ]

CMD ["query"]