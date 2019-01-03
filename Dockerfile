FROM golang:latest 
WORKDIR /go
ADD . "/go/src/github.com/slpeople"
RUN cd "/go/src/github.com/slpeople"; go build -o slpeople.app; cp slpeople.app /go
EXPOSE 3000
ADD start.sh start.sh
RUN ls

ENTRYPOINT ["./start.sh"]
