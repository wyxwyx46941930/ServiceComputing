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
	. "github.com/wtysos11/goAgenda/entity"
	"github.com/spf13/cobra"
	"errors"
)

const meetingPlace = "meeting.txt"
func userExistCheck(userInfo []User,participants []string) error{
	for _,p := range participants {
		pass := false
		for _,u := range userInfo{
			if u.Username == p{
				pass = true
				break
			}
		}
		if(!pass){
			return errors.New("Participants have illegal participant:"+p)
		}
	}
	return nil
}

//check whether all the users are available and have time to attend this meeting
func userTimeCheck(userInfo []User,meetingInfo []Meeting,startTime AgendaTime, endTime AgendaTime,participants []string) error{
	//first, check all participants are available in userInfo
	if existErr := userExistCheck(userInfo,participants); existErr!=nil{
		return existErr
	}
	//for all meetings, if their userlist have participant, check whether this meeting have conflicts.
	for _,m := range meetingInfo{
		inMeeting := false
		for _,user := range m.UserList{
			for _,p := range participants {
				if user == p{
					inMeeting = true
					break
				}
			}
		}
		if !inMeeting{
			continue
		}

		meetingStartTime,_ := String2Time(m.StartTime)
		meetingEndTime,_ := String2Time(m.EndTime)
		if !(CompareTime(endTime,meetingStartTime)<0 || CompareTime(startTime,meetingEndTime)>0){
			return errors.New("Can't pass user and time test. Your participants may not have time to attend this meeting.") 
		}
	}
	return nil
}
//receive member list that going to check. Current meetings time and the meeting list.
//check whether these users are available
func userAvailableCheck(member []string,meetingInfo []Meeting,sTime AgendaTime, eTime AgendaTime) error{
	if userInfo,userReadingerr := ReadUserFromFile(userPlace);userReadingerr!=nil {
		fmt.Println(userReadingerr)
		return userReadingerr
	} else{
		if userCheckError := userTimeCheck(userInfo,meetingInfo,sTime,eTime,member);userCheckError != nil{
			return userCheckError
		}
	}
	return nil
}

//legal check, don't implement yet
func meetingLegalCheck(meetingInfo []Meeting,startTime string, endTime string,title string ,participants []string) (bool,error){
	if len(title)==0{
		return false,errors.New("Meeting must have a title")
	}

	sTime,tserr := String2Time(startTime)
	eTime,teerr := String2Time(endTime)
	if tserr!=nil {
		return false,tserr
	} else if teerr != nil{
		return false,teerr
	}

	if startTimeErr := TimeLegalCheck(sTime); startTimeErr != nil{
		return false,startTimeErr
	} 
	if endTimeErr := TimeLegalCheck(eTime); endTimeErr != nil{
		return false,endTimeErr
	}

	if CompareTime(sTime,eTime)>=0 {
		return false,errors.New("start time should smaller than end time (equal is not allowed)")
	}

	for _,m := range meetingInfo {
		if m.Title == title{
			return false,errors.New("Repeat title")
		}
	}

	//check participants
	if err:=userAvailableCheck(participants,meetingInfo,sTime,eTime); err!=nil{
		return false,err
	}

	return true,nil
}



