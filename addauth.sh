#!/bin/sh

set -eu

if [ -z "$1" ]
then
	echo "usage: $1 [file]" >&2
	exit 1
fi

printf "Username: "
read -r USER || ( echo && exit 1 )

if grep -qs "^$USER" "$1"
then
    printf "\"%s\" already specified in %s\\n" "$USER" "$1"
	printf "Do you want to replace it? [yes] "
    read -r ANS || ( echo && exit 1 )
    if [ "$ANS" = "yes" ] || [ -z "$ANS" ]
	then
		sed -i "/$USER/d" "$1"
    else
		echo "Quitting."
		exit 0
    fi
fi

stty -echo
printf "Password: "
read -r PASS || ( stty echo && echo && exit 1 )
stty echo
echo

HASH="$(printf "%s%s" "$USER" "$PASS" | sha256sum | cut -f1 -d' ')"
printf "%s\\t%s\\n" "$USER" "$HASH" >> "$1"

