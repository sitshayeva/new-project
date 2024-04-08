
package main

import (
    "log"
    "os"
    "github.com/spf13/cobra"
    tb "gopkg.in/telebot.v3"
)

func main() {
    var rootCmd = &cobra.Command{
        Use:   "telegrambot",
        Short: "Telegram bot using Cobra and Telebot",
        Long: `A simple Telegram bot written in Go, 
               using Cobra and Telebot libraries.`,
        Run: func(cmd *cobra.Command, args []string) {
            runBot()
        },
    }

    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}

func runBot() {
    botToken := os.Getenv("TELE_TOKEN")
    if botToken == "" {
        log.Fatal("TELE_TOKEN environment variable not set")
    }

    botSettings := tb.Settings{
        Token: botToken,
    }

    bot, err := tb.NewBot(botSettings)
    if err != nil {
        log.Fatal(err)
        return
    }

    bot.Handle(tb.OnText, func(m tb.Context) error {
        // TODO: add logic to process text messages
        return m.Send("I received your message!")
    })

    bot.Start()
}
