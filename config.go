package main

import (
	"fmt"
	"os"
	"strings"

	cl "bitbucket.org/ai69/so-colorful/colorlogo"
	"github.com/1set/gut/yos"
	"github.com/1set/gut/ystring"
)

var (
	AppName    string
	CIBuildNum string
	BuildDate  string
	BuildHost  string
	GoVersion  string
	GitBranch  string
	GitCommit  string
	GitSummary string
)

var (
	logoArt = `
   ██████╗██╗     ██╗
  ██╔════╝██║     ██║
  ██║     ██║     ██║
  ██║     ██║     ██║
  ╚██████╗███████╗██║
   ╚═════╝╚══════╝╚═╝
`
	logoArtColor = cl.CherryBlossomsByColumn(logoArt)
)

func displayBuildInfo() {
	if ystring.IsBlank(BuildDate) {
		return
	}

	var sb strings.Builder
	flag := "➣ "
	if yos.IsOnWindows() {
		flag = "> "
	} else {
		sb.WriteString(logoArtColor)
	}

	addNonBlankField := func(name, value string) {
		if ystring.IsNotBlank(value) {
			fmt.Fprintln(&sb, flag+name+":", value)
		}
	}

	addNonBlankField("Build Num ", CIBuildNum)
	addNonBlankField("Build Date", BuildDate)
	addNonBlankField("Build Host", BuildHost)
	addNonBlankField("Go Version", GoVersion)
	addNonBlankField("Git Branch", GitBranch)
	addNonBlankField("Git Commit", GitCommit)
	addNonBlankField("GitSummary", GitSummary)

	rs := sb.String() + ystring.NewLine
	os.Stderr.WriteString(rs)
}
