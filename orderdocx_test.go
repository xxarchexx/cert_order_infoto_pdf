package orderdocx

import (
	"fmt"
	"os"
	"testing"
)

func TestAbs(t *testing.T) {
	d := Data{}
	d.Amount = "5.55"
	d.CompanyName = "test"
	d.Count = "1"
	d.Quantity = "52"
	d.ServiceName = "ssssnnn"
	d.ProductName = "test"
	d.WorkDays = "34"
	ttt, err := os.Create("test.docx")
	if err != nil {
		fmt.Print(err)
	}
	d.Writer = ttt
	defer ttt.Close()
	d.Fill()
}
