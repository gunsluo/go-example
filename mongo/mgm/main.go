package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx := context.Background()
	err := mgm.SetDefaultConfig(&mgm.Config{
		CtxTimeout: 5 * time.Second,
	}, "ocr", options.Client().ApplyURI("mongodb://root:password@mongo.infra:27017"))
	if err != nil {
		panic(err)
	}
	conf, client, _, err := mgm.DefaultConfigs()
	fmt.Println("--->", conf)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	err = crud()
	if err != nil {
		panic(err)
	}
}

func crud() error {

	r := newScanRecord("ocrid", "10s", "http://a.b.c")
	coll := mgm.Coll(r)

	r.SetScan(ScanResult{
		License: &DriverLicense{
			LicenceNumber: "0001",
		},
	})

	if err := coll.Create(r); err != nil {
		return err
	}

	r.SetRevised(ScanResult{
		License: &DriverLicense{
			LicenceNumber: "0002",
		},
	})

	if err := coll.Update(r); err != nil {
		return err
	}

	nr := &ScanRecord{}
	ctx := context.Background()
	if err := coll.FindOne(ctx, bson.M{"ocrid": "ocrid"}).Decode(nr); err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return err
		}

		fmt.Println("not found")
	} else {
		fmt.Println("-->", nr.Result.Scan)
	}
	// if err := coll.First(bson.M{"ocrid": "ocrid1"}, nr); err != nil {
	// 	return err
	// }
	// mgm.Err

	resut, err := coll.DeleteOne(ctx, bson.M{"ocrid": "ocrid"})
	return nil
	return coll.Delete(r)
}

// type book struct {
// 	mgm.DefaultModel `bson:",inline"`
// 	Name             string    `json:"name" bson:"name"`
// 	Pages            int       `json:"pages" bson:"pages"`
// 	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
// 	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"`
// }

// func newBook(name string, pages int) *book {
// 	return &book{
// 		Name:      name,
// 		Pages:     pages,
// 		CreatedAt: time.Now().UTC(),
// 		UpdatedAt: time.Now().UTC(),
// 	}
// }

func newScanRecord(id, cost string, urls ...string) *ScanRecord {
	return &ScanRecord{
		OcrId: id,
		Cost:  cost,
		Result: Result{
			ImageUrls: urls,
		},
	}
}

type ScanRecord struct {
	mgm.DefaultModel `bson:",inline"`
	OcrId            string `json:"ocrid" bson:"ocrid"`
	Cost             string `json:"cost" bson:"cost"`
	Result           Result `json:"result" bson:"result"`
}

type Result struct {
	ImageUrls []string   `json:"imageUrls" bson:"imageUrls"`
	Scan      ScanResult `json:"scan" bson:"scan"`
	Revised   ScanResult `json:"revised" bson:"revised"`
}

type ScanResult struct {
	License *DriverLicense `json:"license,omitempty" bson:"license,omitempty"`
	Truck   *Truck         `json:"truck,omitempty" bson:"truck,omitempty"`
}

type DriverLicense struct {
	// Licence Number
	//
	// from driver license front
	LicenceNumber string `json:"licenceNumber" bson:"licenceNumber"`
}

type Truck struct {
	BoardType   string `json:"boardType" bson:"boardType"`
	BoardNumber string `json:"boardNumber" bson:"boardType"`
}

func (sr *ScanRecord) SetScan(r ScanResult) {
	sr.Result.Scan = r
}

func (sr *ScanRecord) SetRevised(r ScanResult) {
	sr.Result.Revised = r
}

func (*ScanRecord) CollectionName() string {
	return "scan_record"
}
