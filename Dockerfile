FROM golang:latest


ADD . /go/src/cd-demo
RUN go install cd-demo
CMD /go/bin/cd-demo

RUN cd $GOPATH/src/k8s.io/klog && git checkout v0.4.0 && cd -
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

EXPOSE 8080
