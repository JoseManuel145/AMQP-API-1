package infraestructure

import (
	"fmt"
	"log"
	"report/src/core"
	"report/src/report/domain/entities"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

// Create inserta un reporte en la base de datos.
func (mysql *MySQL) Create(id int, title, content string) error {
	query := "INSERT INTO reports (id, title, content) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, id, title, content)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

// FindAll obtiene todos los reportes de la base de datos.
func (mysql *MySQL) ViewAll() ([]entities.Report, error) {
	query := "SELECT id, title, content, status FROM reports"
	rows := mysql.conn.FetchRows(query)
	if rows == nil {
		return nil, fmt.Errorf("error al recuperar los registros")
	}
	defer rows.Close()

	var reports []entities.Report
	for rows.Next() {
		var report entities.Report
		if err := rows.Scan(&report.ID, &report.Title, &report.Content, &report.Status); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
		reports = append(reports, report)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}
	return reports, nil
}

// FindByID obtiene un reporte por su ID.
func (mysql *MySQL) ViewOne(id int) (*entities.Report, error) {
	query := "SELECT id, title, content FROM reports WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	if rows == nil {
		return nil, fmt.Errorf("error al recuperar el reporte")
	}
	defer rows.Close()

	var report entities.Report
	if rows.Next() {
		if err := rows.Scan(&report.ID, &report.Title, &report.Content); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
	} else {
		return nil, fmt.Errorf("no se encontró ningún reporte con el ID %d", id)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}
	return &report, nil
}
