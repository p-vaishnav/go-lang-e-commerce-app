package commands

import (
	"errors"
	"fmt"
	"migrations/configs"
	database "migrations/databse"

	"github.com/spf13/cobra"
)

func DropTables() *cobra.Command {
	return &cobra.Command{
		Use: "droptables",
		RunE: func(cmd *cobra.Command, args []string) error {
			if configs.App.Env != "local" {
				fmt.Println("Warning: Environment is not local. Tables won't be dropped")
				return nil
			}
			fmt.Println("App env is lcoal")

			dbConnection, sqlConnection := database.Connection()
			defer sqlConnection.Close()

			var tableNames []string
			if err := dbConnection.Table("information_schema.tables").
				Where("table_schema = ?", "public").Pluck("table_name", &tableNames).Error; err != nil {
				panic(err)
			}

			if len(tableNames) > 0 {
				for index, tableName := range tableNames {
					if err := dbConnection.Migrator().DropTable(tableName); err != nil {
						return errors.New("Error: while dropping tables:" + tableName)
					}
					fmt.Println("[ ", index, " ]", "dropped table: ", tableName)
				}
			}
			fmt.Println("Dropped all tables successfully")
			return nil
		},
	}
}
