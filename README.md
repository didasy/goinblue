# Goinblue
---------------------------------
### A work in progress package for golang to send email and SMS through sendinblue.com

#### Usage
Example:

```
package main

import (
	"github.com/JesusIslam/goinblue"
	"fmt"
)

func main() {
	myApiKey := "thisisyourapikey"

	email := &goinblue.Email{
		To: map[string]string{
			"to@example.com": "Mr. To",
		},
		Subject: "Test",
		From: []string{
			"from@example.com", "From",
		},
		Text: "This is just a test.",
	}

	client := goinblue.NewClient(myApiKey)
	res, err := client.SendEmail(email)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```
