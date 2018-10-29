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
	"io/ioutil"
	"os"
	. "github.com/wtysos11/goAgenda/entity"
	"github.com/spf13/cobra"
	"errors"
)

const userPlace = "user.txt"
const cachePlace = "cache.txt"

//only check username repeat now
func userLegalCheck(userInfo []User,username string, password string,email string ,telphone string) (bool,error){
	for _,user := range userInfo {
		if user.Username == username{
			return false,errors.New("Repeat Username")
		}
	}
	if len(password) == 0{
		return false,errors.New("Must have a password")
	} else if len(email)==0 {
		return false,errors.New("Must have an email")
	} else if len(telphone)==0 {
		return false,errors.New("Must have a telphone")
	}
	return true,nil
}

func checklogin() (bool,error){
	b,err := ioutil.ReadFile(cachePlace)
	if err!=nil {
		return false,err
	}
	str := string(b)

	if str == "logout"{
		return false,nil
	} else{
		return true,nil
	}
}

func getLoginUsername() (string,error){
	b,err := ioutil.ReadFile(cachePlace)
	if err!=nil {
		return "",err
	}
	return string(b),nil
}

func userLogin(username string) error{
	return ioutil.WriteFile(cachePlace,[]byte(username),os.ModeAppend)
}

func userLogout() error{
	return ioutil.WriteFile(cachePlace,[]byte("logout"),os.ModeAppend)
}


// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "help a user to regist",
	Long: `
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

	instruction:	user logout`,
	Run: func(cmd *cobra.Command, args []string) {
		//reading file from store place
		userInfo,userReadingerr := ReadUserFromFile(userPlace)
		if userReadingerr!=nil {
			fmt.Println(userReadingerr)
			return
		}

		//get flags
		username, _ := cmd.Flags().GetString("user")
		password, _:= cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		telphone,_ := cmd.Flags().GetString("telphone")

		if len(args)>0 {
			switch (args[0]){
				case "register":{
					fmt.Println("register")
					//legal check for username(unique),password,email,telphone
					if pass,err := userLegalCheck(userInfo,username,password,email,telphone); err!=nil{
						fmt.Println(err)
						return
					}else if !pass{
						fmt.Println("Register Failed")
						return
					}

					//if pass legal check, add it to userFile
					userInfo = append(userInfo,User{username,password,email,telphone})
					//store the user file into userPlace
					WriteUserToFile(userPlace,userInfo)
					fmt.Println("User register success")
				}
				case "login":{
					fmt.Println("user login")
					//check from cache whether the status is login.
					if check,error := checklogin(); error!=nil{
						fmt.Println(error)
						return
					} else if check {
						fmt.Println("Already Login")
						return
					}
					//validate username and password
					if len(username) == 0 || len(password) == 0 {
						fmt.Println("Must have a username and a password")
						return
					}

					pass := false
					for _,user := range userInfo{
						if user.Username == username && user.Password == password{
							userLogin(user.Username)
							pass = true
							break
						}
					}
					//if no pass, report
					if !pass {
						fmt.Println("login failed")
						return
					}
					
					fmt.Println("Login success. Welcome!")
				}
				case "logout":{
					fmt.Println("user logout")
					//if status is login, make the status logout

					pass,err := checklogin()
					if err!=nil{
						fmt.Println(err)
						return
					} else if !pass {
						fmt.Println("Please login first.")
						return
					}

					userLogout()
					fmt.Println("Logout success")
					
				}
				case "lookup":{
					fmt.Println("user lookup")
					//check the status (login)
					pass,err := checklogin()
					if err!=nil{
						fmt.Println(err)
						return
					} else if !pass {
						fmt.Println("Please login first.")
						return
					}
					//if pass validation, give all info from all users
					fmt.Println("Here is users' info:")
					for _,user := range userInfo{
						fmt.Println(user.Username,user.Email,user.Telphone)
					}
				}
				case "delete":{
					fmt.Println("user delete")
					//check status login
					pass,err := checklogin()
					if err!=nil{
						fmt.Println(err)
						return
					} else if !pass {
						fmt.Println("Please login first.")
						return
					}
					loginUsername,loginErr := getLoginUsername()
					if loginErr!=nil{
						fmt.Println(loginErr)
						return
					}
					//if pass, delete this user and logout
					for i,user := range userInfo{
						if loginUsername == user.Username{
							if i+1 < len(userInfo){
								userInfo = append(userInfo[:i],userInfo[i+1:]...)
							} else{
								userInfo = userInfo[:i]
							}
							
							break
						}
					}
					//update the userPlace
					WriteUserToFile(userPlace,userInfo)
					userLogout()
				}
				default:{
					fmt.Println("Unknown command")
				}
			}
		}
		
	},
}


func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.Flags().StringP("user","u","","Help message for username")
	userCmd.Flags().StringP("password","p","","Help message for password")
	userCmd.Flags().StringP("email","e","","Help message for email")
	userCmd.Flags().StringP("telphone","t","","Help message for telphone")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
