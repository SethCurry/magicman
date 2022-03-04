package gui

import (
	"context"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/SethCurry/magicman/pkg/ent"
)

func Run() error {
	a := app.New()
	win := a.NewWindow("MagicMan")
	win.SetContent(widget.NewLabel("magicman"))
	win.Show()
	a.Run()
	return nil
}

func NewApplication(ctx context.Context, db *ent.Client) *Application {
	a := app.New()
	return &Application{
		db:     db,
		app:    a,
		ctx:    ctx,
		window: a.NewWindow("MagicMan"),
	}

}

type Application struct {
	db     *ent.Client
	app    fyne.App
	window fyne.Window
	ctx    context.Context
}

func (a *Application) Run() {
	tbl, err := CardTable(a.ctx, a.db)
	if err != nil {
		panic(err)
	}
	content := container.NewBorder(DefaultToolbar(), nil, nil, nil, tbl)
	a.window.SetContent(content)
	a.window.Show()
	a.app.Run()
}

func DefaultToolbar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
	)
}

func CardTable(ctx context.Context, db *ent.Client) (*widget.Table, error) {
	cards, err := db.Card.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	list := widget.NewTable(
		func() (int, int) {
			return len(cards), 2
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			var text string
			switch i.Col {
			case 0:
				text = cards[i.Row].MultiverseID
			case 1:
				text = cards[i.Row].Name
			}
			o.(*widget.Label).SetText(text)
		})

	return list, nil
}
