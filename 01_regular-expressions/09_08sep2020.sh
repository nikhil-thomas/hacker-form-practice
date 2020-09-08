#!/usr/bin/env bash

source ../lib/util.sh

printline

data1=$(mktemp)
cat <<EOF>>$data1
Hello, my name is Ben. Please visit
my website at http://www.forta.com/.
EOF

echo data1
cat $data1
printline

echo 'word: Ben'
grep 'Ben' $data1
printline

echo 'word: my'
grep 'my' $data1
printline

data2=$(mktemp)
cat <<EOF>> $data2
sales.xls
sales1.xls
orders3.xls
sales2.xls
sales3.xls
apac1.xls
europe2.xls
na1.xls
na2.xls
sa1.xls
EOF

cat $data2
printline

echo 'match sales+any one character'
grep 'sales.' $data2
printline

echo 'match <any>a<any>dot'
grep '.a.\.' $data2
printline

