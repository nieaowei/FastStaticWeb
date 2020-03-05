FROM centos:latest
RUN mkdir /blog
WORKDIR /blog
ENTRYPOINT ["./faststaticweb-linux"]