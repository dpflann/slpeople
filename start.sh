#!/bin/bash
apikey="$1";
port="$2";
echo "apikey=$apikey, port=$port";
./slpeople.app --apikey "$apikey" --port "$port"
