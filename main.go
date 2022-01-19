package main

import (
	"log"
	"os"
	"regexp"
	"strconv"

    // Could add more message types later
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func checkErr(err error) {
    if err != nil {
        log.Fatalln(err)
    }
}

func main() {
    // Check RUNNER_OS
    if runnerOS := os.Getenv("RUNNER_OS"); runnerOS != "Linux" {
        log.Fatalln("Only Linux runners are supported")
    }

    // Get inputs
    changelog := os.Getenv("INPUT_CHANGELOG")
    token := os.Getenv("INPUT_TOKEN")
    channel := os.Getenv("INPUT_CHANNEL")

    // Validate inputs
    if changelog == "" || token == "" || channel == "" {
        log.Fatalln("Insufficient parameters")
    }

    channel_id, err := strconv.Atoi(channel)
    checkErr(err)

    bot, err := telegram.NewBotAPI(token)
    checkErr(err)

    log.Printf("Using %s to post changelog...", bot.Self.UserName)

    // Construct message, order of replacements is important
    replaceString := func(regex, repl string) {
        changelog = regexp.MustCompile(regex).ReplaceAllString(changelog, repl)
    }

    replaceString(
        " by @\\w+ in https:\\/\\/github\\.com\\/[\\w\\-]+\\/[\\w\\-]+\\/pull\\/\\d+",
        "",
    )  // Strip contributions
    replaceString("<!--.*-->", "")  // Comments
    replaceString("(?:^|\n)#+ (.+)(\n|$)", "*$1*$2")  // Headings
    replaceString("((?:^|\n)\\s*)[\\*\\-] ", "$1• ")  // Lists
    replaceString("\\*+", "*")  // Bolds
    replaceString("_+", "_")  // Italics
    replaceString("\\.", "\\.")  // Dots, yes the 2 strings look the same
    replaceString("-", "\\-")  // Hyphens
    replaceString("!", "\\!")  // Bangs
    replaceString("`", "\\`")  // Tildes

    msg := telegram.NewMessage(int64(channel_id), changelog)
    msg.DisableWebPagePreview = true
    msg.ParseMode = "MarkdownV2"

    // Send message
    _, err = bot.Send(msg)
    checkErr(err)

    log.Println("Changelog posted to channel")
}
