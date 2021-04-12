FROM ubuntu:20.04

MAINTAINER Steampipe www.steampipe.io

COPY $GITHUB_WORKSPACE / 

RUN /bin/bash /install.sh

ENTRYPOINT [ "/usr/local/bin/steampipe" ]

CMD ["query"]