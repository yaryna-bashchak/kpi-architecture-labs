#!/bin/bash+

width=10

curl -X POST http://localhost:17000 -d "green"
curl -X POST http://localhost:17000 -d "bgrect $width $width $((400-width)) $((400-width))"
curl -X POST http://localhost:17000 -d "update"