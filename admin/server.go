package admin

import (
	"fmt"
	"github.com/jchavannes/jgo/jerr"
	"github.com/jchavannes/jgo/jlog"
	"github.com/memocash/server/node"
	"github.com/memocash/server/ref/config"
	"net/http"
)

const (
	UrlIndex              = "/"
	UrlNodeGetAddrs       = "/node/get_addrs"
	UrlNodeConnectDefault = "/node/connect_default"
)

type Server struct {
	Nodes *node.Group
}

func (s *Server) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc(UrlIndex, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Memo Admin 0.1")
	})
	mux.HandleFunc(UrlNodeGetAddrs, func(w http.ResponseWriter, r *http.Request) {
		jlog.Log("Node get addrs request")
		for _, serverNode := range s.Nodes.Nodes {
			serverNode.GetAddr()
		}
	})
	mux.HandleFunc(UrlNodeConnectDefault, func(w http.ResponseWriter, r *http.Request) {
		jlog.Log("Node connect default")
		s.Nodes.AddDefaultNode()
	})
	server := http.Server{
		Addr:    config.GetHost(config.GetAdminPort()),
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		return jerr.Get("error listening and serving admin server", err)
	}
	return nil
}

func NewServer(group *node.Group) *Server {
	return &Server{
		Nodes: group,
	}
}
