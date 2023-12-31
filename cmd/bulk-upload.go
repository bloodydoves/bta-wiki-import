package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"cgt.name/pkg/go-mwclient"
)

var (
	flagApiURL  string
	flagWikiUsername string
	flagWikiPassFile string
)

var BulkUploadCmd = &cobra.Command{
	Use:   "bulk-upload",
	Short: "upload all pictures in current directory",
	Run: func(cmd *cobra.Command, args []string) {
		// first, check the flags for a username. Prefer this over the
		// environment variable.
		username := flagWikiUsername
		url := flagApiURL

		// if there is no username flag set, then check the environment
		if flagWikiUsername == "" {
			if user := os.Getenv(USERNAME_ENV); user != "" {
				username = user
			} else {
				return fmt.Errorf("no wiki username provided")
			}
		}

		password := ""
		// again, first check to see if the password is in a file provided by
		// flags.
		passFile := flagWikiPassFile
		// if the passfile is not empty, open and read that file, trimming off
		// spaces
		if flagWikiPassFile != "" {
			fileContents, err := ioutil.ReadFile(passFile)
			if err != nil {
				return err
			}
			password = strings.TrimSpace(string(fileContents))
		} else {
			password = os.Getenv(PASSWORD_ENV)
		}

		if password == "" && !flagDryRun {
			return fmt.Errorf("no wiki password provided")
		}

		// Create new mediawiki client & log in
		uploadClient, err := mwclient.New(url, "")
		if err != nil {
			return err
		}
		err = uploadClient.Login(username, password)
		csrfToken := uploadClient.GetToken(CSRFToken)

		uploadParams := map[string]string{
			"action":          "upload",
			"filename":        "file_1.jpg",
			"format":          "json",
			"token":           csrfToken,
			"ignorewarnings":  "1",
		}

		dir, err := os.Getwd()
		if err != nil {
			return err
		}
	
		// Use io.WalkDir to walk through the files in the current directory
		err = filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
	
			// Check if the file has one of the specified extensions
			ext := filepath.Ext(path)
			if ext == ".jpg" || ext == ".png" || ext == ".svg" {
				(*uploadParams)["filename"] = filepath.Base(path)
				// Call the arbitrary function with the file name
				return uploader.Upload(url, targetFile, uploadParams)
			}
			(*uploadParams)["filename"] = filepath.Base(path)
			
		})
	
	},
}

func init() {
	BulkUploadCmd.Flags().StringVarP(
		&flagApiURL, "url", "l", "",
		"the wiki URL to execute the cache purge against. Expects https://WEBSITE/api.php",
	)
	BulkUploadCmd.Flags().StringVarP(
		&flagWikiUsername, "username", "u", "",
		"the username to use when logging into the wiki",
	)
	BulkUploadCmd.Flags().StringVar(
		&flagWikiPassFile, "passfile", "",
		"a file to read the wiki password from",
	)
}