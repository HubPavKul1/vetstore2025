package catalogs

import (
	"fmt"

	"fyne.io/fyne/v2/dialog"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)



func VeiwProductCatalog() {
	w := ui_utils.CreateNewWindow("КАТАЛОГ ТОВАРОВ", false)


	products, err := services.GetProductsService()
	if err != nil {
		dialog.NewError(err, w)
	}

	for _, prod := range products {

		fmt.Println(prod.Name)

	}



	w.Show()
}