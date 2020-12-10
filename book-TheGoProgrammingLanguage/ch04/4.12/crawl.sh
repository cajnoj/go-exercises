#!/bin/bash
set -e
for i in {1..2396}; do
	echo $i
	curl https://xkcd.com/$i//info.0.json > data/$i
	sleep 0.01
done
