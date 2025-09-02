package cat

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

const (
	reset   = "\x1b[0m"
	white   = "\x1b[97m"
	yellow  = "\x1b[93m"
	magenta = "\x1b[95m"
)

func Run() {
	// Вернём курсор при Ctrl+C
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func() {
		<-ch
		fmt.Print("\x1b[?25h", reset, "\n")
		os.Exit(0)
	}()

	// Спрятать курсор
	fmt.Print("\x1b[?25l")
	defer fmt.Print("\x1b[?25h", reset)

	eyesFrames := []string{
		yellow + "o o" + reset,
		yellow + "- -" + reset,
		yellow + "o o" + reset,
		yellow + "^ ^" + reset,
	}
	tailFrames := []string{"", "~", "~=", "~==", "~=", "~"}

	for i := 0; ; i++ {
		eyes := eyesFrames[i%len(eyesFrames)]
		tail := tailFrames[i%len(tailFrames)]
		nose := magenta + "^" + reset

		// Очистить экран и в начало
		fmt.Print("\x1b[2J\x1b[H")

		fmt.Printf(white + "  /\\_/\\  \n" + reset)
		fmt.Printf(white+" ( %s ) \n"+reset, eyes)
		fmt.Printf(white+" =(  %s  )=   %s\n"+reset, nose, tail)
		fmt.Printf(white + "  )   (  //\n" + reset)
		fmt.Printf(white + " (__ __)//\n" + reset)
		fmt.Printf("\nНажми Ctrl+C, чтобы выйти.\n")

		time.Sleep(200 * time.Millisecond)
	}
}
