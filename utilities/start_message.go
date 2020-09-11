package utilities
import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"strings"
)

func printLineBreak(messageLength int) {
	fmt.Println(strings.Repeat("=", messageLength))
}

func PrintStartMessage(port string) {
	message := figure.NewFigure("BOILERPLATE SERVICE", "", true)
	message.Print()

	portMessage := "Application Listening on Port:" + port
	printLineBreak(len(portMessage))
	fmt.Println(portMessage)
	printLineBreak(len(portMessage))
}
