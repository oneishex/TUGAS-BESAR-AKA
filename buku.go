package main

import (
	"os"

	"github.com/wcharczuk/go-chart/v2"
)

type info struct {
	ID        int
	title     string
	author    string
	publisher string
	halaman   string
	genre     string
	tipe      string
}

type ArrBuku struct {
	n    int
	buku [10000000]info
}

func inputbukudummy(Arr *ArrBuku, n int) {
	//Catatan
	/*Jika terdapat ID yang sama sebelum melakukan inputan dummy maka ID tersebut yang akan keluar dahulu karena
	dalam prosedur ini tidak ada penanganan ketika ID sama*/

	var data1 info
	data1.title = "The Witcher"
	data1.genre = "Action"
	data1.author = "James Hulker"
	data1.halaman = "943"
	data1.publisher = "PT Gramedia"
	data1.tipe = "Novel"

	for i := Arr.n; i < n; i++ {
		Arr.buku[i] = data1
		if i%2 == 1 {
			Arr.buku[i].ID = 3127 + i
		} else {
			Arr.buku[i].ID = 1273 + i
		}
	}
	Arr.n = n
}

func searchIteratif(Arr ArrBuku, id int) info {
	for i := 0; i < Arr.n; i++ {
		if id == Arr.buku[i].ID {
			return Arr.buku[i]
		}
	}
	var notfound info
	notfound.ID = 0
	return notfound
}

func searchRekursif(Arr ArrBuku, id int, i int) info {
	if Arr.buku[i].ID == id {
		return Arr.buku[i]
	} else if i == -1 {
		var notfound info
		notfound.ID = 0
		return notfound
	} else {
		return searchRekursif(Arr, id, i-1)
	}
}

func inputbukumanual(dataBuku *ArrBuku, x int, judul string, penulis string, penerbit string, halaman string, genre string, tipe string) {
	var y info
	y.ID = x
	y.title = judul
	y.author = penulis
	y.genre = genre
	y.publisher = penerbit
	y.halaman = halaman
	y.tipe = tipe
	dataBuku.buku[dataBuku.n] = y
	dataBuku.n++
}

func createChart(inputSizes, runningTimes []float64) {
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "Input Size (Found on (N))",
		},
		YAxis: chart.YAxis{
			Name: "Time (ms)",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name:    "Running Time",
				XValues: inputSizes,
				YValues: runningTimes,
			},
		},
	}

	file, _ := os.Create("runtime_chart.png")
	defer file.Close()
	graph.Render(chart.PNG, file)
}
