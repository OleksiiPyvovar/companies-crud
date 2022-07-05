package db

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/OleksiiPyvovar/companies-crud/pkg/domain"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	OperatorEqual    = "="
	OperatorNotEqual = "!="
)

type CompaniesPostgresRepository struct {
	conn *pgxpool.Pool
}

func NewCompaniesPostgresRepository(conn *pgxpool.Pool) *CompaniesPostgresRepository {
	return &CompaniesPostgresRepository{
		conn: conn,
	}
}

func (cpr *CompaniesPostgresRepository) Create(company *domain.Company) error {
	query := "INSERT INTO companies (name, code, country, website, phone) " +
		"VALUES ($1,$2,$3,$4,$5) RETURNING id"

	err := cpr.conn.QueryRow(context.Background(), query,
		company.Name, company.Code, company.Country,
		company.Website, company.Phone).Scan(&company.ID)
	if err != nil {
		return fmt.Errorf("query exec: %w", err)
	}

	return nil
}

func (cpr *CompaniesPostgresRepository) Update(company *domain.Company) error {
	query := "UPDATE companies SET name = $1, code = $2, country = $3, website = $4, phone = $5 " +
		"WHERE id = $6"

	_, err := cpr.conn.Exec(context.Background(), query,
		company.Name, company.Code, company.Country,
		company.Website, company.Phone, company.ID)
	if err != nil {
		return fmt.Errorf("query exec: %w", err)
	}

	return nil
}

func (cpr *CompaniesPostgresRepository) DeleteByID(id int) bool {
	var result string
	query := "DELETE FROM companies WHERE id=$1"

	err := cpr.conn.QueryRow(context.Background(), query, id).Scan(&result)

	return err.Error() == "no rows in result set"
}

func (cpr *CompaniesPostgresRepository) GetByID(id int) (domain.Company, error) {
	query := "SELECT name, code, country, website, phone FROM companies WHERE id = $1"

	row := domain.Company{ID: id}
	err := cpr.conn.QueryRow(context.Background(), query, id).Scan(
		&row.Name, &row.Code, &row.Country, &row.Website, &row.Phone,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return row, errors.New("not found")
	}

	if err != nil {
		return row, fmt.Errorf("query row: %w", err)
	}

	return row, nil
}

func (cpr *CompaniesPostgresRepository) List(options domain.ListFilter) ([]domain.Company, error) {
	query := "SELECT id, name, code, country, website, phone FROM companies %s"

	filter, values := buildAttributeParams(options)
	query = fmt.Sprintf(query, filter)

	rows, err := cpr.conn.Query(context.Background(), query, values...)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	var res []domain.Company
	for rows.Next() {
		var row domain.Company
		err := rows.Scan(&row.ID, &row.Name, &row.Code, &row.Country, &row.Website, &row.Phone)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		res = append(res, row)
	}

	return res, nil
}

func buildAttributeParams(options domain.ListFilter) (string, []interface{}) {
	var (
		values     []interface{}
		params     []string
		q          string
		filterTmpl = "%s %s $%d"
	)

	for i, filter := range options.Filters {
		params = append(params,
			fmt.Sprintf(filterTmpl, filter.Attr, convertOperator(filter.Operator), i+1),
		)
		values = append(values, filter.Value)
	}

	if len(params) != 0 {
		q = fmt.Sprintf("WHERE %s ORDER BY id ASC LIMIT $%d", strings.Join(params, " AND "), len(params)+1)
	} else {
		q = "ORDER BY id ASC LIMIT $1"
	}
	values = append(values, options.Limit)

	return q, values
}

func convertOperator(operator string) string {
	switch operator {
	case "eq":
		return OperatorEqual
	case "ne":
		return OperatorNotEqual
	}

	return OperatorEqual
}
