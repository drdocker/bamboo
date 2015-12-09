package haproxy

import (
	"github.com/QubitProducts/bamboo/Godeps/_workspace/src/github.com/samuel/go-zookeeper/zk"
	conf "github.com/QubitProducts/bamboo/configuration"
	"github.com/QubitProducts/bamboo/services/marathon"
	"github.com/QubitProducts/bamboo/services/service"
)

type templateData struct {
	Apps     marathon.AppList
	Services map[string]service.Service
	MaxConn  string
	StatsPort string
	StatsAuth string
	Balance string

}

func GetTemplateData(config *conf.Configuration, conn *zk.Conn) (*templateData, error) {

	apps, err := marathon.FetchApps(config.Marathon, config)

	if err != nil {
		return nil, err
	}

	services, err := service.All(conn, config.Bamboo.Zookeeper)

	if err != nil {
		return nil, err
	}

	maxconn := config.HAProxy.MaxConn
	stats_port := config.HAProxy.StatsPort
	stats_auth := config.HAProxy.StatsAuth
	balance := config.HAProxy.Balance


	return &templateData{apps, services, maxconn, stats_port, stats_auth, balance}, nil
}
