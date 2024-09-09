#!/bin/bash

white="\033[0m"
green="\033[0;32m"
cyan="\033[0;36m"

for file in `ls "$(dirname $0)"/seeds/** | sort`; do
  echo -e "\n${green}Running ${cyan} $file db seed scripts${green} on project db...${white}"
    psql -U $POSTGRES_USER -d $POSTGRES_DB --echo-errors -q -f $file
  echo ""
done
