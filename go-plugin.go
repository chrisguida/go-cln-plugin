package main

import (
	"fmt"
	"log"
	"os"

	"github.com/elementsproject/glightning/glightning"
	"github.com/elementsproject/glightning/jrpc2"
)

type Balance struct{}

func (h *Balance) New() interface{} {
	return &Balance{}
}

func (h *Balance) Name() string {
	return "balance"
}

func (h *Balance) Call() (jrpc2.Result, error) {
	log.Printf("calling listfunds\n")
	funds, err := lightning.ListFunds()
	if err != nil {
		log.Printf("error found: %s\n", err)
		return nil, err
	}
	log.Printf("converting to json\n")

	var totalAmountMsat int64
	for _, output := range funds.Outputs {
		totalAmountMsat += int64(output.AmountMilliSatoshi)
	}

	for _, channel := range funds.Channels {
		totalAmountMsat += int64(channel.AmountMilliSatoshi)
	}

	return fmt.Sprint(totalAmountMsat), nil
}

var lightning *glightning.Lightning
var plugin *glightning.Plugin

func main() {
	plugin = glightning.NewPlugin(onInit)
	lightning = glightning.NewLightning()

	registerMethods(plugin)
	err := plugin.Start(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

func onInit(plugin *glightning.Plugin, options map[string]glightning.Option, config *glightning.Config) {
	log.Printf("successfully init'd! %s\n", config.RpcFile)

	lightning.StartUp(config.RpcFile, config.LightningDir)
	channels, _ := lightning.ListChannels()
	log.Printf("You know about %d channels", len(channels))
	log.Printf("Is this initial node startup? %v\n", config.Startup)
}

func registerMethods(p *glightning.Plugin) {
	rpcBalance := glightning.NewRpcMethod(&Balance{}, "Say hello!")
	rpcBalance.LongDesc = `Returns your node's total funds balance`
	rpcBalance.Category = "utility"
	p.RegisterMethod(rpcBalance)
}
