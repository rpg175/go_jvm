//反转控制IoC – Inversion of Control 是一种软件设计的方法，
//其主要的思想是把控制逻辑与业务逻辑分享，
//不要在业务逻辑里写控制逻辑，
//这样会让控制逻辑依赖于业务逻辑，
//而是反过来，让业务逻辑依赖控制逻辑。
//在《IoC/DIP其实是一种管理思想》中的那个开关和电灯的示例一样，
//开关是控制逻辑，
//电器是业务逻辑，
//不要在电器中实现开关，
//而是把开关抽象成一种协议，
//让电器都依赖之。
//这样的编程方式可以有效的降低程序复杂度，
//并提升代码重用。
package main

import "fmt"

// 有了这样的嵌入，就可以像UI组件一样的在结构构的设计上进行层层分解
type Widget struct {
	X, Y int
}

type Label struct {
	Widget        // Embedding (delegation)
	Text   string // Aggregation
}

type Button struct {
	Label // Embedding (delegation)
}

type ListBox struct {
	Widget          // Embedding (delegation)
	Texts  []string // Aggregation
	Index  int      // Aggregation
}

type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}

func (label Label) Paint() {
	fmt.Printf("%p:Label.Paint(%q)\n", &label, label.Text)
}

//因为这个接口可以通过 Label 的嵌入带到新的结构体，
//所以，可以在 Button 中可以重载这个接口方法以
// Override
func (button Button) Paint() {
	fmt.Printf("Button.Paint(%s)\n", button.Text)
}

func (button Button) Click() {
	fmt.Printf("Button.Click(%s)\n", button.Text)
}

func (listBox ListBox) Paint() {
	fmt.Printf("ListBox.Paint(%q)\n", listBox.Texts)
}

func (listBox ListBox) Click() {
	fmt.Printf("ListBox.Click(%q)\n", listBox.Texts)
}

func main() {
	label := Label{Widget{10, 10}, "State:"}
	label.X = 11
	label.Y = 12
	fmt.Printf("label %v", label)

	button1 := Button{Label{Widget{10, 70}, "OK"}}
	listBox := ListBox{Widget{10, 40}, []string{"AL", "AK", "AZ", "AR"}, 0}
	for _, painter := range []Painter{label, listBox, button1} {
		painter.Paint()
	}

	for _, widget := range []interface{}{label, listBox, button1} {
		widget.(Painter).Paint()
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
		fmt.Println()
	}
}
