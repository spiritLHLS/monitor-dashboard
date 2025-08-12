package response

import "server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
