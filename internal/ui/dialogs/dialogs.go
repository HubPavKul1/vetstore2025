package dialogs

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/HubPavKul1/vetstore2025/internal/ui/app"
)


func SuccessAddDataDialog(parent fyne.Window) dialog.Dialog {
	return dialog.NewInformation("", "ЗАПИСЬ УСПЕШНО ДОБАВЛЕНА В БАЗУ!", parent)
}

func ErrorDialog(parent fyne.Window, e error) dialog.Dialog {
	return dialog.NewError(e, parent)
}


func CreateAddDataDialog(
	parent fyne.Window, 
	title string, 
	) fyne.Window {

	w := app.MyApp.NewWindow(strings.ToUpper(title))
	w.Resize(fyne.NewSize(400, 300))
	w.CenterOnScreen()

	return w


		
	}



