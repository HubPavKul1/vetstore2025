package addProductForm

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/dialogs"
)


func SaveNewProduct( w fyne.Window, f *AddProductFormData) {
    newProduct := models.Product{
        SubCategoryID: f.SubcategoryID, 
        PackagingID: f.PackagingID,
        UnitID: f.UnitID,
    }
    newProduct.Name =f.ProductName
    // Сохраняем товар в базе данных
    _, err := services.CreateProductService(newProduct)
        if err != nil {
            dialog.NewError(err, w).Show()
            return
        }
    dialogs.SuccessAddDataDialog(w).Show()
}