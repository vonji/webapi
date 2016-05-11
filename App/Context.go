package vonji

import "github.com/jinzhu/gorm"

type Context struct {
	App *App
	Db  *gorm.DB
}

var context = Context{}

func Init(app *App, db *gorm.DB) {
	context.App = app
	context.Db = db
}

func GetContext() *Context {
	return &context
}
