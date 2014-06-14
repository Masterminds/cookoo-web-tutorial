#!/bin/bash
DEST="table_of_contents.md"
BASE="https://github.com/Masterminds/cookoo-web-tutorial/tree"

echo "# Contents" > $DEST

list=`git branch --list '*_*'`
for i in $list; do
  echo $i
  echo "* [$i]($BASE/$i)" >> $DEST
done
