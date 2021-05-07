package cmd

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"io"
	"math"
	"os"
	"os/exec"
	"strconv"
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
		fmt.Println("install called")
		createFolders()
		if !skipDownload {
			downloadQASInstall()
		}

		startInstaller()
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

	s := spinner.New(spinner.CharSets[26], 100*time.Millisecond)
	s.Prefix = fmt.Sprintf("Size of QAS Installer %s. \nDownloading ", fileSize(stat.Size()))
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
