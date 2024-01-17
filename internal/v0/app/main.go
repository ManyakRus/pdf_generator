// главный модуль программы

package main

import (
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/contextmain"
	logger "github.com/ManyakRus/starter/logger"
	stopapp "github.com/ManyakRus/starter/stopapp"
	//	// "github.com/ManyakRus/starter/claim_debtors_list/internal/v0/app/config"
)

// // log - глобальный логгер
var log = logger.GetLog()

// main - старт приложения
func main() {
	StartApp()
}

// StartApp - выполнение всех операций для старта приложения
func StartApp() {
	config_main.LoadEnv()

	stopapp.StartWaitStop()

	contextmain.GetContext()

	stopapp.GetWaitGroup_Main().Wait()

}
