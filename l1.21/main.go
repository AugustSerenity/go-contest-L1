package main

import "fmt"

// интерфейс клиента
type Computer interface {
	InsertIntoLightningPort()
}

type Client struct {
}

// клиентский код
func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

// Сервис
type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

// Что-то не известное
type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

// Адаптер
type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsAdapter)
}

// Применимость
//  Когда вы хотите использовать сторонний класс, но его интерфейс не соответствует остальному коду приложения.
//  Адаптер позволяет создать объект-прокладку, который будет превращать вызовы приложения в формат,
//  понятный стороннему классу.
//  Когда вам нужно использовать несколько существующих подклассов,
//  но в них не хватает какой-то общей функциональности, причём расширить суперкласс вы не можете.

// Плюсы
// Отделяет и скрывает от клиента подробности преобразования различных интерфейсов.

// Минусы
// Усложняет код программы из-за введения дополнительных классов.

// Реальное применение:
// Приложение получает данные в формате xml
// необходимо подключить сторонню библиотеку для более детальной аналитики
// для расширения функциональности приложения, но данная библиотека работает с форматом данных JSON
// поэтому в наше приложение добавляем Адаптер.
// Адаптер получает эти вызовы и перенаправляет их второму объекту,
// но уже в том формате и последовательности, которые понятны второму объекту.
// Наше приложение посылало бы адаптеру запросы в формате XML,
// а адаптер сначала транслировал входящие данные в формат JSON,
// а затем передавал бы их методам обёрнутого объекта аналитики
