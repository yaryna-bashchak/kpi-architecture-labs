#!/bin/bash

startPos=0
finishPos=400
pos=$startPos
step=10
interval=0.01

curl -X POST http://localhost:17000 -d "figure $startPos $startPos"

curl -X POST http://localhost:17000 -d "update"

sleep $interval

while true; do
  while ((pos < finishPos)); do
    curl -X POST http://localhost:17000 -d "move $step $step"
    pos=$((pos + step))
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
  done

  while ((pos > startPos)); do
    curl -X POST http://localhost:17000 -d "move $((-step)) $((-step))"
    pos=$((pos - step))
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
  done
done
