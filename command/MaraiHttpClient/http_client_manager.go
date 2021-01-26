package MaraiHttpClient

import (
	"fmt"
)


func managerRun() maraiHttpClient{

	var client maraiHttpClient

	fmt.Println("this is template main")
	authKey := "INITKEYpHQzpWXk"
	adminQQNumber := "31792690"

	client = newMaraiClient(authKey,adminQQNumber)

	result, _ := client.getAbout()

	fmt.Println(result.String())


	_ = client.verifySession()

	return client


}
