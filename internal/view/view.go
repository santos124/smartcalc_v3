package view

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"smartcalc_v3_wadina/internal/viewmodel"
	"strconv"
	"strings"
	"time"
)

/*
история вычислений
Программа должна сохранять историю операций, разрешать загрузку выражений из истории и очищать всю историю
История должна сохраняться между запусками приложения
В программу могут быть введены как целые, так и действительные числа, записанные либо через точку, либо в экспоненциальной форме.
Расчет следует производить после полного ввода вычисляемого выражения и нажатия символа=
справочный раздел произвольной формы
ввод действительных чисел
*/
/*
Программа должна быть разработана на языке Go 1.19.2Код программы должен быть расположен в папке src
При написании кода необходимо придерживаться стиля Google Code StyleВам необходимо разработать настольное приложение
Подготовить инсталлятор, который установит приложение в систему со стандартными настройками (путь установки, создание ярлыка)Подготовить реализацию с графическим интерфейсом пользователя для Mac OS, основанную на любой библиотеке или фреймворке GUI (допустима реализация GUI слоя в HTML/CSS/JS)
Программа должна быть реализована с использованием паттерна MVVM или MVP, ив коде представления не должно быть кода бизнес-логики
в модели, презентаторе и модели представления не должно быть кода интерфейса"ядро" калькулятора в виде алгоритма формирования и вычисления польской нотации и различных вычислительных функций подключить в виде динамической библиотеки на C/C++ из проектов SmartCalc v1.0 или SmartCalc v2.0
Модель должна представлять собой "ядро" с оберткой на языке GoМодель должна иметь всю функциональность калькулятора, чтобы в будущем ее можно было использовать без других слоев
Подготовить полное покрытие методов в слое модели модульными тестамиПриложение должно иметь раздел помощи с описанием интерфейса программы в произвольной форме
Программа должна сохранять историю операций, позволять загружать выражения из истории и очищать всю историюИстория должна сохраняться между запусками приложения
В программу можно вводить как целые, так и вещественные числа, записанные либо через точку, либо в экспоненциальной формеВычисление должно выполняться после полного ввода вычисляемого выражения и нажатия символа =
Вычисление произвольных арифметических выражений со скобками в инфиксной нотацииВычисление произвольных арифметических выражений со скобками в инфиксной системе с заменой переменной x на число
Построение графика функции, определенной с помощью выражения в инфиксной нотации с переменной x (с координатными осями, масштабным маркером и сеткой с адаптивным шагом).Нет необходимости предоставлять пользователю возможность изменять масштаб
Область определения и область значений функций, как минимум, ограничены числами от -1000000 до 1000000Для построения графика функции необходимо дополнительно указать отображаемую область определения и область значения
Проверенная точность дробной части составляет не менее 7 знаков после запятойПользователь должен иметь возможность вводить до 255 символов
Арифметические выражения в скобках в инфиксной нотации должны поддерживать следующие арифметические операции и математические функции:

*/
type View struct {
	app             fyne.App
	window          fyne.Window
	timer           *widget.Label
	infoEntry       *widget.Label
	expressionEntry *widget.Entry
	operatorEntry   *widget.Entry
	resultLabel     *widget.Label
	buttonResult    *widget.Button
	vm              *viewmodel.ViewModel
}

/*1 поле ввода строки с вычисляемым выражением (должно закончится щнаком равно "=")
2 поле результата текста3 график в случае задания функции
4 справка5 история

*/func NewView(vm *viewmodel.ViewModel) *View {
	return &View{app: nil,
		window: nil, timer: &widget.Label{},
		infoEntry: &widget.Label{}, expressionEntry: &widget.Entry{},
		operatorEntry: &widget.Entry{}, resultLabel: &widget.Label{},
		buttonResult: &widget.Button{}, vm: vm,
	}
}

func (v *View) createBaseUI() {
	a := app.New()
	w := a.NewWindow("Calculator")
	w.Resize(fyne.NewSize(800, 600))
	v.app = a
	v.window = w
}

func (v *View) createEntry() {
	infoEntry := widget.NewLabel("Введи выражение:")
	entry1 := widget.NewEntry()
	v.expressionEntry = entry1
	v.infoEntry = infoEntry
}

func (v *View) createResult() {
	result := widget.NewLabel("Тут будет результат")
	v.resultLabel = result
}

func (v *View) createButton() {
	btn := widget.NewButton("посчитать", func() {
		num1, err1 := strconv.ParseFloat(v.expressionEntry.Text, 64)
		num2 := 2.123
		if err1 != nil {
			v.resultLabel.SetText("ERROR")
			return
		}
		v.resultLabel.SetText(fmt.Sprintf(
			"сумма: %v\nразность: %v\nпроизведение: %v\nделение: %v\n",
			num1+num2, num1-num2, num1*num2, num1/num2))
	})
	v.buttonResult = btn
}

func (v *View) createKeys() *fyne.Container {

	vb1 := container.NewVBox(
		widget.NewButton(" ÷  ", func() {}),
		widget.NewButton(" ×  ", func() {}),
		widget.NewButton(" -  ", func() {}),
		widget.NewButton(" +  ", func() {}),
		widget.NewButton(" =  ", func() {}))
	vb2 := container.NewVBox(
		widget.NewButton(" %  ", func() {}),
		widget.NewButton(" 9  ", func() {}),
		widget.NewButton(" 6  ", func() {}),
		widget.NewButton(" 3  ", func() {}),
		widget.NewButton(" AC ", func() {}))
	vb3 := container.NewVBox(
		widget.NewButton(" -x  ", func() {}),
		widget.NewButton(" 8  ", func() {}),
		widget.NewButton(" 5  ", func() {}),
		widget.NewButton(" 2  ", func() {}),
		widget.NewButton(" .  ", func() {}))
	vb4 := container.NewVBox(
		widget.NewButton(" +x  ", func() {}),
		widget.NewButton(" 7  ", func() {}),
		widget.NewButton(" 4  ", func() {}),
		widget.NewButton(" 1  ", func() {}),
		widget.NewButton(" 0  ", func() {}))
	row := container.NewHBox(vb4, vb3, vb2, vb1)
	return row
}

func (v *View) createButtons() {
	mainFields := container.NewHBox(container.NewVBox(
		v.timer, v.infoEntry, v.resultLabel))
	rowEntry := container.NewVBox(v.expressionEntry, mainFields, v.createKeys())
	v.window.SetContent(rowEntry)
}

func (v *View) checkExpr() {
	for {
		if strings.Contains(v.expressionEntry.Text, "=") {
			v.resultLabel.SetText(v.vm.Calculate(v.expressionEntry.Text))
			log.Println("est ravno")
		} else {
			v.resultLabel.SetText("Ожидание ввода выражения")
			log.Println("net ravno")
		}
		time.Sleep(time.Second)
	}
}

func (v *View) Fire() {
	v.createBaseUI()
	v.showTime() //кнопку пока отключим
	// v.createButton()
	v.createResult()
	v.createEntry()
	go v.checkExpr()
	v.createButtons()
	v.window.Show()
	v.app.Run()
	tidyUp()
}
func (v *View) showTime() {
	timer := widget.NewLabel("")
	go func() {
		for {
			timer.SetText(time.Now().Format("15:04:05"))
			time.Sleep(time.Second)
		}
	}()
	v.timer = timer
}

func tidyUp() {
	fmt.Println("Exited")
}
