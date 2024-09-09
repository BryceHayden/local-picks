#!/bin/bash

# -e or --echo-errors (outputs sql commands sent to server to console)
# https://www.postgresql.org/docs/devel/app-psql.html#APP-PSQL-OPTION-ECHO-QUERIES

# -q or --quiet (reduces output to console from psql)
# https://www.postgresql.org/docs/devel/app-psql.html#APP-PSQL-OPTION-QUIET

# -f or --file=filename (read from file instead of standard input)
# https://www.postgresql.org/docs/devel/app-psql.html#APP-PSQL-OPTION-FILE
psql -U $POSTGRES_USER -d $POSTGRES_DB --echo-errors -q -f ~/migrations/00-init.sql

bin/bash ~/schema-script.sh
/bin/bash ~/seed-script.sh
