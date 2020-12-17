package module

import (
	"sifiv/internal/core/domain"
	repo "sifiv/internal/core/ports/module"
	adapter "sifiv/internal/infraestructure/sql"
)

type moduleRepoImp struct {
	connection adapter.Connection
}

func NewModuleRepository(cn adapter.Connection) repo.ModuleRepository {
	return &moduleRepoImp{connection: cn}
}

func (mr moduleRepoImp) Get() (modules []domain.Module) {

	db := mr.connection.Get()
	defer db.Close()

	q := `SELECT
			modulo_codigo, modulo_descripcion, modulo_estado
		 FROM modulo
		`
	rows, err := db.Query(q)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		e := domain.Module{}
		err = rows.Scan(
			&e.Code,
			&e.Description,
			&e.State,
		)

		if err != nil {
			return
		}

		modules = append(modules, e)
	}

	return modules
}
