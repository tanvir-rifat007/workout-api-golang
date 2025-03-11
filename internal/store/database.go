package store

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)


func Open() (*sql.DB,error){
	dbUrl:= os.Getenv("DB_URL")
	if dbUrl == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable not set")
	}
	db,err:= sql.Open("pgx",dbUrl)

	if err!=nil{
		return nil,err
	}

	return db,nil

} 

func MigrateFS(db *sql.DB, migrateFS fs.FS,dir string)error{
	goose.SetBaseFS(migrateFS)

	defer func(){
		goose.SetBaseFS(nil)

	}()

	return Migrate(db,dir)
}


func Migrate(db *sql.DB,dir string)error{
	err:= goose.SetDialect("postgres")

	if err!=nil{
		return fmt.Errorf("failed to set dialect: %w",err)
	}

	err = goose.Up(db,dir)

	if err!=nil{
		return fmt.Errorf("failed to migrate db: %w",err)
	}

	return nil

}