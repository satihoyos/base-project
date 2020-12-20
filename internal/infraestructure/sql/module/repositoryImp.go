package module

import (
	"errors"
	"sifiv/internal/core/domain"
	repo "sifiv/internal/core/ports/module"
	adapter "sifiv/internal/infraestructure/sql"
	"fmt"
)

//TODO manejar transaccionalidad
//TODO loggear errores
//TODO ver pull de conexion 
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

func (mr moduleRepoImp) GetByCode(code string) (module domain.Module) {
	
	q := `SELECT
			modulo_codigo, modulo_descripcion, modulo_estado
		 FROM modulo WHERE modulo_codigo = ($1)
		`
	db := mr.connection.Get()
	defer db.Close()
	rows, err := db.Query(q, code)

	if err != nil {		
		fmt.Println(err)
		return 
	}

	defer rows.Close()
	for rows.Next() {
		
		err = rows.Scan(
			&module.Code,
			&module.Description,
			&module.State,
		)

		if err != nil {
			fmt.Println(err)
			return
		}		
	}

	return module
}

func (mr moduleRepoImp) Save(module domain.Module) error {	

	q := `INSERT INTO 
			modulo (modulo_codigo, modulo_descripcion,modulo_estado) 
				VALUES ($1, $2, $3)`

	db := mr.connection.Get()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	r, err := stmt.Exec(module.Code, module.Description, module.State)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error: se esperaba 1 fila afectada")
	}

	return nil
}

func (mr moduleRepoImp) Edit(module domain.Module){
	q := `UPDATE 
			modulo 
			SET modulo_descripcion = ($1), modulo_estado = ($2)
			WHERE modulo_codigo = ($3)`

	db := mr.connection.Get()
	defer db.Close()

	_, err := db.Exec(q, module.Description, module.State,  module.Code)

	if err != nil {
		fmt.Println(err)
		return 
	}

}
