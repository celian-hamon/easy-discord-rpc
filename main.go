package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/hugolgst/rich-go/client"
)

type rpcItem struct {
	RpcAppId    string `json:"rpcAppId"`
	RpcState    string `json:"rpcState"`
	RpcDetails  string `json:"rpcDetails"`
	RpcLargeImg string `json:"rpcLargeImg"`
	RpcLargeTxt string `json:"rpcLargeTxt"`
	RpcSmallImg string `json:"rpcSmallImg"`
	RpcSmallTxt string `json:"rpcSmallTxt"`
	RpcPartyID  string `json:"rpcPartyID"`
	RpcPartyNum string `json:"rpcPartyNum"`
	RpcPartyMax string `json:"rpcPartyMax"`
	RpcStart    string `json:"rpcStart"`
	RpcEnd      string `json:"rpcEnd"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("webapp.html"))
	tmpl.Execute(w, nil)
}

// constructor function
func (rpcItem *rpcItem) fill_defaults() {
	if rpcItem.RpcAppId == "" {
		rpcItem.RpcAppId = "935173881256345751"
	}
	if rpcItem.RpcState == "" {
		rpcItem.RpcState = "In Progress"
	}
	if rpcItem.RpcDetails == "" {
		rpcItem.RpcDetails = "A game of chess"
	}
	if rpcItem.RpcLargeImg == "" {
		rpcItem.RpcLargeImg = "https://i.imgur.com/BJjw1VN.png"
	}
	if rpcItem.RpcLargeTxt == "" {
		rpcItem.RpcLargeTxt = "Chess"
	}
	if rpcItem.RpcSmallImg == "" {
		rpcItem.RpcSmallImg = "https://i.imgur.com/BJjw1VN.png"
	}
	if rpcItem.RpcSmallTxt == "" {
		rpcItem.RpcSmallTxt = "Chess"
	}
	if rpcItem.RpcPartyID == "" {
		rpcItem.RpcPartyID = "party-id"
	}
	if rpcItem.RpcPartyNum == "" {
		rpcItem.RpcPartyNum = "1"
	}
	if rpcItem.RpcPartyMax == "" {
		rpcItem.RpcPartyMax = "2"
	}
}

func setRPC(w http.ResponseWriter, r *http.Request) {
	var newRPC rpcItem

	err := json.NewDecoder(r.Body).Decode(&newRPC)

	newRPC.fill_defaults()

	bodyBytes, _ := ioutil.ReadAll(r.Body)

	//print result
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	if err != nil {
		panic(err)
	}

	go rpc(newRPC)

	fmt.Fprintf(w, "RPC set")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)

	router.HandleFunc("/setRPC", setRPC).Methods("POST")

	http.ListenAndServe(":8080", router)
}

func rpc(rpcItem rpcItem) {
	err := client.Login(rpcItem.RpcAppId)
	if err != nil {
		panic(err)
	}
	party, _ := strconv.Atoi(rpcItem.RpcPartyNum)
	max, _ := strconv.Atoi(rpcItem.RpcPartyMax)

	err = client.SetActivity(client.Activity{
		State:      rpcItem.RpcState,
		Details:    rpcItem.RpcDetails,
		LargeImage: rpcItem.RpcLargeImg,
		LargeText:  rpcItem.RpcLargeTxt,
		SmallImage: rpcItem.RpcSmallImg,
		SmallText:  rpcItem.RpcSmallTxt,
		Party: &client.Party{
			ID:         rpcItem.RpcPartyID,
			Players:    party,
			MaxPlayers: max,
		},
	})

	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 5)
}
