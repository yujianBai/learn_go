package main

import("fmt")

type ISpeaker interface{
	Speak()
}

type SimpleSpraker struct{
	Message string
}

func (speaker * SimpleSpraker)Speak(){
	fmt.Println("I am speaking", speaker.Message)
}

func main(){
	var speaker ISpeaker
	speaker = &SimpleSpraker{"Hell"}
	speaker.Speak()
}
