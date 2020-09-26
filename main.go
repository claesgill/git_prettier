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
// TODO: Add flag -b branch -m modidied -d deleted -u untraced (-bmdu)

func main() {
	vers   := flag.Bool("v", false, "get the version")

	status := flag.Bool("status", false, "formatting git status if any")
	flag.Parse()
	if *vers == true {
		fmt.Println("Git Prettier version v0.0.2")
		os.Exit(1)
	}
	if *status {
		nMod, nDel, nUntr := gitStatus()
		branch 			  := gitBranch()
	
		if branch == "" { os.Exit(1) }
		
		fmtStats := fmtResults(branch, nMod, nDel, nUntr)
		fmt.Println(fmtStats)
	}
}

func gitBranch() (string){
	cmd := exec.Command("git", "branch")
	output, err := cmd.Output()
	if err != nil {}
	rBranch    := regexp.MustCompile(`^\*\s.*`)
	branch     := rBranch.FindAllString(string(output[:]), 1000)
	branch     = strings.Split(strings.Join(branch, ""), "* ")
	parsBranch := strings.Join(branch, "")
	return strings.TrimSuffix(parsBranch, "\n")
}

func gitStatus() (int, int, bool){
	cmd := exec.Command("git", "status")
	output, err := cmd.Output()
	if err != nil {}
	rMod  := regexp.MustCompile(`modified:`)
	nMod  := len(rMod.FindAllString(string(output[:]), 1000))
	rDel  := regexp.MustCompile(`deleted:`)
	nDel  := len(rDel.FindAllString(string(output[:]), 1000))
	rNew  := regexp.MustCompile(`new file:`)
	nNew  := len(rNew.FindAllString(string(output[:]), 1000))
	rUntr := regexp.MustCompile(`Untracked files:`)
	nMod = nMod + nNew
	return nMod, nDel, rUntr.MatchString(string(output[:]))
}

func fmtResults(branch string, nMod int, nDel int, isUntr bool) (string) {
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