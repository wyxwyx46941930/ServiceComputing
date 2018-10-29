// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "helo user to do something",
	Long: `

	you can use this app to create or remove meetings.Also you must register a user to have the rights to use the functions.

	Usage:
		agenda [command]

	Available Commands:
		user : commands about user operation
		meeting : commands about meeting operation

	Use "agenda [command] --help" for more information about a command.

	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			if args[0] == "user" {
				if args[1] == "register" {
					fmt.Println(`
Command : user register
Function: Registe a user.
instruction: user register --username/-u [UserName] --password/-p [Pass] --email/-e [Email] --telphone/-t [Phone]
Args :
	[Username] 	register's name
	[Pass] 		register's Password
	[Email] 	register's email
	[Phone] 	register's phone-number`)
				} else if args[1] == "delete" {
					fmt.Println(`
Command : user delete
Function: Delete the current user.
instruction: user delete
`)
				} else if args[1] == "lookup" {
					fmt.Println(`
Command : user lookup
Function: Lookup all the information of the users. (Need to login first)
instruction: user lookup
`)
				} else if args[1] == "login" {
					fmt.Println(`
Command : user login
Function: Login
(PS: If you don't login, you can only use user login and user register commands.)
instruction : user login --username/-s [UserName] --password/-p [Pass]
Args :
	[Username] 	register's name
	[Pass] 		register's Password`)
				} else if args[1] == "logout" {
					fmt.Println(`
Command : user logout
Function: Logout
instruction: Logout the user
`)
				}
			} else if args[0] == "meeting" {
				if args[1] == "addUser" {
					fmt.Println(`
Command : meeting addUser
Function: Add users to a meeting. It require all the users don't have meetings at the period this meeting held.
instruction: meeting addUser --title/-t [Title] --participant/-p [Participants]
Args :
	[Title] title of the meeting (unique)
	[Participants] Participants of the meetings
(PS: If multiple participants add at the same time, need multiple -p flag.
This : meeting addUser -p "wty" "wyx" is not available
But this : meeting addUser -p "wty" -p "wyx" is available)
`)
				} else if args[1] == "deleteUser" {
					fmt.Println(`
Command : meeting deleteUser
Function: Delete user from the meeting. If this operation make the meeting empty, then delete the meeting.
instruction: meeting deleteUser --title/-t [Title] --participant/-p [Participants]
Args :
	[Title] title of the meeting (unique)
	[Participants] Participants of the meetings
(PS: If multiple participants delete at the same time, need multiple -p flag.
This : meeting deleteUser -p "wty" "wyx" is not available
But this : meeting deleteUser -p "wty" -p "wyx" is available)
`)
				} else if args[1] == "create" {
					fmt.Println(`
Command : meeting create
Function: Create a meeting.
instruction: meeting create --startTime/-s [StartTime] --endTime/-e [EndTime] --title/-t [Title] --participant/-p [Participants]
Args :
	[Title] title of the meeting (unique)
	[Participants] Participants of the meetings
	[StartTime] the StartTime of the meeting
	[EndTime] the EndTime of the meeting 
(PS: If multiple participants delete at the same time, need multiple -p flag.
This : meeting create -p "wty" "wyx" is not available
But this : meeting create -p "wty" -p "wyx" is available)
`)
				} else if args[1] == "cancel" {
					fmt.Println(`
Command : meeting cancel
Function: User can cancel the meeting created by himself.
instruction: meeting cancel --title/-t [Title]
Args :
	[Title] title of the meeting (unique)
`)
				} else if args[1] == "lookup" {
					fmt.Println(`
Command : meeting lookup
Function: Login user can look up the meeting he attended in the specific period.
instruction: meeting create --startTime/-s [StartTime] --endTime/-e [EndTime] 
Args :
	[StartTime] the StartTime of the meeting
	[EndTime] the EndTime of the meeting 
`)
				} else if args[1] == "exit" {
					fmt.Println(`
Command : meeting exit
Function: User can exit the meeting he attended.
instruction: meeting cancel --title/-t [Title]
Args :
	[Title] title of the meeting (unique)
`)
				} else if args[1] == "clear" {
					fmt.Println(`
Command : meeting clear
Function: User can clear all the meetings he created.
instruction: meeting clear
`)
				}
			}
		} else if len(args) == 1 {
			if args[0] == "user" {
				fmt.Println(`
	1. To add a register: 
	
	instruction:	user register -u [UserName] -p [Pass] -e [Email] -t [Phone]
		
		[Username] register's name
		[Pass] register's Password
		[Email] register's email
		[Phone] register's phone-number
	
	2. To remove a register

	instruction:	user delete

	(attention: you can delete your account in the database of Agenda)

	3. To query user

	instruction:	user lookup

	(attention: you can query all user's name who has registed)

	4. To login 

	instruction:	user login -u [UserName] -p [PassWord]

		[Username] register's name
		[Pass] register's Password

	5. To logout

	instruction:	user logout`)
			} else if args[0] == "meeting" {
				fmt.Println(` 
	1.To add Participator of the meeting: 

	instruction:	meeting addUser -p [Participator] -t [Title]
				
		[Participator] Participator's name
		[Title] meeting's name

	(attention:If the Participator cannot attend during the time, add fail.)
	(PS: There can be multiple -p flags at the same time, but each -p flag only aims at one participator.)

	2. To remove a Participator of the meeting

	instruction:	meeting deleteUser -p [Participator] -t [Title]

		[Participator] Participator's name
		[Title] meeting's name

	3. To create a new meeting:

	instruction:	meeting create -t [Title] -p [Participator] -s [StartTime] -e [EndTime]

		[Title] the Title of the meeting
		[Participator] the Participator of the meeting,the Participator can only attend one meeting during one meeting time
		[StartTime] the StartTime of the meeting
		[EndTime] the EndTime of the meeting 

	4. To cancel a meeting:

	instruction:    meeting cancel -t [title]

		[Title] the Title of the meeting
	(PS: Only creator of this meeting can cancel it.)

	5. To query meetings in the specific time for login user:

	instruction:	meeting lookup -s [StartTime] -e [EndTime]

		[StartTime] the StartTime of the meeting
		[EndTime] the EndTime of the meeting

	6. To quit a meeting:

	instruction:	meeting exit -t [title]

		[Title] the Title of the meeting

	(attention: if there is no participators in this meeting,the meeting will be deleted)
				
	7. To clear all meetings:

	instruction:	clear

	`)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
	//rootCmd.SetHelpCommand(helpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
