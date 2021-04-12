FROM ubuntu:20.04

MAINTAINER Steampipe www.steampipe.io

RUN useradd -ms /bin/bash steampipe

USER steampipe

RUN tar -xf ~/artifacts/linux.tar.gz -C /usr/local/bin/steampipe

ENTRYPOINT [ "/usr/local/bin/steampipe" ]

CMD ["query"]