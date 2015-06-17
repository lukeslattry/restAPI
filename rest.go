package main

import  (
	"fmt"
	"encoding/json"
	"net/http"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit Fruits
	Veg Vegetables
}

type Fruits map[string]string
type Vegetables map[string]string

func serveRest( w http.ResponseWriter, r *http.Request ) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf( w, string( response ) )
}

func main() {
	http.HandleFunc( "/", serveRest )
	http.ListenAndServe( "localhost:1337", nil )
}

func getJsonResponse() ( []byte, error ) {
	fruits := make( map[string]string )
	fruits["Apples"] = "25"
	fruits["Oranges"] = "11"
	
	vegetables := make( map[string]string )
	vegetables["Carrots"] = "21"
	vegetables["Peppers"] = "0"

	d := Data{ fruits, vegetables }
	p := Payload{ d }

	return json.MarshalIndent( p, "", "  " )
}
