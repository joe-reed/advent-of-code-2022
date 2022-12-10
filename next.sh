#!/bin/bash

latest=$(ls -d -r day-* | grep -o '[0-9]\+' | head -1)
next=$(printf "%02d\n" $((${latest#0} + 1)))
cp -R templates day-$next
mv day-$next/day.go day-$next/day-$next.go
mv day-$next/day_test.go day-$next/day-$next\_test.go
