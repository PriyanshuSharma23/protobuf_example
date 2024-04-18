package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PriyanshuSharma23/grpc_example/pb/employee"
	"google.golang.org/protobuf/proto"
)

type Employee struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Salary float32 `json:"salary"`
}

func main() {
	// employees := []Employee{
	// 	{ID: 1001, Name: "Priyanshu Sharma", Salary: 1000},
	// 	{ID: 1002, Name: "Ayush Sharma", Salary: 4000},
	// 	{ID: 1003, Name: "Pawan Sharma", Salary: 10000},
	// }

	// enc, err := json.Marshal(employees)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// out, err := os.Create("out.json")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer out.Close()

	// _, err = out.Write(enc)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println("Written successfully!")

	// Using protobuf
	// emps := &pb.Employees{}

	// for _, emp := range employees {
	// 	emps.Employees = append(emps.Employees, &pb.Employee{
	// 		Name:   emp.Name,
	// 		Id:     emp.ID,
	// 		Salary: emp.Salary,
	// 	})
	// }

	// protoData, err := proto.Marshal(emps)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// out, err := os.Create("out_proto")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer out.Close()

	// _, err = out.Write(protoData)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	data, err := os.ReadFile("out_proto")
	if err != nil {
		log.Fatalln(err)
	}

	emps := new(employee.Employees)
	err = proto.Unmarshal(data, emps)

	if err != nil {
		log.Fatalln(err)
	}

	for _, emp := range emps.Employees {
		fmt.Println(emp)
	}
}
