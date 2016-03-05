package goan

import (
	"testing"
)

func Test_Logging(t *testing.T) {
	SetupLogger()
	LogInfo.Println("Loggin info")
	LogWarning.Println("Logging warning")
	LogError.Println("Logging error")
}
