package main


import (
	_ "github.com/joho/godotenv/autoload"
)



func main(){
	stat := stats{}
	langs := stat.most_used_languages( 10 )
	dump( langs )
}