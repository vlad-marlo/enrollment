package migrator_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vlad-marlo/enrollment/internal/pkg/pgx/client"
	"github.com/vlad-marlo/enrollment/internal/pkg/pgx/migrator"
	"testing"
)

func TestMigrate_Positive(t *testing.T) {
	cli, td := client.NewTest(t)
	defer td()
	i, err := migrator.MigrateUp(cli)
	assert.NoError(t, err)
	if assert.NotEmpty(t, i) {
		assert.Equal(t, migrator.Migrations, i)
	}
}

func TestMigrate_Negative(t *testing.T) {
	cli := client.BadCli(t)
	i, err := migrator.MigrateUp(cli)
	assert.Error(t, err)
	assert.Empty(t, i)
}

func TestMigrateDown_Negative(t *testing.T) {
	cli := client.BadCli(t)
	i, err := migrator.MigrateDown(cli)
	assert.Error(t, err)
	assert.Empty(t, i)
}
