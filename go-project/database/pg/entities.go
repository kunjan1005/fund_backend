package database

// Define the column struct to hold the column name, type, and constraints.
type Column struct {
	Name        string
	Type        string
	Constraints []string // Constraints like "NOT NULL", "UNIQUE", etc.
}

// Define the CreateTable struct with a table name and columns.
type CreateTable struct {
	TableName string
	Columns   []Column
}

// Define an interface for storage that includes the CreateTable method.
type Storage interface {
	CreateTableSQL(createTable CreateTable) (string, error)
}

var UserTable = CreateTable{
	TableName: "user",
	Columns: []Column{
		{Name: "id", Type: "SERIAL", Constraints: []string{"PRIMARY KEY"}},
		{Name: "name", Type: "VARCHAR(100)", Constraints: []string{"NOT NULL"}},
		{Name: "email", Type: "VARCHAR(100)", Constraints: []string{"NOT NULL", "UNIQUE"}},
		{Name: "created_at", Type: "TIMESTAMP", Constraints: []string{"DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: "TIMESTAMP", Constraints: []string{"DEFAULT CURRENT_TIMESTAMP"}},
	},
}
