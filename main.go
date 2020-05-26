package main

import (
	"flag"
	"github.com/FanYoung/ping-pong-demo/pkg"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
)

var (
	configFile string
	port int
)

func init() {
	flag.StringVar(&configFile, "config", "", "config name")
	flag.IntVar(&port, "port", 0, "server port")
	flag.Parse()

	log.SetFlags(log.Ldate|log.Lshortfile)
}

func main() {
	bs, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("read config failed, error %+v\n", err)
		return
	}

	config := pkg.ServerConfig{}
	err = yaml.Unmarshal(bs, &config)
	if err != nil {
		log.Printf("parse config failed, error %+v\n", err)
		return
	}

	if len(config.ClusterHosts) == 0 || len(config.ClusterHosts) / 2 == 0 {
		log.Printf("the number of hosts should be odd")
		return
	}


}