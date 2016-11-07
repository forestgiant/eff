#!/bin/bash

#This is intended to be run inside the examples folder
for var in "$@"
do
    echo "Creating a new example named: $var"
    mkdir $var
    cp boilerplate/main.go $var/main.go
    mkdir $var/drawable
    cp boilerplate/drawable/mydrawable.go $var/drawable/mydrawable.go
    echo "$var\n!$var/" >> ../.gitignore
    awk -v var=$var '/ifeq/{print;print "\tCGO_ENABLED=1 GOARCH=386 go build -o examples/" var "/" var " examples/" var "/main.go";next}1' ../Makefile > make.tmp
    awk -v var=$var '/else/{print;print "\tgo build -o examples/" var "/" var " examples/" var "/main.go";next}1' make.tmp > ../Makefile
    rm make.tmp
done