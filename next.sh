#!/bin/bash

latest=$(ls -d -r day-* | grep -o '[0-9]\+' | head -1)
next=$(printf "%02d\n" $((${latest#0} + 1)))
cp -R templates day-$next
mv day-$next/day.go.tmpl day-$next/day-$next.go
mv day-$next/day_test.go.tmpl day-$next/day-$next\_test.go
echo day-$next
