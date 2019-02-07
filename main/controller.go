package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

/** JSONデコード用に構造体定義 */
type NodeId struct {
	TrustScore       uint64    `json:"trust_score"`
	PrivateTrust     uint64 `json:"private_trust"`
	GroupTrust 		 uint64 `json:"group_trust"`
}

type ExchangeInformation struct {
	Id       		 string    `json:"id"`
	Location 		struct{
		TrustScore       uint64    `json:"trust_score"`
		PrivateTrust     uint64 `json:"private_trust"`
		GroupTrust 		 uint64 `json:"group_trust"`
	}`json:"location"`
	Company 		struct{
		TrustScore       uint64    `json:"trust_score"`
		PrivateTrust     uint64 `json:"private_trust"`
		GroupTrust 		 uint64 `json:"group_trust"`
	}`json:"company"`
}

func main(){
	iter := 1000000
	status := "if"
	//json読み込み
	var node_id, exchange = read_json()

	switch status {
	case "if": performance_if(node_id, exchange, iter)
	case "bit": performance_bit(node_id, exchange, iter)
	case "routine": performance_routine(node_id, exchange, iter)
	case "all":
		performance_if(node_id, exchange, iter)
		performance_bit(node_id, exchange, iter)
		performance_routine(node_id, exchange, iter)
	}
}

func performance_routine(node_id *NodeId , exchange *ExchangeInformation, iter int)  {
	t_start := time.Now()
	for i := 0; i < iter; i++{
		routine_cal(node_id, exchange)
	}
	t_finish := time.Now()
	duration_routine := t_finish.Sub(t_start)
	fmt.Println("routine: ",duration_routine)
}

func performance_bit(node_id *NodeId , exchange *ExchangeInformation, iter int)  {
	t_start := time.Now()
	for i := 0; i < iter; i++{
		 bit_cal(node_id, exchange)
	}
	t_finish := time.Now()
	duration_bit := t_finish.Sub(t_start)
	fmt.Println("bit: ",duration_bit)
}

func performance_if(node_id *NodeId , exchange *ExchangeInformation, iter int)  {
	t_start := time.Now()
	for i := 0; i < iter; i++{
		if_cal(node_id, exchange)
	}
	t_finish := time.Now()
	duration_if := t_finish.Sub(t_start)
	fmt.Println("if: ",duration_if)
}

func read_json() (*NodeId, *ExchangeInformation){
	// JSONファイル読み込み
	bytes_NI, err := ioutil.ReadFile("routing_node_id.json")
	if err != nil {
		log.Fatal(err)
	}
	bytes_EI, err := ioutil.ReadFile("taxi_exchange_information.json")
	if err != nil {
		log.Fatal(err)
	}
	// JSONデコード
	var node_id = new(NodeId)
	if err := json.Unmarshal(bytes_NI, &node_id); err != nil {
		log.Fatal(err)
	}
	var exchange = new(ExchangeInformation)
	if err := json.Unmarshal(bytes_EI, &exchange); err != nil {
		log.Fatal(err)
	}

	return node_id, exchange
}