# Driver

在 DB 連線前必須先指定 DB Driver

## MSSQL

```go

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

```

[driver](https://github.com/denisenkom/go-mssqldb)

## MySQL

```go

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

```

[driver](https://github.com/go-sql-driver/mysql)

---- 

[DB Driver](https://github.com/golang/go/wiki/SQLDrivers)