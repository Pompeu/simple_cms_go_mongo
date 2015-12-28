#!/bin/bash
while true; do
  find `pwd` | grep -E "\.go$" | entr -d go run $1 
done
