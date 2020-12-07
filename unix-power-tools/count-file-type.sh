#!/usr/bin/env bash

find ${*-.} -type f -print | xargs file |
awk '{
  $1=NULL
  t[$0]++
  
}
END {
  for (i in t) printf("%d\t%s\n", t[i], i); 
}' |
sort -nr

