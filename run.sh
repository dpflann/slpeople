#!/bin/bash
apikey="$1";
port="$1";
docker run --rm -it -p $port:$port slpeople "$apikey" "$port"
