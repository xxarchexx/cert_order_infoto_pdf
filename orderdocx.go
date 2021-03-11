package orderdocx

import (
	"io"

	"github.com/nguyenthenguyen/docx"
)

type Data struct {
	io.Writer
	ServiceName string
	ProductName string
	WorkDays    string
	Count       string
	Amount      string
	Quantity    string
	Price       string
	CompanyName string
}

var templateName = "./document.docx"

func (data *Data) Fill() error {
	// err := copy()
	// if err != nil {
	// 	return err
	// }
	r, err := docx.ReadDocxFile(templateName)
	if err != nil {
		return err
	}

	docx1 := r.Editable()
	docx1.Replace("$$serv-name", data.ServiceName, -1)
	docx1.Replace("$$product-name", data.ProductName, -1)
	docx1.Replace("$$work-days", data.WorkDays, -1)
	docx1.Replace("$$cont", data.Count, -1)
	docx1.Replace("$$amount", data.Amount, -1)
	docx1.Replace("$$quantity", data.Quantity, -1)
	docx1.Replace("$$count", data.Count, -1)
	docx1.Replace("$$companyname", data.CompanyName, -1)

	docx1.Write(data.Writer)
	r.Close()
	return nil
}
