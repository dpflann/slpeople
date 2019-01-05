FROM golang:latest 
WORKDIR /app
ADD . "/go/src/github.com/slpeople"
RUN cd "/go/src/github.com/slpeople"; go test -v -cover ./... && go build -o slpeople.app; cp slpeople.app /app
EXPOSE 3000
COPY start.sh start.sh
COPY index.html index.html
COPY static/ static/

ENTRYPOINT ["./start.sh"]
