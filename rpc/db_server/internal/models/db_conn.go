package models

import (
	"rpc/db_server/internal/config"
	"rpc/db_server/internal/sql"
)

var sqlConfig = config.New()
var conn = sql.Connect(sqlConfig)
