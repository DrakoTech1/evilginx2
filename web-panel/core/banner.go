package core

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	VERSION = "3.0.2"
)

func putAsciiArt(s string) {
	for _, c := range s {
		d := string(c)
		switch string(c) {
		case " ":
			color.Set(color.BgRed)
			d = " "
		case "@":
			color.Set(color.BgBlack)
			d = " "
		case "#":
			color.Set(color.BgHiRed)
			d = " "
		case "W":
			color.Set(color.BgWhite)
			d = " "
		case "_":
			color.Unset()
			d = " "
		case "\n":
			color.Unset()
		}
		fmt.Print(d)
	}
	color.Unset()
}

func printLogo(s string) {
	for _, c := range s {
		d := string(c)
		switch string(c) {
		case "_":
			color.Set(color.FgWhite)
		case "\n":
			color.Unset()
		default:
			color.Set(color.FgHiBlack)
		}
		fmt.Print(d)
	}
	color.Unset()
}

func printUpdateName() {
	nameClr := color.New(color.FgHiWhite)
	txt := nameClr.Sprintf("               - --  DEATH STROK  -- -")
	fmt.Fprintf(color.Output, "%s", txt)
}

func printOneliner1() {
	handleClr := color.New(color.FgHiBlue)
	versionClr := color.New(color.FgGreen)
	textClr := color.New(color.FgHiBlack)
	spc := strings.Repeat(" ", 10-len(VERSION))
	txt := textClr.Sprintf("      by SMARTGRENEDIER (") + handleClr.Sprintf("@SMARTGRENEDIER") + textClr.Sprintf(")") + spc + textClr.Sprintf("version ") + versionClr.Sprintf("%s", VERSION)
	fmt.Fprintf(color.Output, "%s", txt)
}

func printOneliner2() {
	textClr := color.New(color.FgHiBlack)
	red := color.New(color.FgRed)
	white := color.New(color.FgWhite)
	txt := textClr.Sprintf("                   no ") + red.Sprintf("nginx") + white.Sprintf(" - ") + textClr.Sprintf("Hell ") + red.Sprintf("Fire")
	fmt.Fprintf(color.Output, "%s", txt)
}

func Banner() {
	fmt.Println()

	putAsciiArt("                                                               \n")
	putAsciiArt("  _____  ______       _______ _    _  _____ _______ _____   ______   __
	              |  __ \|  ____|   /\|__   __| |  | |/ ____|__   __|  __ \ / __ \ \ / /
	              | |  | | |__     /  \  | |  | |__| | (___    | |  | |__) | |  | \ V / 
	              | |  | |  __|   / /\ \ | |  |  __  |\___ \   | |  |  _  /| |  | |> <  
	              | |__| | |____ / ____ \| |  | |  | |____) |  | |  | | \ \| |__| / . \ 
	              |_____/|______/_/    \_\_|  |_|  |_|_____/   |_|  |_|  \_\\____/_/ \_\
																		                 ")
	printLogo(`                                                                          `)
	fmt.Println()
	putAsciiArt("__                                                                     \n")
	fmt.Println()
}
