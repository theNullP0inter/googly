package command

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(migrateCmd)
}

func getMigration() *migrate.Migrate {
	db, err := sql.Open("mysql", viper.GetString("RDB_URL"))
	if err != nil {
		panic(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}
	return m
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migration Lib",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(("Running Migrations"))
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Ups all the migrations",
	Run: func(cmd *cobra.Command, args []string) {
		m := getMigration()
		m.Steps(2)
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Downs all the migrations",
	Run: func(cmd *cobra.Command, args []string) {
		m := getMigration()
		m.Steps(-2)
	},
}
