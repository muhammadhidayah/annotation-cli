package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config/cmd"
	pb "github.com/muhammadhidayah/annotation-service/proto/annotation"
)

func parseFile(file string) (*pb.Annotation, error) {
	var annotation *pb.Annotation
	data, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &annotation)

	return annotation, err
}

func main() {
	cmd.Init()

	srv := micro.NewService(
		micro.Name("annotation.cli"),
	)

	client := pb.NewAnnotationService("annotation.service", srv.Client())

	// // insert to database using data dummy
	// file := "annotation.json"

	// annotation, err := parseFile(file)
	// if err != nil {
	// 	log.Fatalf("Could not parse file: %v", err)
	// }

	// r, err := client.CreateAnnotation(context.Background(), annotation)
	// if err != nil {
	// 	log.Fatalf("Could not gree: %v", err)
	// }

	// log.Printf("Created: %t", r.Created)

	// get annotatio by ID
	// var annotation2 pb.Annotation
	// annotation2.AnnotateId = "annotate_000003"

	// r2, err := client.GetAnnotationByID(context.Background(), &annotation2)
	// if err != nil {
	// 	log.Fatalf("Could not greet: %v", err)
	// }

	// log.Println(r2.Annotation)

	// // get annotatio by file ID
	// var annotation2 pb.Annotation
	// annotation2.FileId = "file_00003"

	// r2, err := client.GetAnnotationByFileID(context.Background(), &annotation2)
	// if err != nil {
	// 	log.Fatalf("Could not greet: %v", err)
	// }

	// log.Println(r2.Annotations)

	// update to database using data dummy
	file := "annotation.json"

	annotation, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}
	annotation.FileId = "file_00006"

	r, err := client.UpdateAnnotation(context.Background(), annotation)
	if err != nil {
		log.Fatalf("Could not gree: %v", err)
	}

	log.Printf("Created: %t", r.Updated)

}
