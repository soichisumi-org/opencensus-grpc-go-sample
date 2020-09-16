#!/bin/bash

for i in $(seq 10)
do
   echo '{"message": "hello"}' | evans -r cli call grpctesting.EchoService.Echo -p 8080
done