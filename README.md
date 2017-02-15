# MCMarket
MagicCardMarket Client (more like SDK...)

If your editor have autocomplete you can navigate to see all the options! :D

This is a sugar version! You can use the more raw version with this project : [mkm](https://github.com/ngmachado/mkm)

How to Use
```go
go get github.com/ngmachado/mcmarket
```

Example : 

Creating a client to make requests. This client will target sandbox server.
After that we are asking for account information and outputing the result.

If there is a error, we will log and exit.




```go
func main() {
	//Use your keys
	myKeys := mkm.Keys{
		ConsumerKey:       "123456789",
		ConsumerSecret:    "123456789",
		AccessToken:       "123456789",
		AccessTokenSecret: "123456789",
	}
	//Create a http client to make request to market
	client := mcmarket.NewClient(&myKeys,mkm.Sandbox)
	resp,err := client.Account.Information()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)

}
```


To generate the structs i used [json-to-go](https://github.com/mholt/json-to-go).
