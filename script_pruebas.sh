#!/bin/bash
for (( var=1; var<=50; var++ ))
do 
    docker run -t -v `pwd`:/test prueba_alpine3.12
done