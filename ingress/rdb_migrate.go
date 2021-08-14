package ingress

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

func getMigration(path string, driver_name string, driver database.Driver) *migrate.Migrate {

	m, err := migrate.NewWithDatabaseInstance(
		path,
		driver_name,
		driver,
	)
	if err != nil {
		panic(err)
	}
	return m
}

func NewMigrateCommand(config *CommandConfig, path string, driver_name string, driver database.Driver) *cobra.Command {

	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migration on a database driver",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(("Running Migrations"))
		},
	}

	var migrateUpCmd = &cobra.Command{
		Use:   "up",
		Short: "Ups all the migrations",
		Run: func(cmd *cobra.Command, args []string) {
			m := getMigration(path, driver_name, driver)
			if err := m.Up(); err != nil && err != migrate.ErrNoChange {
				panic(err)
			}
		},
	}

	var migrateDownCmd = &cobra.Command{
		Use:   "down",
		Short: "Downs all the migrations",
		Run: func(cmd *cobra.Command, args []string) {
			m := getMigration(path, driver_name, driver)
			if err := m.Down(); err != nil && err != migrate.ErrNoChange {
				panic(err)
			}
		},
	}

	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)

	return migrateCmd

}

type MigrationIngress struct {
	*Ingress
}

func NewMigrationIngress(name string, path string, driver database.Driver) *MigrationIngress {

	cmd := NewMigrateCommand(
		&CommandConfig{
			Name:  name,
			Short: "DB Migrator",
		},
		fmt.Sprintf("file://%s", path),
		"mysql", driver,
	)
	ingress := NewIngress(cmd)
	return &MigrationIngress{ingress}

}
