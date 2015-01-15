#!/bin/bash

vercomp () {
    if [[ $1 == $2 ]]
    then
        return 0
    fi
    local IFS=.
    local i ver1=($1) ver2=($2)
    # fill empty fields in ver1 with zeros
    for ((i=${#ver1[@]}; i<${#ver2[@]}; i++))
    do
        ver1[i]=0
    done
    for ((i=0; i<${#ver1[@]}; i++))
    do
        if [[ -z ${ver2[i]} ]]
        then
            # fill empty fields in ver2 with zeros
            ver2[i]=0
        fi
        if ((10#${ver1[i]} > 10#${ver2[i]}))
        then
            return 1
        fi
        if ((10#${ver1[i]} < 10#${ver2[i]}))
        then
            return 2
        fi
    done
    return 0
}

GOVERSION=`go version|awk '{print $3}'| sed 's/go//'`
vercomp $GOVERSION '1.4'
if [ $? == 2 ]; then
	echo "You need to run this with go version >= 1.4"
	exit 1
fi

go get -v github.com/mattn/go-sqlite3
go install github.com/mattn/go-sqlite3

go get -v github.com/gchaincl/dotsql
go get -v github.com/jinzhu/gorm
go get -v github.com/jmoiron/sqlx
go get -v github.com/lann/squirrel
go get -v github.com/astaxie/beego

go test -bench . -benchmem