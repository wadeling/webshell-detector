# webshell-detector

## usage

- first: make -C php
	- build ok in ubuntu 18.04
	- build failed in Macos:  may change php version to solve this which i am not try yet.

- second: go build -o test
- run: ./test

## comment
file "sample.php" has webshell vulnerability, after run test,we can see the score 

## notice
not support concurrent running
