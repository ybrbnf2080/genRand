#!/bin/bash
set -e

if [[ -z "${SLEEP_TIME}" ]]; then
	SLEEP_TIME=0.05
fi

if [[ -z "${COLOR_SHIFT}" ]]; then
	COLOR_SHIFT=-1000
fi

BASEDIR=/tmp/genRand
mkdir -p $BASEDIR
PATH_TO_ANIM=$1

BASE=$(basename "$1")
filename=$(printf '%b' ${BASE//%/\\x})
FILE_PATH=$PATH_TO_ANIM

if [[ "$1" == *"http"* ]]; then
    wget $1 -P $BASEDIR -nc
fi


CHASHEDIR=$BASEDIR/cache/$filename
echo $CHASHEDIR
mkdir -p $BASEDIR/cache/$filename

if ! ls $CHASHEDIR/img* >/dev/null 2>&1; then
	echo CHASHEDIR
	if [[ "$1" != *"webp"* ]]; then
		ffmpeg -i $FILE_PATH  $BASEDIR/cache/$filename/img%03d.jpg
	fi

	if [[ "$1" == *"webp"* ]]; then
		magick  $FILE_PATH "$BASEDIR/cache/$filename/img.jpg"
	fi 
fi
PATH_TO_ANIM="$BASEDIR/cache/$filename"
while :
do
for o in $PATH_TO_ANIM/*.jpg; do gen-rand -color $COLOR_SHIFT -s "$o" && sleep $SLEEP_TIME ; done
done

