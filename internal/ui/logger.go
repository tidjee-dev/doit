package ui

import "fmt"

func Success(msg string) {
	fmt.Println("✅", SuccessStyle.Render("[SUCCESS]"), msg)
}

func Warn(msg string) {
	fmt.Println("⚠️ ", WarnStyle.Render("[WARN]"), msg)
}

func Info(msg string) {
	fmt.Println("ℹ️ ", InfoStyle.Render("[INFO]"), msg)
}

func Error(msg string) {
	fmt.Println("❌", ErrorStyle.Render("[ERROR]"), msg)
}
