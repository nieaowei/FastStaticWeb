FROM centos:latest
ADD . /home
RUN ls && ls /home && mkdir /blog
WORKDIR /home
RUN pwd && ls && ls /
ENTRYPOINT ["./bin/faststaticweb-linux"]