# Prakerja Batch 9

## Days 1 ~ Belajar Git

Tugas
- Buat sebuah repository di Github
- Implementasikan penggunaan branching yang terdiri dari master,
development, featureA, dan featureB
- Implementasikan intruksi git untuk push, pull, stash dan merge
- Implementasikan sebuah penanganan conflict di branch developement ketika
setelah merge dari branch featureA lalu merge dari branch featureB (Conflict
bisa terjadi jika kedua branch mengerjakan di file dan line code yang sama)
- Gunakan merge no fast forward
- Screenshot hasil network di github

## Days 2 ~ Basic programing - Go 

Tugas 

- buatlah sebuah program untuk menentukan bilang prima
- buatlah sebuah program untuk menentukan bilangan kelipatan 7
- buatlah sebuah program untuk menghitung luas trapesium


## Days 3 ~ Structure data
Tugas 

- Buatlah sebuah program menggabungkan 2 array yang diberikan, dan jangan sampai terdapat nama yang sama di data yang sudah tergabung tadi!
- buatlah sebuah program yang dapat menghitung berapa banyak sebuah string yang sama didalam sebuah slice!
- Buat program sesuai dengan deskripsi di bawah. Input merupakan variable
string berisi kumpulan angka. Output merupakan list / array berisi angka yang
hanya muncul 1 kali pada input.

##  Days 4 ~ String, Advance Function, Pointer, Heap & Stack, Struct, Method, Interface, Garbage Collector, Package & Error Handling
Tugas

- Buatlah sebuah method untuk menghitung perkiraan jarak yang bisa
ditempuh berdasarkan bahan bakar yang sedang terisi (1.5 L / mill), method
tersebut receiver kepada struct Car yang memiliki property type dan fuelln!
- Buat sebuah struct dengan nama Student yang mempunyai properti name
dan score dalam bentuk slice kemudian simpan data siswa sebanyak 5 siswa.
Setelah 5 siswa dimasukkan maka program menunjukkan skor rata-rata,
siswa yang memiliki skor minimum dan maksimal? (implementasikan
method)
- Buatlah program di Golang untuk menemukan nilai maksimum serta minimum
di antara 6 angka inputan. Gunakan multiple return fungsi, pointer untuk
referencing maupun deferencing!

## Days 5 ~ Introduction to RESTfull api & echo
Tugas

- tampilkan data yang diperoleh dari API : [link api](https://jsonplaceholder.typicode.com/posts)
- tampilkan data dengan id 3 yang diperoleh dari API : [link api](https://jsonplaceholder.typicode.com/posts)
- simpan data postingan ke server melalui API : [link api](https://jsonplaceholder.typicode.com/posts)
- hapus suatu data melalui API : [link api](https://jsonplaceholder.typicode.com/posts)
- Clone repository Github : [link github](https://github.com/kokolopo/task-intro-echo.git) lalu
lengkapi tiap-tiap controller yang ada pada file server.go


## Days 6 ~ Connecting DB dan ORM + mvc

Buat API CRUD User dengan spesifikasi seperti berikut!
- Pada bagian Sample Code artinya sudah disediakan contoh code yang bisa kamu
implementasikan.
- Pada bagian Need Solution Code kamu perlu membuat sendiri code-nya!






## Catatan

* nilai null dgi golang adalah nil
* pada for range terdapat 2 buah nilai yaitu key/index dan value. jika hanya ingin mendapatkan index/key value tidak perlu ditambahkan. tapi jika ingin mendapatkan value tanpa key/index diganti dengan _ (undescore)
* Interface digunakan untuk komunikasi
* Pada struct terdapat id tapi tidak bisa diexport kalo mau diexport harus dirubah jadi Id(huruf depan kapital)