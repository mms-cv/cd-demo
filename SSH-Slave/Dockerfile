FROM jenkins/ssh-slave

## Add Docker Client

ENV DOCKERVERSION=18.06.1-ce
RUN curl -fsSLO https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKERVERSION}.tgz \
  && tar xzvf docker-${DOCKERVERSION}.tgz --strip 1 \
                 -C /usr/local/bin docker/docker \
  && rm docker-${DOCKERVERSION}.tgz

## Setup SSH

EXPOSE 22

ENTRYPOINT ["setup-sshd"]