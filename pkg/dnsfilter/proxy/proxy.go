package proxy

import (
	"errors"

	"github.com/miekg/dns"
	"github.com/mustafanafizdurukan/pomodoro/pkg/dnsfilter/cache"
)

var (
	errorNoQuestion     = errors.New("there is no question")
	errorNoResponse     = errors.New("there is no response")
	errorBlockedService = errors.New("blocked service dns request")
)

type DNS struct {
	Cache            *cache.Cache
	IsBlockingActive bool
	blockedServices  []string
}

func New() *DNS {
	return &DNS{
		Cache: cache.New(),
	}
}

func (d *DNS) Start() {
	server := &dns.Server{Addr: ":53", Net: "udp"}
	go server.ListenAndServe()

	dns.HandleFunc(".", func(w dns.ResponseWriter, req *dns.Msg) {
		resp, err := d.getResponse(req)
		if err != nil {
			// fmt.Println(err, req.Question[0].Name)
			return
		}

		w.WriteMsg(resp)
	})
}

func (d *DNS) getResponse(req *dns.Msg) (*dns.Msg, error) {
	if len(req.Question) < 1 {
		return nil, errorNoQuestion
	}

	dnsReq := req.Question[0]
	if d.IsBlockingActive {
		if d.isBlockedService(dnsReq.Name) {
			return nil, errorBlockedService
		}
	}

	answer, err := d.handleAllTypes(req)
	if err != nil {
		return nil, err
	}
	return answer, nil

	// switch dnsReq.Qtype {
	// case dns.TypeA:
	// 	answer, err := d.handleTypeA(req)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return answer, nil

	// default:
	// 	answer, err := d.handleOtherTypes(req)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return answer, nil
	// }
}

func (d *DNS) handleAllTypes(req *dns.Msg) (*dns.Msg, error) {
	resp, err := sendDNSRequest(req)
	if err != nil {
		return nil, err
	}

	if len(resp.Answer) < 1 {
		return nil, errorNoResponse
	}

	return resp, nil
}

func (d *DNS) handleTypeA(req *dns.Msg) (*dns.Msg, error) {
	resp := new(dns.Msg)

	dnsQst := req.Question[0]
	e, ok := d.Cache.Get(dnsQst.Name)
	if ok {
		resp.SetReply(req)

		resp.Answer = append(resp.Answer, *e.Answer)
		return resp, nil
	}

	resp, err := sendDNSRequest(req)
	if err != nil {
		return nil, err
	}

	if len(resp.Answer) > 0 {
		d.Cache.Set(dnsQst.Name, &resp.Answer[0])
		return resp, nil
	}

	return nil, errorNoResponse
}

func (d *DNS) handleOtherTypes(req *dns.Msg) (*dns.Msg, error) {
	msg, err := sendDNSRequest(req)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func sendDNSRequest(req *dns.Msg) (*dns.Msg, error) {
	dnsClient := new(dns.Client)
	dnsClient.Net = "udp"

	resp, _, err := dnsClient.Exchange(req, "8.8.8.8:53")
	if err != nil {
		return nil, err
	}

	return resp, nil
}
