#!/bin/bash
if $1 == ''
then
	echo 'Please give the cert name as an argument!!'
	exit 1
fi

certname=$1

openssl req -newkey rsa:4096 \
            -x509 \
            -sha256 \
            -days 3650 \
            -nodes \
            -out $certname.crt \
            -keyout $certname.key
