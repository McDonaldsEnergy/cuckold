//	Don't judge, I don't code in GO it's just this specific project I found would work the best if
//	main API calling was in GO instead of Python or C#, etc.

package main

import (
	"fmt"
	"net/http"
)

func main() {
	Call("Authorization token here")
}

func Call(uri string) {
	req, err := http.NewRequest("PUT", fmt.Sprintf("https://0.0.0.0/api/v2/cuckold/call/%s", uri), nil)

	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode == 204 {
		fmt.Printf("Called successfully, check panel.")
		return
	} else {
		fmt.Printf("Error calling: %s", resp.Status)
		return
	}

}
