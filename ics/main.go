package main

import (
	"fmt"
	"strings"
	"time"

	ics "github.com/gunsluo/go-example/ics/ical"
)

func main() {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	cal.SetProductId("fulxble meeting")
	event := cal.AddEvent("sg28dreu7")
	event.AddAttendee("ya.yu@target-energysolutions.com", ics.WithCN("Ray Paseur"), ics.WithRSVP(false))
	event.SetOrganizer("no-reply@target-energysolutions.com")
	event.SetStartAt(time.Now())
	event.SetEndAt(time.Now())
	event.SetSummary("xxx")
	event.SetDescription("https://fluxble-meeting.dev.meeraspace.com/meeting/sg28dreu7")

	fmt.Println("->", cal.Serialize())

	ncal, err := ics.ParseCalendar(strings.NewReader(cal.Serialize()))
	if err != nil {
		panic(err)
	}
	fmt.Println("-->", ncal.Serialize())
}
