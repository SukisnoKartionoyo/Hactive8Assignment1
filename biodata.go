package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type biodata struct {
	Nama      string
	Alamat    string
	Pekertaan string
	Alasan    string
}

func main() {
	var allbio = []biodata{
		{
			Nama:      "sukisno",
			Alamat:    "Jakarta",
			Pekertaan: "Koding",
			Alasan:    "balajar aja ",
		},
		{
			Nama:      "yudistro Wahyu",
			Alamat:    "Jakarta",
			Pekertaan: "Koding",
			Alasan:    "balajar aja ",
		},
		{
			Nama:      "Wisma Eka",
			Alamat:    "Jakarta",
			Pekertaan: "Koding",
			Alasan:    "balajar aja ",
		},
		{
			Nama:      "Wisma Eka",
			Alamat:    "Jakarta",
			Pekertaan: "Koding",
			Alasan:    "balajar aja ",
		},
	}
	args := os.Args[1:2]          // dapatnya string dan array
	idx := args[0]                // arraynya masukin ke var
	ind, err := strconv.Atoi(idx) // saya rubah jadi int

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(ind)
	//fmt.Println(args)
	if ind < 4 {
		fmt.Printf("Nama = %s", allbio[ind].Nama)
		fmt.Println("")
		fmt.Printf("Alamat = %s", allbio[ind].Alamat)
		fmt.Println("")
		fmt.Printf("Pekerjaan = %s", allbio[ind].Pekertaan)
		fmt.Println("")
		fmt.Printf("Alasan = %s", allbio[ind].Alasan)
	} else {
		fmt.Println("Tidak ada data ")
	}
}
