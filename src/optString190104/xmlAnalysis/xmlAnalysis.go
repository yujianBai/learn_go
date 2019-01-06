package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/xml"
)

func getAttributeValue(attr []xml.Attr, name string)string{
	for _, a := range attr{
		if a.Name.Local == name{
			return a.Value
		}
	}
	return ""
}

func main(){
	
	content, err := ioutil.ReadFile("test.csproj")
	decoder := xml.NewDecoder(bytes.NewBuffer(content))

	var t xml.Token
	flaghe := false
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token(){
		// fmt.Printf("token t:%v\n", t)
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			if flaghe{
				if name == "hehehe"{
					fmt.Println(getAttributeValue(token.Attr, "hehehe"))	
				}
			}else{
				if name == "hehehe"{
					flaghe = true 
				}else{
					flaghe = false
				}
			}
			// fmt.Printf("name:%s\n", name)
		case xml.EndElement:	
			if flaghe{
				if token.Name.Local == "hehehe"{
					flaghe = false
				}
			}
		}
	}

}