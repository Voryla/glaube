package berouter

import "github.com/fatih/color"

var logo = "  _____ _             _            /\\      /\\ \n / ____| |           | |          |/\\|  _ |/\\|\n| |  __| | __ _ _   _| |__   ___       (_)    \n| | |_ | |/ _` | | | | '_ \\ / _ \\             \n| |__| | | (_| | |_| | |_) |  __/     _ _ _   \n \\_____|_|\\__,_|\\__,_|_.__/ \\___|    (_|_|_)  "

func printLogo() {
	color.Blue(logo)
}
