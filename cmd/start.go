package cmd

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start SOPRE UI's",
	Long: `Start SOPRE UI's over the cli. 

Start a ui with: sopre start <platform> <dataset> <ui>
Example: sopre start d ap client

All Shortcuts created while installing local model are supported:

Platforms: 	d, dev, b, bc, bctests
UI's: 		designer, client, editor
datasets: 	dev
			  - ap
			  - ep
			bc
			  - ep
			  - global
			  - le
			  - plst
`,
	Run: func(cmd *cobra.Command, args []string) {

		runos := runtime.GOOS

		if runos != "windows" {
			color.Red("ERROR: This software only runs on Window Machines\n\n")
			os.Exit(1)
		}

		if len(args) < 3 {
			color.Red("ERROR: Too few arguments supplied to start a sopre ui\n\n")
			_ = cmd.Help()
			os.Exit(0)
		}
		openApplication(args)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.SetUsageTemplate("Usage:\n" +
		"\tsopre start <platform> <dataset> <ui>\n" +
		"\n" +
		"Flags:\n" +
		"\t-h, --help\t\t\thelp for start\n" +
		"\t--config string\t\tconfig file (default is $HOME/.sopre.yaml)\n")
}

func openApplication(args []string) {

	pf := args[0]
	ds := args[1]
	exe := args[2]

	platform := ""

	var p []string
	p = append(p, "dev")
	p = append(p, "d")
	p = append(p, "bctest")
	p = append(p, "b")
	p = append(p, "bc")

	var dts []string
	dts = append(dts, "ap")
	dts = append(dts, "ep")
	dts = append(dts, "global")
	dts = append(dts, "le")
	dts = append(dts, "plst")

	var e []string
	e = append(e, "designer")
	e = append(e, "client")
	e = append(e, "editor")

	if !stringInSlice(pf, p) {
		fmt.Printf(fmt.Sprintf("%s is not a possible plattform", pf))
		os.Exit(0)
	}

	if pf == "d" || pf == "dev" {
		platform = "d"
	}
	if pf == "b" || pf == "bctest" || pf == "bc" {
		platform = "bc"
	}

	if !stringInSlice(ds, dts) {
		fmt.Printf(fmt.Sprintf("%s is not a possible dataset", pf))
		os.Exit(0)
	}

	if !stringInSlice(exe, e) {
		fmt.Printf(fmt.Sprintf("%s is not a possible executable", pf))
		os.Exit(0)
	}

	installPath := viper.GetString("sopre-install-path")
	executable := viper.GetString(fmt.Sprintf("sopre-gui-path.%s.%s.%s", platform, ds, exe))
	exeSt := fmt.Sprintf("%s.exe", executable)
	gui := fmt.Sprintf("%s\\%s", installPath, executable)
	propertiesPath := viper.GetString("sopre-properties-path")
	var qtq = exec.Command("")
	properties := fmt.Sprintf("/startup.properties=%s\\%s.properties", propertiesPath, executable)
	if exe == "client" || exe == "editor" {
		qtq = exec.Command(gui, properties)
	} else {
		qtq = exec.Command(gui)
	}

	qtq.Dir = viper.GetString("sopre-install-path")
	qtq.Path = exeSt
	qtq.Args[0] = exeSt

	s := spinner.New(spinner.CharSets[26], 100*time.Millisecond)
	s.Prefix = fmt.Sprintf("Starting %s ", executable)
	s.Start()

	qtq_err := qtq.Start()
	if qtq_err != nil {
		panic(qtq_err)
	}

	pid := qtq.Process.Pid

	time.Sleep(5 * time.Second)
	s.Stop()
	fmt.Printf("started %s [%d]\n\n", executable, pid)

}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
