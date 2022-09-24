package bridge

import "fmt"

type App interface {
	DoDesign()
	DoCode()
	DoTest()
}

type ChatApp struct{}

func (c *ChatApp) DoDesign() {
	fmt.Println("do Chat app Design ...")
}

func (c *ChatApp) DoCode() {
	fmt.Println("do Chat app DoCode ...")
}

func (c *ChatApp) DoTest() {
	fmt.Println("do Chat app DoTest ...")
}

type ICorp interface {
	developApp()
}

type AppCorp struct {
	App App
}

func (c *AppCorp) developApp() {
	c.App.DoDesign()
	c.App.DoCode()
	c.App.DoTest()
}
