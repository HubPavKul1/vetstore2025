package catalogs

import (
	"fmt"

	"fyne.io/fyne/v2/dialog"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
)



func VeiwProductCatalog() {
	w := uiUtils.CreateNewWindow("КАТАЛОГ ТОВАРОВ", false)


	products, err := services.GetProductsService()
	if err != nil {
		dialog.NewError(err, w)
	}

	for _, prod := range products {

		fmt.Println(prod.Name)

	}



	w.Show()
}