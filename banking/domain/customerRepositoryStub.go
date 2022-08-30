package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Even", City: "New City", ZipCode: "241", Birth: "1990-01-01", Status: "1"},
		{Id: "1002", Name: "Ian", City: "New City", ZipCode: "241", Birth: "2000-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
