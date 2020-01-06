FROM golang
MAINTAINER Vip Consult Solutions <team@vip-consult.solutions>


ADD /home/jenkins/workspace/cd-demo /go/src/cd-demo
RUN go install cd-demo
CMD /go/bin/cd-demo

EXPOSE 8080
