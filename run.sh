#!/bin/bash
# TODO: Add better flag handling through getopts
apikey="$1";
port="$2";
docker run --rm -it -p $port:$port slpeople "$apikey" "$port"
