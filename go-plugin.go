package main

import (
	"errors"
	"log"
	"os"

	"github.com/chrisguida/go-cln-plugin/util"
	"github.com/elementsproject/glightning/glightning"
	"github.com/elementsproject/glightning/jrpc2"
)

type Balance struct {
	// Required string `json:"required"`
	Mode string `json:"optional,omitempty"` // Add 'omitempty' to mark optional
}

func (h *Balance) New() interface{} {
	return &Balance{}
}

func (h *Balance) Name() string {
	return "balance"
}
func (h *Balance) Call() (jrpc2.Result, error) {
	funds, err := lightning.ListFunds()
	if err != nil {
		log.Printf("error found: %s\n", err)
		return nil, err
	}

	var result int64
	switch h.Mode {
	case "":
		fallthrough // Treat empty mode as total balance
	case "total":
		result = calculateTotalBalance(funds)
	case "chain":
		result = calculateChainBalance(funds)
	case "channels":
		result = calculateChannelsBalance(funds)
	case "inbound":
		return nil, errors.New("not implemented")
	case "outbound":
		return nil, errors.New("not implemented")
	default:
		return nil, errors.New("invalid balance mode")
	}
	return util.FormatMsat(result), nil
}

func calculateTotalBalance(funds *glightning.FundsResult) int64 {
	var totalBalance int64
	totalBalance += calculateChainBalance(funds)
	totalBalance += calculateChannelsBalance(funds)
	return totalBalance
}

func calculateChainBalance(funds *glightning.FundsResult) int64 {
	var chainBalance int64
	for _, output := range funds.Outputs {
		chainBalance += int64(output.AmountMilliSatoshi)
	}
	return chainBalance
}

func calculateChannelsBalance(funds *glightning.FundsResult) int64 {
	var channelsBalance int64
	for _, channel := range funds.Channels {
		channelsBalance += int64(channel.OurAmountMilliSatoshi)
	}
	return channelsBalance
}

func addCommas(s string) string {
	if len(s) <= 3 {
		return s
	}
	return addCommas(s[:len(s)-3]) + "," + s[len(s)-3:]
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
	rpcBalance := glightning.NewRpcMethod(&Balance{}, "Returns your node's total funds balance")
	rpcBalance.LongDesc = `Returns your node's total funds balance`
	rpcBalance.Category = "utility"
	p.RegisterMethod(rpcBalance)
}
