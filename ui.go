package main

import (
	"bufio"
	"fmt"
//	"io"
	"log"
	"os"
	"os/user"
	"strings"
	//    "io"
)

const goodStr = "👍"
const badStr = "🤔"

var Status = true
var commandCount int
var scripting bool

func Init() {
    Status = true
    commandCount = 0
    stat, _ := os.Stdin.Stat()
    if (stat.Mode() & os.ModeCharDevice) == 0 {
        log.Print("Data piped in on stdin")
        scripting = true;
    }
}

func PromptLine() string {
    var statusString string

    if Status {
        statusString = goodStr
    } else {
        statusString = badStr
    }
    
    user := promptUsername()
    host := promptHostname()
    cwd := promptCWD()
    commandNum := commandNumber()

    formatString := "%s[%d] %s@%s:%s "
    promptString := fmt.Sprintf(formatString, statusString, commandNum, user, host, cwd)
    return promptString
}

func promptUsername() string {
    user, _ := user.Current()
    return user.Username
}

func promptHostname() string {
    hostname, _ := os.Hostname()
    return hostname
}

func promptCWD() string {
    cwd, _ := os.Getwd()
    home, _ := os.UserHomeDir()

    if (strings.HasPrefix(cwd, home)) {
        cwd = strings.Replace(cwd, home, "~", 1)
    }

    return cwd
}

func commandNumber() int {
    return commandCount
}

func ReadCommand() (string, bool) {
    reader := bufio.NewReader(os.Stdin)
    scanner := bufio.NewScanner(os.Stdin)
    commandCount++
    if !scripting {
        fmt.Println(PromptLine())
        command, _ := reader.ReadString('\n')
        return strings.TrimSuffix(command, "\n"), false
    } else {
        scanner.Scan()
        command := scanner.Text()
        return strings.TrimSuffix(command, "\n"), scanner.Err() == nil
    }
}
