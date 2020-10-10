package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"flag"
)

// TODO: implement random color func

func main() {
	vers      := flag.Bool("v", false, "get the version")
	status    := flag.Bool("status", false, "formatting git status if any")
	branch    := flag.Bool("b", false, "get active branch")
	modified  := flag.Bool("m", false, "get number of tracked modified files")
	deleted   := flag.Bool("d", false, "get number of tracked deleted files")
	untracked := flag.Bool("u", false, "get number of untracked files")
	color     := flag.Bool("c", false, "color the status output")

	flag.Parse()
	if *vers == true {
		fmt.Println("Git Prettier version v0.0.3")
		os.Exit(1)
	}
	if *status {
		nMod, nDel, isUntr := gitStatus()
		branch 			  := gitBranch()
		if branch == "" { os.Exit(1) }
		var fmtStats string
		if *color {
			fmtStats = fmtResultsColor(branch, nMod, nDel, isUntr)
		} else {
			fmtStats = fmtResults(branch, nMod, nDel, isUntr)
		}
		fmt.Println(fmtStats)
	}
	if *branch    { fmt.Println(gitBranch()) }
	if *modified  { 
		if nMod, _, _ := gitStatus(); nMod > 0 { fmt.Printf("+%v\n", nMod) }
	}
	if *deleted   {
		if _, nDel, _ := gitStatus(); nDel > 0 { fmt.Printf("-%v\n", nDel) }
	}
	if *untracked {
		if _, _, isUntr := gitStatus(); isUntr { fmt.Println("U…") }
	}
}

func gitBranch() (string){
	cmd := exec.Command("git", "branch")
	output, err := cmd.Output()
	if err != nil {}
	rBranch    := regexp.MustCompile(`\*\s.*`)
	branch     := rBranch.FindAllString(string(output[:]), 1000)
	branch     = strings.Split(strings.Join(branch, ""), "* ")
	parsBranch := strings.Join(branch, "")
	return strings.TrimSuffix(parsBranch, "\n")
}

func gitStatus() (int, int, bool){
	cmd := exec.Command("git", "status")
	output, err := cmd.Output()
	if err != nil {}
	rMod := regexp.MustCompile(`modified:`)
	nMod := len(rMod.FindAllString(string(output[:]), 1000))
	rNew := regexp.MustCompile(`new file:`)
	nNew := len(rNew.FindAllString(string(output[:]), 1000))
	nMod =  nMod + nNew
	rDel  := regexp.MustCompile(`deleted:`)
	nDel  := len(rDel.FindAllString(string(output[:]), 1000))
	rUntr := regexp.MustCompile(`Untracked files:`)
	return nMod, nDel, rUntr.MatchString(string(output[:]))
}

func fmtResults(branch string, nMod int, nDel int, isUntr bool) (string) {
	stats := "["
	stats += branch
	if nMod  > 0 || nDel > 0 { stats += " |" }
	if nMod  > 0 { stats = fmt.Sprint(stats, " +", nMod) }
	if nDel  > 0 { stats = fmt.Sprint(stats, " -", nDel) }
	if isUntr    { stats = fmt.Sprint(stats, " U…")      }
	stats += "]"
	return stats
}

func fmtResultsColor(branch string, nMod int, nDel int, isUntr bool) (string) {
	RED     := "\033[91m"
	BLUE    := "\033[34m"
	GREEN   := "\033[32m"
	PURPLE  := "\033[95m"
	DEFAULT := "\033[0m"

	stats := "["
	stats += PURPLE + branch + DEFAULT
	if nMod  > 0 || nDel > 0 { stats += " |" }
	if nMod  > 0 { stats = fmt.Sprint(stats, GREEN, " +",  nMod,  DEFAULT) }
	if nDel  > 0 { stats = fmt.Sprint(stats, RED,   " -",  nDel,  DEFAULT) }
	if isUntr    { stats = fmt.Sprint(stats, BLUE,  " U…", DEFAULT)        }
	stats += "]"
	return stats
}