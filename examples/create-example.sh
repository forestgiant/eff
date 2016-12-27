#!/bin/bash

#This is intended to be run inside the examples folder
for var in "$@"
do
    echo "Creating a new example named: $var"
    mkdir $var
    cp boilerplate/main.go $var/main.go
    echo "$var\n!$var/" >> ../.gitignore
    awk -v var=$var '/build:/{print;print "\tcd examples/" var "; go build; cd ../../";next}1' ../Makefile > make.tmp
    mv make.tmp ../Makefile
done