package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

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
	Long:  `Clone all holaluz repositories you want`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Sprintf("git@github.com:holaluz/%s.git", args[0])
		clone(args[0], "git@github.com:holaluz/"+args[0]+".git")
	},
}

func getSSHKeyAuth(privateSSHKeyFile string) transport.AuthMethod {
	var auth transport.AuthMethod
	sshKey, _ := ioutil.ReadFile(privateSSHKeyFile)
	signer, _ := ssh.ParsePrivateKey([]byte(sshKey))
	auth = &go_git_ssh.PublicKeys{User: "git", Signer: signer}
	return auth
}

func clone(projectName string, url string) {
	directory := viper.GetString("code_path") + "/" + projectName

	fmt.Println(directory)

	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:      url,
		Auth:     getSSHKeyAuth(viper.GetString("github_key_path")),
		Progress: os.Stdout,
	})

	CheckIfError(err)

	fmt.Println("Cloned")
}
