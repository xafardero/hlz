package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

var name string
var env string
var app string
var region string
var user string

func init() {
	sshCmd.Flags().StringVarP(&env, "env", "e", "*", "env")
	sshCmd.Flags().StringVarP(&app, "app", "a", "*", "app")
	sshCmd.Flags().StringVarP(&region, "region", "r", "eu-west-1", "region")
	sshCmd.Flags().StringVarP(&user, "user", "u", "", "user")
	rootCmd.AddCommand(sshCmd)
}

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "ssh to your aws instances.",
	Long:  `select and enter to your aws instances via ssh.`,
	Run: func(cmd *cobra.Command, args []string) {
		var ec2 string

		if len(args) != 0 {
			ec2 = args[0]
		} else {
			ec2 = "*"
		}

		rawInstanceList := filterInstances(
			cmd.Flag("region").Value.String(),
			cmd.Flag("env").Value.String(),
			cmd.Flag("app").Value.String(),
			ec2,
		)
		instances := getInstancesInfo(rawInstanceList)
		showInstanceList(instances, cmd.Flag("user").Value.String())
		selectedInstance := selectInstanceIndex(instances)
		launchSSH(instances[selectedInstance].IP, cmd.Flag("user").Value.String())
	},
}

//Instance struct
type Instance struct {
	Name string
	IP   string
	ID   string
	Size string
}

//NewInstance creator
func NewInstance(name, ip, id, size string) *Instance {
	newInstance := new(Instance)
	newInstance.Name = name
	newInstance.IP = ip
	newInstance.ID = id
	newInstance.Size = size
	return newInstance
}

func selectInstanceIndex(instanceList []Instance) int {
	var input string
	fmt.Println("\nWhich one do you want to ssh in?")
	fmt.Scanln(&input)
	index, err := strconv.Atoi(input)
	if err != nil || index > len(instanceList)-1 || index < 0 {
		fmt.Println("Please enter a valid number.")
		index = selectInstanceIndex(instanceList)
	}
	return index
}

func showInstanceList(instanceList []Instance, user string) {
	if len(instanceList) == 0 {
		fmt.Printf("There are no instances matching your request.\n")
		os.Exit(0)
	}
	for idx, inst := range instanceList {
		fmt.Printf("[%d] %s - %s (%s, %s) \n", idx, inst.Name, inst.IP, inst.ID, inst.Size)
	}
	if len(instanceList) == 1 {
		fmt.Printf("\n\n")
		launchSSH(instanceList[0].IP, user)
		os.Exit(0)
	}
}

func launchSSH(destIP string, username string) {
	destination := destIP
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	time.Sleep(1 * time.Second)
	if username != "" {
		destination = username + "@" + destination
	}

	fmt.Printf(">> Starting a new ssh session to %s\n", destination)
	proc, err := os.StartProcess("/usr/bin/ssh", []string{"ssh", destination}, &pa)

	if err != nil {
		panic(err)
	}

	state, err := proc.Wait()
	if err != nil {
		panic(err)
	}

	fmt.Printf("<< Exited ssh session: %s\n", state.String())
}

func getInstancesInfo(describeInstancesOutput *ec2.DescribeInstancesOutput) []Instance {
	var instanceList []Instance
	for idx := range describeInstancesOutput.Reservations {
		for _, inst := range describeInstancesOutput.Reservations[idx].Instances {
			name := ""
			for _, tag := range inst.Tags {
				if *tag.Key == "Name" {
					name = *tag.Value
				}
			}
			newInstance := NewInstance(name, *inst.PrivateIpAddress, *inst.InstanceId, *inst.InstanceType)
			instanceList = append(instanceList, *newInstance)
		}
	}
	return instanceList
}

func filterInstances(region, env, app, name string) *ec2.DescribeInstancesOutput {

	fmt.Printf("\nApplication: %s   Environment: %s   Name: %s   Region: %s\n", app, env, name, region)
	fmt.Printf("---------------------------------------------------------\n\n")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	ec2svc := ec2.New(sess)
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag:Environment"),
				Values: []*string{aws.String(env)},
			},
			{
				Name:   aws.String("tag:Name"),
				Values: []*string{aws.String("*" + name + "*")},
			},
			{
				Name:   aws.String("tag:Application"),
				Values: []*string{aws.String(app + "*")},
			}, {
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running"), aws.String("pending")},
			},
		},
	}
	resp, err := ec2svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("there was an error listing instances in", err.Error())
		log.Fatal(err.Error())
	}
	return resp
}
