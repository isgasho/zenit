package mysql

import (
  "gitlab.com/swapbyt3s/zenit/common"
  "gitlab.com/swapbyt3s/zenit/config"
  "gitlab.com/swapbyt3s/zenit/output"
)

type Table struct {
  schema    string
  table     string
  size      float64
  rows      float64
  increment float64
}

const QUERY_SQL_TABLES = `
SELECT table_schema AS 'schema',
       table_name AS 'table',
       data_length + index_length AS 'size',
       table_rows AS 'rows',
       auto_increment AS 'increment'
FROM information_schema.tables
WHERE table_schema NOT IN ('mysql','sys','performance_schema','information_schema','percona')
ORDER BY table_schema, table_name;
`

func GatherTables() {
  conn, err := common.MySQLConnect(config.DSN_MYSQL)
  defer conn.Close()
  if err != nil {
    panic(err)
  }

  rows, err := conn.Query(QUERY_SQL_TABLES)
  defer rows.Close()
  if err != nil {
    panic(err)
  }

  var a = output.LoadAccumulator()

  for rows.Next() {
    var t Table

    rows.Scan(
      &t.schema,
      &t.table,
      &t.size,
      &t.rows,
      &t.increment)

    a.AddItem(output.Metric{
      Key:   "mysql_stats_tables",
      Tags:  []output.Tag{output.Tag{"schema", t.schema},
                          output.Tag{"table", t.table},
                          output.Tag{"type", "size"}},
      Value: t.size,
    })
    a.AddItem(output.Metric{
      Key:   "mysql_stats_tables",
      Tags:  []output.Tag{output.Tag{"schema", t.schema},
                          output.Tag{"table", t.table},
                          output.Tag{"type", "rows"}},
      Value: t.rows,
    })
    a.AddItem(output.Metric{
      Key:   "mysql_stats_tables",
      Tags:  []output.Tag{output.Tag{"schema", t.schema},
                          output.Tag{"table", t.table},
                          output.Tag{"type", "increment"}},
      Value: t.increment,
    })
  }
}