// meetingCmd represents the meeting command
var meetingCmd = &cobra.Command{
	Use:   "meeting",
	Short: "operation on meeting",
	Long: ` 
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

	`,
	Run: func(cmd *cobra.Command, args []string) {
		//all meeting operation need a login status
		if login,err:=checklogin(); err!=nil{
			fmt.Println(err)
			return
		} else if !login{
			fmt.Println("Please login first")
			return
		}
		//get login username, for founder of the conference and some operations
		loginUsername,loginErr := getLoginUsername()
		if loginErr !=nil {
			fmt.Println(loginErr)
			return
		}

		//get store meeting data in JSON format
		meetingInfo,meetingReadingerr := ReadMeetingFromFile(meetingPlace)
		if meetingReadingerr!=nil {
			fmt.Println(meetingReadingerr)
			return
		}
		fmt.Println("meeting called")
		startTime,_ := cmd.Flags().GetString("start")
		endTime,_ := cmd.Flags().GetString("end")
		title,_ := cmd.Flags().GetString("title")
		participants,_ := cmd.Flags().GetStringArray("participant")

		/*
		fmt.Println("flags test.")
		fmt.Println(startTime)
		fmt.Println(endTime)
		fmt.Println(title)
		fmt.Println(participants)
*/
		if len(args)>0{
			switch (args[0]){
				case "create":{
					fmt.Println("create")

					if pass,err := meetingLegalCheck(meetingInfo,startTime,endTime,title,participants); err !=nil {
						fmt.Println(err)
						return
					} else if !pass{
						fmt.Println("Meeting create failed")
						return
					}
					meetingInfo = append(meetingInfo,Meeting{loginUsername,startTime,endTime,title,participants})

					WriteMeetingToFile(meetingPlace,meetingInfo)
					fmt.Println("Meeting create success")
				}
				case "addUser":{
					fmt.Println("add user")
					//check. Need title and at least one valiable participants(username correct and have time to attend)
					//check arguments first
					if len(participants) == 0{
						fmt.Println("Add user command needs a user list.")
						return
					} else if len(title)==0 {
						fmt.Println("Add user command needs title")
					}

					//check time valid and operate
					var sTime,eTime AgendaTime
					titleCheck := false
					for i,m := range meetingInfo{
						if m.Title == title{
							titleCheck = true
							sTime,_ = String2Time(m.StartTime)
							eTime,_ = String2Time(m.EndTime)

							//repeat check
							for _,user := range m.UserList{
								for _,p := range participants {
									if user == p{
										fmt.Println("Your participants :"+p+" has already attend this meeting. Add failed.")
										return
									}
								}
							}
							//check time validation here
							if timeErr := userAvailableCheck(participants,meetingInfo,sTime,eTime); timeErr!=nil{
								fmt.Println(timeErr)
								return
							}
							//if valid operate
							meetingInfo[i].UserList = append(meetingInfo[i].UserList,participants...)
							break
						}
					}
					if !titleCheck{
						fmt.Println("Invalid title.")
						return
					}
					
					WriteMeetingToFile(meetingPlace,meetingInfo)
					fmt.Println("Meeting add users success")
				}
				case "deleteUser":{
					fmt.Println("delete user")
					//check. title and participants name
					//check arguments first
					if len(participants) == 0{
						fmt.Println("Add user command needs a user list.")
						return
					} else if len(title)==0 {
						fmt.Println("Add user command needs title")
					}
					//read userPlace file and validate participants
					if userInfo,userReadingerr := ReadUserFromFile(userPlace);userReadingerr!=nil {
						fmt.Println(userReadingerr)
						return
					} else{
						if existErr:=userExistCheck(userInfo,participants); existErr!=nil{
							fmt.Println(existErr)
							return
						}
					}
					//find meeting
					pass := false
					for i := 0; i < len(meetingInfo);i++ {
						meeting := meetingInfo[i]
						if meeting.Title == title{ //find the meeting
							pass = true
							//check whether participants in this meeting
							for _ , p := range participants {
								ok := false
								for _ , user := range meeting.UserList{
									if p == user {
										ok = true
										break
									}
								}
								if !ok {
									fmt.Println("Participants "+p+" not in meeting's userlist.")
									return
								}
							}

							//delete participants from this meeting
							//warning: may have bugs. Not sure
							for j := 0; j < len(meeting.UserList) ; j++ {
								user := meeting.UserList[j]
								for k:=0 ; k < len(participants) ; k++ {
									deleteUser:=participants[k]
									if user == deleteUser{
										if j+1 < len(meetingInfo[i].UserList) {
											meetingInfo[i].UserList = append(meetingInfo[i].UserList[:j],meetingInfo[i].UserList[j+1:]...)
											j--;
										} else {
											meetingInfo[i].UserList = meetingInfo[i].UserList[:j]
										}
										
										if k+1 < len(participants) {
											participants = append(participants[:k],participants[k+1:]...)
											k--;
										} else {
											participants = participants[:k]
										}
										
										break
									}
								}
							}
							//if the delete operation make this meeting empty, clear the meeting
							if len(meetingInfo[i].UserList) == 0 {
								if i+1 < len(meetingInfo){
									meetingInfo = append(meetingInfo[:i],meetingInfo[i+1:]...)
									i--;
								} else {
									meetingInfo = meetingInfo[:i]
								}
								
							}
							break
						}
					}

					if !pass {
						fmt.Println("Meeting delete users failed.")
						return
					}
					WriteMeetingToFile(meetingPlace,meetingInfo)
					fmt.Println("Meeting delete users success")
				}
				case "lookup":{
					//transfer startTime and endTime
					fmt.Println("meeting lookup")
					if len(startTime) == 0 {
						fmt.Println("Meeting lookup command needs a start time.")
					} else if len(endTime) == 0{
						fmt.Println("Meeting lookup command needs a end time.")
					}
					sTime,tserr := String2Time(startTime)
					eTime,teerr := String2Time(endTime)
					if tserr!=nil {
						fmt.Println(tserr)
						return
					} else if teerr != nil{
						fmt.Println(teerr)
						return
					}

					fmt.Println("Meeting lookup:")
					for _,m := range meetingInfo{
						meetingStartTime,_ := String2Time(m.StartTime)
						meetingEndTime,_ := String2Time(m.EndTime)
						if (CompareTime(sTime,meetingStartTime)<0 && CompareTime(meetingStartTime,eTime)<0)||(CompareTime(sTime,meetingEndTime)<0 && CompareTime(meetingEndTime,eTime)<0){
							fmt.Println("Meeting : "+ m.Title+" will start at  "+m.StartTime+" and end at "+m.EndTime+".")
						}
					}
				}
				case "cancel":{
					fmt.Println("meeting cancel")
					if len(title) == 0{
						fmt.Println("Meeting cancel command needs a title.")
					}
					pass := false
					for i := 0 ; i<len(meetingInfo) ; i++{
						meeting := meetingInfo[i]
						if meeting.Title == title && meeting.Creator == loginUsername{
							pass = true
							if(i+1<len(meetingInfo)){
								meetingInfo = append(meetingInfo[:i],meetingInfo[i+1:]...)
								i--
							} else{
								meetingInfo = meetingInfo[:i]
							}
							break
							
						} else if meeting.Title == title && meeting.Creator != loginUsername{
							fmt.Println("You can only cancel the meeting that create by yourself")
							return
						}
					}
					if !pass{
						fmt.Println("Meeting cancel failed")
						return
					}
					WriteMeetingToFile(meetingPlace,meetingInfo)
					fmt.Println("Meeting cancel success")
				}
				case "exit":{
					fmt.Println("meeting exit")
					//find the specific meeting

					//check whether user join
					pass := false
					for i:=0 ; i < len(meetingInfo) ; i++ {
						meeting := meetingInfo[i]
						if meeting.Title == title{
							//check whether the user in the user list
							for j := 0 ; j < len(meeting.UserList) ; j++ {
								user := meeting.UserList[j]
								if user == loginUsername{
									pass = true
									if(j+1 < len(meetingInfo[i].UserList)){
										meetingInfo[i].UserList = append(meetingInfo[i].UserList[:j],meetingInfo[i].UserList[j+1:]...)
									} else{
										meetingInfo[i].UserList = meetingInfo[i].UserList[:j]
									}
									// if UserList is empty, delete the meeting
									if len(meetingInfo[i].UserList) == 0{
										if i+1 < len(meetingInfo){
											meetingInfo = append(meetingInfo[:i],meetingInfo[i+1:]...)
										} else{
											meetingInfo = meetingInfo[:i]
										}
									}

									break
								}
							}

							break
						}
					}

					if !pass {
						fmt.Println("Meeting Exit failed")
						return
					}

					//delete it
					WriteMeetingToFile(meetingPlace,meetingInfo)
					fmt.Println("Meeting Exit success")
				}
				case "clear":{
					fmt.Println("meeting clear")
					for i:=0 ; i < len(meetingInfo) ; i++ {
						if meetingInfo[i].Creator == loginUsername{
							if i+1 < len(meetingInfo){
								meetingInfo = append(meetingInfo[:i],meetingInfo[i+1:]...)
							} else{
								meetingInfo = meetingInfo[:i]
							}
							i--
						}
					}
					WriteMeetingToFile(meetingPlace,meetingInfo)
					fmt.Println("Meeting clear success")
				}
			}
		} else{
			fmt.Println("Need some command.")
		}
	},
}

func init() {
	rootCmd.AddCommand(meetingCmd)
	meetingCmd.Flags().StringP("start","s","","Help message for start time")
	meetingCmd.Flags().StringP("end","e","","Help message for end time")
	meetingCmd.Flags().StringP("title","t","","Help message for meeting title")
	meetingCmd.Flags().StringArrayP("participant","p",[]string{},"Help message for participant")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// meetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// meetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
