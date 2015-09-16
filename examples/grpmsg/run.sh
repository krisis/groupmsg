#! /bin/bash

./grpmsg server ":5555" & server1=$!
./grpmsg server ":5556" & server2=$!
sleep 4
./grpmsg client ":5555" ":5556" & client=$!
sleep 2

kill $server1 $server2 $client
