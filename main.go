package main

import (
	_ "github.com/joho/godotenv/autoload"

	"encoding/json"
	"fmt"
)

// For debugging
func dump(data interface{}) {
	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Error printing data:", err)
		return
	}
	fmt.Println(string(b))
}




func main(){
	stat := newStats()
	langs := stat.most_used_languages( 14 )
	stat.generate_svg( "elmyrockers_stat.svg", langs )
}