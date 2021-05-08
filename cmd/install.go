package cmd

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"golang.org/x/sys/windows"
	"io"
	"math"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var skipDownload bool
var sourceFilePath = "\\\\i90256\\Quintiq 609\\QAS-Setup-64-6_0_9.exe"
var destFilePath = "C:\\Program Files\\Quintiqdev\\QAS-Setup-64-6_0_9.exe"

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install GUI for Dev and BCTest on your machine",
	Long:  `Install GUI for Dev and BCTest on your machine`,
	Run: func(cmd *cobra.Command, args []string) {

		if !skipDownload {
			downloadQASInstall()
		}

		if !amAdmin() {
			runMeElevated()
		}
		time.Sleep(10 * time.Second)

	},
	PreRun: func(cmd *cobra.Command, args []string) {
		dirname, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		destFilePath = path.Join(dirname, "Downloads", "QAS-Setup-64-6_0_9.exe")
	},
}

var (
	sizeInMB float64 = 999 // This is in megabytes
	suffixes [5]string
)

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.PersistentFlags().BoolVar(&skipDownload, "skip-download", false, "Skip download of QAS Setup File")
}

func createFolders() {
	err := os.MkdirAll("C:\\Program Files\\Quintiqdev\\Quintiq 60", 0777)
	if err != nil {
		panic(err)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}
}

func fileSize(filesize int64) string {
	size := float64(filesize)
	suffixes[0] = "B"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[int(math.Floor(base))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}

func downloadQASInstall() {

	srcFile, err := os.Open(sourceFilePath)
	check(err)
	defer srcFile.Close()

	stat, err := srcFile.Stat()

	check(err)
	fmt.Printf("Size of QAS Installer %s.\n", fileSize(stat.Size()))

	s := spinner.New(spinner.CharSets[26], 100*time.Millisecond)
	s.Prefix = "Downloading "
	s.Start()

	destFile, err := os.Create(destFilePath) // creates if file doesn't exist
	check(err)
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		panic(err)
	}

	s.Stop()
}

func startInstaller() {
	cmd := exec.Command(destFilePath)
	err := cmd.Run()
	check(err)
}

func runMeElevated() {
	verb := "runas"
	//exe, _ := os.Executable()
	//cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(destFilePath)
	cwdPtr, _ := syscall.UTF16PtrFromString(path.Join(dirname, "Downloads"))
	argPtr, _ := syscall.UTF16PtrFromString(args)

	fmt.Println(verbPtr, exePtr, cwdPtr, argPtr)

	var showCmd int32 = 0 //SW_NORMAL

	err = windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func amAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}
