FROM ubuntu:20.04

MAINTAINER Steampipe www.steampipe.io

RUN apt-get update && \
    apt-get install -y curl \
    useradd -ms /bin/bash steampipe

COPY $GITHUB_WORKSPACE / 

USER steampipe

RUN /bin/bash /install.sh

ENTRYPOINT [ "/usr/local/bin/steampipe" ]

CMD ["query"]