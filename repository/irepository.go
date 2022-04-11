package repository

import (
	sqlclient "a2billing-go-api/internal/sqldb/mysql/driver"
)

var BillingSqlClient sqlclient.IMySqlConnector
