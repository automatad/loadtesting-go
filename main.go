package main

import (
	"fmt"
	"loadingtesting-go/models/beeswax/bid"
	"loadingtesting-go/models/beeswax/openrtb"
	"log"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
	"google.golang.org/protobuf/proto"
)

func main() {

	var lineitemId uint64 = 612
	var campaignId uint64 = 185
	var id string = "175"
	var ua string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36"
	var domain string = "patch.com"
	var page string = "https://www.colgate.com/en-in/oral-health/implants/what-are-dental-implants"
	var segmentId string = "domain_list"
	reqModel := &bid.BidAgentRequest{
		BidRequest: &openrtb.BidRequest{
			Id:     &id,
			Device: &openrtb.BidRequest_Device{Ua: &ua},
			Site:   &openrtb.BidRequest_Site{Domain: &domain, Page: &page},
			Ext:    &openrtb.BidRequestExtensions{AugmentorData: []*openrtb.BidRequestExtensions_AugmentorData{{Segment: []*openrtb.BidRequestExtensions_AugmentorData_Segment{{Id: &segmentId}}}}}},
		Adcandidates: []*bid.Adcandidate{{LineItemId: &lineitemId, CreativeIds: []uint64{1182, 1183, 1184}, CampaignId: &campaignId}},
	}
	data, err := proto.Marshal(reqModel)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	rate := vegeta.Rate{Freq: 80, Per: time.Second}
	duration := 60 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    "http://bw1.dev.90d.io/beeswax",
		Body:   data,
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
	fmt.Printf("95th percentile: %s\n", metrics.Latencies.P95)
	fmt.Printf("50th percentile: %s\n", metrics.Latencies.P50)
	fmt.Printf("Mean percentile: %s\n", metrics.Latencies.Mean)
}
