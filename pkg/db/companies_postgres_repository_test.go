package db_test

import (
	"context"
	"testing"

	"github.com/OleksiiPyvovar/companies-crud/pkg/db"
	"github.com/OleksiiPyvovar/companies-crud/pkg/domain"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	connURL = "user=postgres password=postgres host=35.238.184.102 port=5432 dbname=companies"
)

type CompaniesRepositoryTestSuite struct {
	suite.Suite
	pool *pgxpool.Pool
	repo *db.CompaniesPostgresRepository
}

func TestCompaniesRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &CompaniesRepositoryTestSuite{})
}

func (s *CompaniesRepositoryTestSuite) SetupSuite() {
	pool, err := pgxpool.Connect(context.Background(), connURL)
	require.NoError(s.T(), err)

	s.pool = pool
	s.repo = db.NewCompaniesPostgresRepository(s.pool)
}

func (s *CompaniesRepositoryTestSuite) SetupTest() {
	_, err := s.pool.Exec(context.Background(), "DELETE FROM companies")
	require.NoError(s.T(), err)

	for _, company := range getTestCompanies() {
		q := "INSERT INTO companies (name, code, country, website, phone, id) " + "VALUES ($1,$2,$3,$4,$5,$6)"
		_, err = s.pool.Exec(context.Background(), q,
			company.Name, company.Code, company.Country,
			company.Website, company.Phone, company.ID)
		require.NoError(s.T(), err)
	}
}

func (s *CompaniesRepositoryTestSuite) TestCompaniesList() {
	expected := getTestCompanies()
	got, err := s.repo.List(domain.ListFilter{Limit: 10})
	require.NoError(s.T(), err)
	assert.ElementsMatch(s.T(), got, expected)

	got, err = s.repo.List(domain.ListFilter{Limit: 5})
	require.NoError(s.T(), err)
	assert.ElementsMatch(s.T(), got, expected[:5])

	filter := []domain.Filter{
		{
			Attr:     "name",
			Value:    "Test-Co-1",
			Operator: "eq",
		},
	}
	got, err = s.repo.List(domain.ListFilter{Filters: filter, Limit: 5})
	require.NoError(s.T(), err)
	assert.Equal(s.T(), got[0], expected[0])

	filter = []domain.Filter{
		{
			Attr:     "country",
			Value:    "test-land-2",
			Operator: "eq",
		},
		{
			Attr:     "code",
			Value:    "test-2",
			Operator: "eq",
		},
	}
	got, err = s.repo.List(domain.ListFilter{Filters: filter, Limit: 5})
	require.NoError(s.T(), err)
	assert.Equal(s.T(), got[0], expected[1])

	filter = []domain.Filter{
		{
			Attr:     "name",
			Value:    "Test-Co-1",
			Operator: "ne",
		},
	}
	got, err = s.repo.List(domain.ListFilter{Filters: filter, Limit: 10})
	require.NoError(s.T(), err)
	assert.ElementsMatch(s.T(), got, expected[1:])

	filter = []domain.Filter{
		{
			Attr:     "name",
			Value:    "Test-Co-1",
			Operator: "ne",
		},
		{
			Attr:     "phone",
			Value:    "1-1-1-1",
			Operator: "eq",
		},
	}
	got, err = s.repo.List(domain.ListFilter{Filters: filter, Limit: 10})
	require.NoError(s.T(), err)
	assert.ElementsMatch(s.T(), got, expected[1:])

	filter = []domain.Filter{
		{
			Attr:     "name",
			Value:    "Test-Co-1",
			Operator: "ne",
		},
		{
			Attr:     "website",
			Value:    "www.test-10.com",
			Operator: "ne",
		},
	}
	got, err = s.repo.List(domain.ListFilter{Filters: filter, Limit: 10})
	require.NoError(s.T(), err)
	assert.ElementsMatch(s.T(), got, expected[1:9])
}

func (s *CompaniesRepositoryTestSuite) TestCompanyGetByID() {
	expected := getTestCompanies()

	got, err := s.repo.GetByID(1)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), got, expected[0])

	got, err = s.repo.GetByID(10)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), got, expected[9])
}

func (s *CompaniesRepositoryTestSuite) TestCompanyUpdate() {
	c := &domain.Company{
		ID:      2,
		Name:    "Test-Co-1-updated",
		Code:    "test-1-updated",
		Country: "test-land-1-updated",
		Website: "www.test-1.updated.com",
		Phone:   "1-1-1-1",
	}

	err := s.repo.Update(c)
	require.NoError(s.T(), err)

	gotAfterUpdate, err := s.repo.GetByID(2)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), gotAfterUpdate, *c)
}

func (s *CompaniesRepositoryTestSuite) TestCompanyCreate() {
	c := &domain.Company{
		Name:    "Test-Co-11-created",
		Code:    "test-11-created",
		Country: "test-land-11-created",
		Website: "www.test-1.created.com",
		Phone:   "1-1-1-1",
	}

	err := s.repo.Create(c)
	require.NoError(s.T(), err)

	gotAfterCreate, err := s.repo.GetByID(c.ID)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), gotAfterCreate, *c)
}

func (s *CompaniesRepositoryTestSuite) TestCompanyDelete() {
	deleted := s.repo.DeleteByID(1)
	assert.Equal(s.T(), deleted, true)

	_, err := s.repo.GetByID(1)
	assert.Contains(s.T(), err.Error(), "not found")
}

func getTestCompanies() []domain.Company {
	return []domain.Company{
		{
			ID:      1,
			Name:    "Test-Co-1",
			Code:    "test-1",
			Country: "test-land-1",
			Website: "www.test-1.com",
			Phone:   "1-1-1-1",
		},
		{
			ID:      2,
			Name:    "Test-Co-2",
			Code:    "test-2",
			Country: "test-land-2",
			Website: "www.test-2.com",
			Phone:   "1-1-1-1",
		},
		{
			ID:      3,
			Name:    "Test-Co-3",
			Code:    "test-3",
			Country: "test-land-3",
			Website: "www.test-3.com",
			Phone:   "1-1-1-1",
		},
		{
			ID:      4,
			Name:    "Test-Co-4",
			Code:    "test-4",
			Country: "test-land-1",
			Website: "www.test-4.com",
			Phone:   "1-1-1-1",
		},
		{
			ID:      5,
			Name:    "Test-Co-5",
			Code:    "test-5",
			Country: "test-land-2",
			Website: "www.test-5.com",
			Phone:   "1-1-1-1",
		},
		{
			ID:      6,
			Name:    "Test-Co-6",
			Code:    "test-6",
			Country: "test-land-1",
			Website: "www.test-6.com",
			Phone:   "1-1-1-1",
		},
		{
			ID:      7,
			Name:    "Test-Co-7",
			Code:    "test-7",
			Country: "test-land-7",
			Website: "www.test-7.com",
			Phone:   "1-1-1-1",
		},
		{
			ID:      8,
			Name:    "Test-Co-8",
			Code:    "test-8",
			Country: "test-land-8",
			Website: "www.test-8.com",
			Phone:   "1-1-1-1",
		},
		{
			ID:      9,
			Name:    "Test-Co-9",
			Code:    "test-9",
			Country: "test-land-9",
			Website: "www.test-9.com",
			Phone:   "1-1-1-1",
		},
		{
			ID:      10,
			Name:    "Test-Co-10",
			Code:    "test-10",
			Country: "test-land-10",
			Website: "www.test-10.com",
			Phone:   "1-1-1-1",
		},
	}
}
