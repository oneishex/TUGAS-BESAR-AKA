package main

import (
	"image"
	"image/color"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var btn1 *widget.Button
var btn2 *widget.Button
var btn3 *widget.Button
var btnShowChart *widget.Button

func main() {
	a := app.NewWithID("com.anydom.com")
	w1 := a.NewWindow("PERPUSTAKAAN BOS")
	var dataBuku ArrBuku
	var runningTimes []float64
	var inputSizes []float64

	img := canvas.NewImageFromFile("d:/Programming/Golang/cobafyne/img/header_bookshelf.jpg")
	img.SetMinSize(fyne.NewSize(400, 80))

	lineX1 := canvas.NewLine(color.White)
	lineX1.StrokeWidth = 2.0

	lineX2 := canvas.NewLine(color.White)
	lineX2.StrokeWidth = 2.0

	judul2 := widget.NewLabel("PERPUSTAKAAN")
	judul2.Alignment = fyne.TextAlignCenter

	judul3 := widget.NewLabel("Selamat Datang di Perpustakaan BOS!")
	judul3.Alignment = fyne.TextAlignCenter

	judul4 := widget.NewLabel("Cari Buku Favorit Anda dengan Mudah dan Cepat!")
	judul4.Alignment = fyne.TextAlignCenter

	judul5 := widget.NewLabel("Cari Buku:")

	judul6 := widget.NewLabel("Masukan banyaknya data buku dummy:")

	con1 := container.NewVBox(
		img, lineX1, judul2, lineX2,
	)

	con2 := container.NewVBox(
		judul3, judul4,
	)

	entry1 := widget.NewEntry()
	entry1.SetPlaceHolder("Masukkan ID buku yang ingin dicari:")

	w1.Resize(fyne.NewSize(800, 580))
	w1.SetContent(
		container.NewVBox(con1, con2),
	)

	btn1 = widget.NewButton("Cari", func() {
		x, err := strconv.Atoi(entry1.Text)

		if err != nil {
			w1.SetContent(
				container.NewVBox(
					con1,
					widget.NewLabel("Masukkan hanya angka!"),
					entry1,
					btn1,
				),
			)
		} else {
			//DISINI ISI BRO
			start := time.Now()
			hasil := searchIteratif(dataBuku, x)
			elapsed := time.Since(start).Seconds() * 1000
			runningTimes = append(runningTimes, elapsed)
			inputSizes = append(inputSizes, float64(x))
			println("Jumlah data untuk grafik:", len(runningTimes))
			//DISINI ISI BRO

			if hasil.ID == 0 {
				w1.SetContent(
					container.NewVBox(
						con1,
						widget.NewLabel("ID tidak ditemukan"),
						entry1,
						btn1,
					),
				)
			} else {
				w1.SetContent(
					container.NewVBox(
						con1,
						widget.NewLabel("ID Buku: "+entry1.Text+"\n"+"Judul Buku: "+hasil.title+"\n"+"Author buku: "+hasil.author+
							"\n"+"Publisher buku: "+hasil.publisher+"\n"+"Genre buku: "+hasil.genre+
							"\n"+"Tipe Buku :"+hasil.tipe+"\n"+"Halaman Buku: "+hasil.halaman),
						entry1,
						btn1,
					),
				)
			}
		}
	})

	entry2 := widget.NewEntry()
	entry2.SetPlaceHolder("Masukkan N!")

	btn2 = widget.NewButton("Masukan", func() {
		x, err := strconv.Atoi(entry2.Text)

		if err != nil {
			w1.SetContent(
				container.NewVBox(
					con1,
					widget.NewLabel("Masukkan hanya angka!"),
					entry2,
					btn2,
				),
			)
		} else {
			if x > len(dataBuku.buku) {
				w1.SetContent(
					container.NewVBox(
						con1,
						widget.NewLabel("Data terlalu besar!"),
						widget.NewLabel("Silahkan masukan banyaknya data kembali:"),
						entry2,
						btn2,
					),
				)
			} else {
				//DISINI ISI BRO
				inputbukudummy(&dataBuku, x)
				//DISINI ISI BRO

				w1.SetContent(
					container.NewVBox(
						con1,
						widget.NewLabel("Berhasil menambahkan "+entry2.Text+" data buku!"),
					),
				)
			}
		}
	})

	MenuItems1 := fyne.NewMenuItem("Book Search", func() {
		w1.SetContent(
			container.NewVBox(con1, judul5, entry1, btn1),
		)
	})

	MenuItems2 := fyne.NewMenuItem("Input Book Auto", func() {
		w1.SetContent(
			container.NewVBox(con1, judul6, entry2, btn2),
		)
	})

	entry3 := widget.NewEntry()
	entry3.SetPlaceHolder("Masukkan ID buku!")

	entry4 := widget.NewEntry()
	entry4.SetPlaceHolder("Masukkan judul buku!")

	entry5 := widget.NewEntry()
	entry5.SetPlaceHolder("Masukkan penulis buku!")

	entry6 := widget.NewEntry()
	entry6.SetPlaceHolder("Masukkan penerbit buku!")

	entry7 := widget.NewEntry()
	entry7.SetPlaceHolder("Masukkan halaman buku!")

	entry8 := widget.NewEntry()
	entry8.SetPlaceHolder("Masukkan genre buku!")

	entry9 := widget.NewEntry()
	entry9.SetPlaceHolder("Masukkan tipe buku!")

	btn3 = widget.NewButton("Masukan", func() {
		x, err := strconv.Atoi(entry3.Text)
		if err != nil {
			w1.SetContent(
				container.NewVBox(con1, widget.NewLabel("DATA BUKU:"), widget.NewLabel("Masukkan ID harus angka!"), entry3, entry4, entry5, entry6, entry7, entry8, entry9, btn3),
			)
		} else {
			y := searchIteratif(dataBuku, x)
			if y.ID == 0 {
				if dataBuku.n+1 > len(dataBuku.buku) {
					w1.SetContent(
						container.NewVBox(con1, widget.NewLabel("Data Buku Sudah Penuh!")),
					)
				} else {
					inputbukumanual(&dataBuku, x, entry4.Text, entry5.Text, entry6.Text, entry7.Text, entry8.Text, entry9.Text)
					w1.SetContent(
						container.NewVBox(con1, widget.NewLabel("Berhasil Memasukkan data!")),
					)
				}
			} else {
				w1.SetContent(
					container.NewVBox(con1, widget.NewLabel("ID Sudah Ada Silahkan Ganti!"), widget.NewLabel("DATA BUKU:"), entry3, entry4, entry5, entry6, entry7, entry8, entry9, btn3),
				)
			}
		}
	})

	MenuItems3 := fyne.NewMenuItem("Input Book Manual", func() {
		w1.SetContent(
			container.NewVBox(con1, widget.NewLabel("DATA BUKU:"), entry3, entry4, entry5, entry6, entry7, entry8, entry9, btn3),
		)
	})

	btnShowChart := widget.NewButton("Show Chart", func() {
		// Pastikan data runningTimes sudah diisi
		if len(runningTimes) == 0 {
			w1.SetContent(
				container.NewVBox(
					widget.NewLabel("Tidak ada data untuk grafik. Jalankan pencarian terlebih dahulu!"),
				),
			)
			return
		}

		// Panggil fungsi untuk membuat grafik
		createChart(inputSizes, runningTimes)

		// Buka file grafik yang baru dibuat
		file, err := os.Open("runtime_chart.png")
		if err != nil {
			w1.SetContent(
				container.NewVBox(
					widget.NewLabel("Gagal membuka file grafik: " + err.Error()),
				),
			)
			return
		}
		defer file.Close()

		// Decode gambar dari file
		img, _, err := image.Decode(file)
		if err != nil {
			w1.SetContent(
				container.NewVBox(
					widget.NewLabel("Gagal membaca gambar grafik: " + err.Error()),
				),
			)
			return
		}

		// Tampilkan gambar di canvas
		imageCanvas := canvas.NewImageFromImage(img)
		imageCanvas.SetMinSize(fyne.NewSize(400, 250))
		w1.SetContent(container.NewVBox(imageCanvas))
	})

	MenuItems4 := fyne.NewMenuItem("Show Graph", func() {
		w1.SetContent(
			container.NewVBox(con1, widget.NewLabel("Graph:"), btnShowChart),
		)
	})

	menu := fyne.NewMenu("Main", MenuItems1, MenuItems2, MenuItems3, MenuItems4)

	main := fyne.NewMainMenu(menu)

	w1.SetMainMenu(main)
	w1.ShowAndRun()
}
