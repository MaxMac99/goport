#!/bin/sh
count=1
while test $count -lt 1000
do
	echo "Test $count"
	count=$((count+1))
	sleep 1
done
