//Models/CustomerModel.go
package Models

type Customer struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Note    string `json:"note"`
}

func (b *Customer) TableName() string {
	return "Customer"
}
