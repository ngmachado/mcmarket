# MCMarket
MagicCardMarket Client (more like SDK...)

If your editor have autocomplete you can navigate to see all the options! :D

This is a sugar version! You can use the more raw version with this project : [mkm](https://github.com/ngmachado/mkm)

How to Use
```
go get github.com/ngmachado/mcmarket
```

Example : 
```
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
