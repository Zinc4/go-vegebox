![splash screen1](https://github.com/Zinc4/go-vegebox/assets/65228679/575fa45c-7308-460e-9b39-1e6f7ff0e8ed)

# Vegebox | E-commerce Backend Project

## About Project

VEGEBOX merupakan project berbasis ecommerce yang menggunakan backend Golang. Proyek ini bertujuan untuk menghubungkan konsumen dengan petani lokal dan produsen makanan, serta memudahkan dalam memesan produk-produk segar. Platform ini didesain untuk menyederhanakan proses pembelian produk makanan segar dan lokal secara efisien dan praktis.

## Features

1. **User**
   - **Authentication:** Register, Login, Verifikasi Email dengan, Kirim Ulang OTP
   - **Profile:** Mengubah data user sendiri, update user avatar sendiri
   - **Produk:** Melihat semua produk, Melihat kategori produk
   - **Cart:** Tambah produk ke keranjang, Edit produk di keranjang, Hapus produk di keranjang, Melihat checkout, Melihat keranjang sendiri
   - **Order:** Membuat pesanan berdasarkan keranjang, Melihat pesanan sendiri
   - **Transaction:** Melakukan pembayaran pesanan, melihat riwayat pembayaran sendiri
2. **Admin**
   - **Authentication:** Login dan Register
   - **Produk:** Mendapatkan data semua produk,Menambah produk, Menghapus produk, Memperbarui produk, Membuat kategori produk, Menghapus kategori produk
   - **User:** Melihat semua users, Menghapus user
   - **Transaction:** Melihat semua riwayat pembayaran users

## Tech Stacks

- Code Editor : **[Visual Studi Code](https://code.visualstudio.com/download)**
- Backend Framework : **[Echo](https://echo.labstack.com/)**
- ORM : **[GORM](https://gorm.io/index.html)**
- API Deployment : **[AWS EC2](https://aws.amazon.com/pm/ec2/?gclid=CjwKCAjw0YGyBhByEiwAQmBEWgU7A2vW-SxsWNH4QFqQIJ1ahXK9YST-yb4vVPm6S99PRFvkFqPRqxoCXQcQAvD_BwE&trk=361ccc4f-68c4-4038-bf6c-0586bee109dc&sc_channel=ps&ef_id=CjwKCAjw0YGyBhByEiwAQmBEWgU7A2vW-SxsWNH4QFqQIJ1ahXK9YST-yb4vVPm6S99PRFvkFqPRqxoCXQcQAvD_BwE:G:s&s_kwcid=AL!4422!3!476956795566!e!!g!!aws%20ec2!11543056243!112002963829)**
- Database Deployment : **[RDS](https://aws.amazon.com/free/database/?gclid=CjwKCAjw0YGyBhByEiwAQmBEWviCXIEtUNS0IlMQSE-o64FINgri6vL8QCihqB6qUot-jJx5eReF2hoC4N4QAvD_BwE&trk=fc551e06-56b0-418c-9ddd-5c9dba18569b&sc_channel=ps&ef_id=CjwKCAjw0YGyBhByEiwAQmBEWviCXIEtUNS0IlMQSE-o64FINgri6vL8QCihqB6qUot-jJx5eReF2hoC4N4QAvD_BwE:G:s&s_kwcid=AL!4422!3!548908918497!e!!g!!aws%20rds!11543056228!112002957989)**
- ERD Design tools : **[Lucidchart](https://www.lucidchart.com/pages/)**
- Containerize : **[Docker](https://www.docker.com/)**
- Password Hash : **[Argon2id](https://github.com/alexedwards/argon2id)**
- Testing : **[Testify](https://github.com/stretchr/testify)**
- Payment Gateway : **[Midtrans](https://dashboard.midtrans.com/login)**
- Image Cloud : **[Cloudinary](https://cloudinary.com/)**
- Mail Sender : **[Go-mail](github.com/wneessen/go-mail)**
- API Documentation : **[Postman](https://www.postman.com/)**

## API Documentation

- API SPEC : **[API-SPEC](https://documenter.getpostman.com/view/21327885/2sA3JNaLGA#1cbfed67-9322-4b11-9bbc-0b465d73f18c)**

## ERD

- ERD : **ERD-Vegebox ![ERD-Vegebox](https://github.com/Zinc4/go-vegebox/assets/65228679/a11cd85a-24e7-4092-8a1d-5830600201cf)**

## Setup

Clone this repository

```bash
git clone https://github.com/Zinc4/go-vegebox.git
```

Navigate to the project directory:

```bash
cd go-vegebox
```

Create .env file using this command and open the .env file and edit the value

```bash
cp .env.example .env
```

Install module using this command

```bash
go mod download
```

Run the project using this command

```bash
go run main.go
```
