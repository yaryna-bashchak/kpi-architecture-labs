#!/bin/bash

topLeft=75
bottomLeft=325
bottomRight=325
topRight=75
posX=$topLeft
posY=$topLeft
step=10
interval=0.01

curl -X POST http://localhost:17000 -d "figure $posX $posY"

curl -X POST http://localhost:17000 -d "update"

sleep $interval

while true; do
  while ((posY + step < bottomLeft)); do
    curl -X POST http://localhost:17000 -d "move 0 $step"
    posY=$((posY + step))
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
  done

  while ((posX + step < bottomRight)); do
    curl -X POST http://localhost:17000 -d "move $step 0"
    posX=$((posX + step))
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
  done

  while ((posY - step > topRight)); do
    curl -X POST http://localhost:17000 -d "move 0 $((-step))"
    posY=$((posY - step))
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
  done

  while ((posX - step > topLeft)); do
    curl -X POST http://localhost:17000 -d "move $((-step)) 0"
    echo $((-step))
    posX=$((posX - step))
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
  done
done