package internal

import (
	"net"
	"sync"

	"google.golang.org/protobuf/proto"
)

type Options struct {
	Bind        string
	HttpService string
}

type Server struct {
	Options
	ID                   int64
	listener             net.Listener
	serviceInfoMapLocker sync.RWMutex
	serviceInfoMap       map[string]*serviceChannelSet
	serviceChannelMap    map[string]*ServiceChannel
}

type serviceChannelSet struct {
	callIndex     int64
	ClientStreams bool              `json:"client_streams"`
	ServerStreams bool              `json:"server_streams"`
	ServiceInfo   []*ServiceChannel `json:"service_info"`
}

type rpcReq struct {
	reqID      []int64
	reqChannel *ServiceChannel
}

type streamOpenReq struct {
	streamID   []int64
	reqChannel *ServiceChannel
}

type StreamChannel struct {
	req      bool
	method   string
	StreamID int64
	close    func()
	channel  *ServiceChannel
}

type ServiceChannel struct {
	Addr             string                     `json:"addr"`
	ChannelID        int64                      `json:"channel_id"`
	ServiceRegistrar *proto.ServiceRegistrarReq `json:"service"`
	mutext           sync.RWMutex
	closeCh          chan interface{}
	messages         chan *proto.Message
	rpcReqMap        map[int64]*rpcReq
	streamOpenReqMap map[int64]*streamOpenReq
	streamChannelMap map[int64]*StreamChannel
	onCC             map[int64]func()
}

func New(options Options) *Server {
	return &Server{
		Options:              options,
		serviceInfoMapLocker: sync.RWMutex{},
		serviceInfoMap:       map[string]*serviceChannelSet{},
		serviceChannelMap:    map[string]*ServiceChannel{},
	}
}

func (c *ServiceChannel) handleRpcReq(channel *ServiceChannel, req *proto.RpcReq) *proto.Message {

}
