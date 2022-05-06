package main

import (
    "fmt"
    "strings"
    "os"
    "log"
    "strconv"
    "os/exec"
)

type commandStruct struct {
    tokens []string
    stdoutFile string
    stdinFile string
    stdoutPipe bool
}

func executeCommand(command string) bool {
    tokens := strings.Split(command, " ")
    cmd := exec.Command(tokens[0], tokens[1:]...)
    out, err := cmd.Output()

    fmt.Print(string(out))
    if (err != nil) {
        log.Println(err)
    }
    return err == nil
}
//func executePipeline()

func commandHandler (command string) (bool, bool) {
    commentIndex := strings.Index(command, "#")
    var actualCommand string
    if commentIndex != -1 {
       actualCommand = command[0 : commentIndex]
    } else {
        actualCommand = command
    }
    commandTokens := strings.Split(actualCommand,  " ")
    if commandTokens[0] == "exit" {
        return true, true;

    } else if commandTokens[0] == "" {
        return false, true;

    } else if commandTokens[0] == "history" {
        HistAdd(command)
        HistPrint()
        return false, true

    } else if commandTokens[0] == "cd" {
        if tilda := strings.Index(commandTokens[1], "~"); tilda != -1 {
            home, _ := os.UserHomeDir()
            commandTokens[1] = strings.Replace(commandTokens[1], "~", home, 1)
        }
        err := os.Chdir(commandTokens[1])
        if (err != nil) {
            log.Print(err)
        }
        HistAdd(command)
        return false, err == nil

    } else if commandTokens[0][0] == '!' {
        num, err := strconv.Atoi(commandTokens[0][1:])
        var result string
        var isFound bool

        if err != nil {
            if commandTokens[0][1] == '!' {
                result, _ = HistSearchCnum(HistLastCnum())
            } else {
                result, isFound = HistSearchPrefix(commandTokens[0][1:])
                if !isFound {
                    return false, false
                }
            }

        } else {
            result, isFound = HistSearchCnum(num)
            if !isFound {
                return false, false
            }

        }
//        log.Println(result)
        return commandHandler(result)

    }else {
        HistAdd(command)

        return false, executeCommand(command)
    }

}

func main() {
    Init()
//    fmt.Println("Hello shell")
    for {

        command, end := ReadCommand()
        e, s := commandHandler(command)
        Status = s
        if e || end {
           break
       }

   }

}
