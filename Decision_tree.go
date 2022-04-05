package main
import "fmt"

type Store struct {
	nama string
	harga int
}

const N = 1000000
 
type arrayType[N] Store

var data arrayType

func menu (data *arrayType) {
	
	data[0].nama="StrawberryCrispy"
	data[0].harga= 3500
	data[1].nama="ChocoOreo"
	data[1].harga= 4000
	data[2].nama="BiskuitVanilla"
	data[2].harga= 4000
	data[3].nama="GummyBear"
	data[3].harga= 3000
	data[4].nama="MiloSpesial"
	data[4].harga= 3500
	data[5].nama="Mystery"
	data[5].harga= 4000
	
}

func main() {
	var pilihan int
	var total, total2, total3 int
	var porsi int
	var uang int
	var kembalian int
	var pilih string
	var tot int
	var bantu int
	var jum int
	var nhar int
	var totruh1, totruh2 int
	var jawaban int
	var cek bool
	var jawab string
	menu(&data)
	jum = 5
	for !cek{
		fmt.Println ("Selamat Datang di NL Ice Store")
		fmt.Println ("Silahkan pilih : ")
		fmt.Println ("1.Beli")
		fmt.Println ("2.Tambah Menu")
		fmt.Println ("3.Keluar")
		fmt.Println (" ")
		fmt.Println ("Masukan pilihan")
		fmt.Scan(&pilihan)
		if pilihan==1{
			fmt.Println ("//////////////////// Menu NL Ice Store///////////////////////")
			fmt.Println ("/////////////////////////////////////////////////////////////")
			Menu(&data,jum)
			fmt.Println("Masukkan menu pilihan : ")
			fmt.Scan(&pilih)
			bantu = Search(data,pilih)
				
			fmt.Println("Mau Berapa Porsi ?")
				fmt.Scan(&porsi)
				total = (data[bantu].harga) * porsi 
				tot	=tot + total
					
				fmt.Println(tot)
				
			fmt.Println("Mau Menambah Menu Lainnya ?")
			fmt.Println("tekan 1 untuk ya, 2 untuk tidak")
			fmt.Scan(&jawaban)
			if jawaban==1 {
				fmt.Println("Masukan Menu Pilihan : ")
				fmt.Scan(&pilih)
				bantu = Search(data,pilih)
				
				fmt.Println("Mau Berapa Porsi ?")
				fmt.Scan(&porsi)
				total2 = (data[bantu].harga) * porsi 
				totruh1 = tot + total2
				fmt.Println(totruh1)
					
				fmt.Println("Mau Menambah Menu ?")
				fmt.Println("tekan 1 untuk ya, 2 untuk tidak")
				fmt.Scan(&jawaban)
					
				if jawaban==1 {
					fmt.Println("Masukan Menu Pilihan : ")
					fmt.Scan(&pilih)
					bantu = Search(data,pilih)
				
					fmt.Println("Mau Berapa Porsi ?")
					fmt.Scan(&porsi)
					total3 = (data[bantu].harga) * porsi 
					totruh2 = totruh1 + total3
					fmt.Println(totruh2)
					
					fmt.Println("Masukan Jumlah Uang Anda : ")
					fmt.Scan(&uang)
					if uang < totruh2 {
						fmt.Println("Uang Anda Kurang, Mohon Membayar Sesuai Pesanan")
					} else{
						kembalian = uang - totruh2
						fmt.Println("Kembalian Anda : ",kembalian)
						fmt.Println("Terima Kasih Telah Membeli")
					}
				
				} else if jawaban==2 {
					
					fmt.Println("Masukan Jumlah Uang Anda : ")
					fmt.Scan(&uang)
					if uang < totruh1 {
						fmt.Println("Mohon Maaf Uang Anda Kurang, Mohon Membayar Sesuai Pesanan")
					} else{
						kembalian = uang - totruh1
						fmt.Println("Kembalian Anda : ",kembalian)
						fmt.Println("Terima Kasih Telah Membeli")
					}
				}
				
				
			} else if jawaban==2 {
					 
				fmt.Println("Masukan Jumlah Uang Anda : ")
				fmt.Scan(&uang)
				if uang < tot {
					fmt.Println("Mohon Maaf Uang Anda Kurang, Mohon Membayar Sesuai Pesanan")
				} else{
					kembalian = uang - tot 
					fmt.Println("Kembalian Anda : ",kembalian)
					fmt.Println("Terima Kasih Telah Membeli")
				}
			}
			
		} else if pilihan==2 {		
			fmt.Println("Masukkan Menu Baru")
			fmt.Scan(&pilih)
			fmt.Println("Masukan Harga Menu: ")
			fmt.Scan(&nhar)
			data[jum+1].nama= pilih
			data[jum+1].harga= nhar
			jum++
			
			fmt.Println("Menu Berhasil Ditambahkan")
			fmt.Println(jum)
			
		} else if pilihan==3 {
			fmt.Println ("Terimakasih telah datang NL Ice Store")
			fmt.Println ("Semoga hari anda menyenangkan")
		}
	
		fmt.Print("Apakah Anda mau kembali ke menu utama ? (Y/T)")
		fmt.Scan(&jawab)
		if jawab == "Y" || jawab == "y"{
			cek = false
		}else if jawab == "N" || jawab == "n"{
			cek = true
		}
	}
	
	
}



func Search (data arrayType, name string) int {
	var i int
	i = 0 
	for data[i].nama != name{
		i++
	}
	
	
	
	return i
}
	
func insertionSort (data *arrayType ,jum int) {
	for i := 0; i <= jum; i++ {
		tmp := data[i].harga
		tmp2:= data[i].nama
		j := i
		for j > 0 && data[j-1].harga > tmp {
			data[j].harga = data[j-1].harga
			data[j].nama = data[j-1].nama
			j--
		}
	data[j].harga = tmp
	data[j].nama = tmp2
	}
}

func Menu(data *arrayType,jum int){
	var i int
	insertionSort(data,jum)
	for i = 0;i<=jum; i++{
		fmt.Println(i+1,"Nama Makanan :", data[i].nama)
		fmt.Println("Harga :" , data[i].harga)
	}
}
