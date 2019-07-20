

package main

import (
"encoding/json"
"fmt"
"log"
)

type movie struct {
	Adult         bool
	Backdrop_path string
	Budget        int
	Genres        []struct {
		Id   int // string
		Name string
	}

}

type AccessToken struct {
	Access_token string
	Expires_in   int64
}


func main() {
	var movieData movie
	str := `
    {
    "adult":false,
    "backdrop_path":"/mWuHbFc7qVmVcpybx3ezhXLj5VO.jpg",
    "belongs_to_collection":null,
    "budget":25000000,
    "genres":
        [
        {
            "id":35,
            "name":"Comedy"
        },
        {
            "id":37,
            "name":"Western"
        }
        ]
    }`





	err := json.Unmarshal([]byte(str), &movieData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(movieData)

	var a AccessToken
	str1  := `{"access_token":"asdfasfd","expires_in":7200}`
	json.Unmarshal([]byte(str1),&a)
	fmt.Println(a)
}