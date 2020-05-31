# Setup for macOS

- installation
  - `brew install postgresql`
- initialize database
  - `mkdir <database name>`
  - `initdb -D <database name> -U postgres -E utf8 -W`
- start database
  - `pg_ctl start -D <database name>`
- create user&database	
  - `createuser -P -d <user(role)name> -U postgres`
	- `-P`: require password
	- `-d`: authorization for database creation
  - `createdb <database name> -U <username> -`
- setup database
  - `psql -U <username> -f setup.sql -d <database name>`
- check database status
  - `pg_ctl status -D <database name>`
- check users
  - `psql <database name> -U <username> -`
  - `gwp=> \du`
- stop database
  - `pg_ctl stop -D <database name>`
  
