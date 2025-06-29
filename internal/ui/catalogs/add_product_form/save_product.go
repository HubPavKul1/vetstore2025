package add_product_form

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/dialogs"
)

type AddProductForm struct {
    SubcategoryID uint
    Name string
    PackID uint
    UnitID uint
}

func SaveNewProduct( w fyne.Window, f *AddProductForm) {
    newProduct := models.Product{
        SubCategoryID: f.SubcategoryID, 
        PackagingID: f.PackID,
        UnitID: f.UnitID,
    }
    newProduct.Name =f.Name
    // Сохраняем товар в базе данных
    _, err := services.CreateProductService(newProduct)
        if err != nil {
            dialog.NewError(err, w).Show()
            return
        }
    dialogs.SuccessAddDataDialog(w).Show()
}