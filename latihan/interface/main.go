package main

import (
	"fmt"
	"sync"

	"github.com/ibrahimker/hacktiv/latihan/interface/service"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)
	names := []string{"budi", "cahya", "tono", "andi", "fika", "taslim", "burok"}
	var wg1 sync.WaitGroup
	wg1.Add(len(names))
	for _, n := range names {
		go func(name string) {
			res := userSvc.Register(&service.User{Nama: name})
			fmt.Println(res)
			wg1.Done()
		}(n)
	}
	wg1.Wait()

	resGet := userSvc.GetUser()
	fmt.Println("-----------Hasil get user-------------")
	var wg sync.WaitGroup
	wg.Add(len(resGet))
	for _, v := range resGet {
		go cetakNama(&wg, v.Nama)
		//fmt.Println(v.Nama)
	}
	wg.Wait()
	//time.Sleep(1 * time.Second)
	//k2 := NewKotak(3)
	//k2.Luas()

}

func cetakNama(wg *sync.WaitGroup, nama string) {
	fmt.Println(nama)
	wg.Done()
}

//type kotak struct {
//	sisi int
//}
//
//type BangunanIface interface {
//	Luas()
//}
//
//func NewKotak(sisi int) BangunanIface {
//	return &kotak{sisi: sisi}
//}
//
//func (k *kotak) Luas() {
//	fmt.Println(k.sisi * k.sisi)
//}
