package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	go_git_ssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

func init() {
	rootCmd.AddCommand(cloneCmd)
}

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone holaluz repositories",
	Long:  `Clone holaluz repositories from github`,
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf("git@github.com:holaluz/%s.git", args[0])
		clone(args[0], url)
	},
}

func clone(projectName string, url string) {
	directory := getProjectsDirectory(projectName)

	fmt.Printf("\n%s\n\n", directory)

	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:      url,
		Auth:     getSSHKeyAuth(getGithubSSHKey()),
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Println(fmt.Errorf("Repository clone fail because %s", err))
		os.Exit(1)
	}

	fmt.Printf("\nCloned %s in directory %s\n", projectName, directory)
}

func getSSHKeyAuth(privateSSHKeyFile string) transport.AuthMethod {
	var auth transport.AuthMethod
	sshKey, _ := ioutil.ReadFile(privateSSHKeyFile)
	signer, _ := ssh.ParsePrivateKey([]byte(sshKey))
	auth = &go_git_ssh.PublicKeys{User: "git", Signer: signer}
	return auth
}

func getGithubSSHKey() string {
	configSSHKey := viper.GetString("github_key_path")

	if configSSHKey != "" {
		return configSSHKey
	}

	home, _ := homedir.Dir()

	return fmt.Sprintf("%s/.ssh/id_rsa", home)
}
