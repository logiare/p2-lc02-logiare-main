[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-2e0aaae1b6195c2367325f4f02e2d04e9abb55f0b24a779b69b11b9e10269abc.svg)](https://classroom.github.com/online_ide?assignment_repo_id=23735798&assignment_repo_type=AssignmentRepo)
# FTGO-P2-V1-LC2
## RULES
1. **Untuk kampus remote**: **WAJIB** melakukan **share screen**(**DESKTOP/ENTIRE SCREEN**) dan **unmute microphone** ketika Live Code
berjalan (tidak melakukan share screen/salah screen atau tidak unmute microphone akan di ingatkan).
2. Kerjakan secara individu. Segala bentuk kecurangan (mencontek ataupun diskusi) akan menyebabkan skor live code ini 0.
3. Clone repo ini kemudian buatlah **branch dengan nama kalian**.
4. Kerjakan pada file Golang (\*.go) yang telah disediakan.
5. Waktu pengerjaan: **150 menit** untuk **2 soal**.
6. **Pada text editor hanya ada file yang terdapat pada repository ini**.
7. Membuka referensi eksternal seperti Google, StackOverflow, dan MDN diperbolehkan.
8. Dilarang membuka repository di organisasi tugas, baik pada organisasi batch sendiri ataupun batch lain, baik branch sendiri maupun branch orang
lain (**setelah melakukan clone, close tab GitHub pada web browser kalian**).
9. Lakukan `git push origin <branch-name>` dan create pull request **hanya jika waktu Live Code telah usai (bukan ketika kalian sudah selesai
mengerjakan)**. Tuliskan nama lengkap kalian saat membuat pull request dan assign buddy.
10. **Penilaian berbasis logika dan hasil akhir**. Pastikan keduanya sudah benar.


## Notes
Live code ini memiliki bobot nilai sebagai berikut:

| **Kriteria** | **Deskripsi** |
|--------------|---------------|
| **Problem Solving - REST API** | Implementasi semua endpoint API yang diminta, termasuk handling edge cases dan negative cases. |
| **Autentikasi dan Keamanan** | Penggunaan JWT untuk autentikasi, proteksi endpoint dengan benar, dan keamanan data pengguna. |
| **Desain Database** | Desain ERD yang sesuai dengan requirement dan penggunaan hubungan yang tepat antara tabel. |
| **Implementasi ORM** | Penggunaan GORM yang benar untuk mengelola database, termasuk raw query untuk kasus tertentu. |
| **Handling Error** | Implementasi error handling yang benar pada setiap endpoint sesuai dengan status kode (400, 404, 500). |
| **Best Practices REST API** | Mengikuti best practices REST API dalam hal metode HTTP, struktur URL, dan format response. |
| **Kualitas Kode** | Kode bersih, terstruktur, mudah dipahami, dan mengikuti konvensi penulisan kode yang baik. |



### KETENTUAN
Here are some dos and don'ts to consider when working on this livecode:

Do:

- Identify all the entities involved in the system, including customers, products, orders, and payments.
- Use clear and concise naming conventions for each entity and attribute.
- Clearly define the relationships between entities, such as the one-to-many relationship between customers and orders, and the many-to-many relationship between orders and products.
- Include foreign keys in the appropriate tables to ensure data integrity and enforce referential integrity constraints.
- Use appropriate data types for each attribute, such as using a date data type for order dates and a decimal data type for prices.
- Use indexes where necessary to improve performance for frequently queried data.

Don't:

- Use vague or ambiguous naming conventions that make it difficult to understand the purpose of each entity and attribute.
- Overcomplicate the relationships between entities, such as creating unnecessary junction tables or introducing circular references.
- Store redundant data in multiple tables, which can lead to data inconsistencies and update anomalies.
- Use inappropriate data types for attributes, such as using a string data type for numerical values.
- Neglect to define primary and foreign keys, which can lead to data integrity issues and orphaned records.
- Neglect to consider the performance implications of our schema, such as the need for indexes on frequently queried data.


# LIVECODE 2

## Objectives
- Mampu Merancang dan Mengimplementasikan Database PostgreSQL
- Mampu Membangun REST API dari Awal Menggunakan Go dan Echo Framework


## Directions
Sebuah perusahaan pariwisata terkemuka tidak hanya mengorganisir berbagai tur ke destinasi unik, tetapi juga baru-baru ini menjelajahi ranah digital dengan menyediakan portal online untuk pelanggannya. Di sini, calon pengguna dapat melakukan registrasi, login, menjelajahi data tur yang tersedia, dan melakukan pemesanan. Setelah berhasil mendaftar atau login, setiap pelanggan akan menerima token JWT, yang memungkinkan mereka mengakses beberapa fungsionalitas di platform tersebut.


## DB Requirement Tasks

Untuk memenuhi kebutuhan aplikasi, Anda perlu mendesain ERD yang mencakup tabel-tabel berikut:

A. Customers:

Attributes: customer_id (primary key), user_id (foreign key), name, email, phone_number, address.

Explanation: Menyimpan detail pelanggan yang melakukan booking tur. Setiap pelanggan juga merupakan pengguna terdaftar yang terhubung melalui user_id.

B. Users:

Attributes: user_id (primary key), email, password_hash, last_login_date, jwt_token.

Explanation: Menyimpan informasi digital pengguna terdaftar, termasuk kata sandi terenkripsi dan token autentikasi untuk akses yang aman.

C. Tours:

Attributes: tour_id (primary key), tour_name, tour_description, tour_price, tour_duration.

Explanation: Menyimpan informasi tentang tur yang tersedia, termasuk nama, deskripsi, harga, dan durasi tur.

D. Bookings:

Attributes: booking_id (primary key), customer_id (foreign key), booking_date, booking_status.

Explanation: Menyimpan detail setiap pemesanan tur yang dilakukan oleh pelanggan.

E. Tour_Bookings:

Attributes: tour_booking_id (primary key), booking_id (foreign key), tour_id (foreign key), tour_start_date, tour_end_date, num_of_people.

Explanation: Menyimpan hubungan antara pelanggan dan tur yang mereka pesan, termasuk tanggal mulai dan berakhirnya tur serta jumlah peserta.

F. Payments:

Attributes: payment_id (primary key), booking_id (foreign key), payment_date, payment_amount.

Explanation: Menyimpan catatan transaksi pembayaran untuk setiap pemesanan tur.

Untuk memudahkan kalian silahkan lihat pada l2p2.sql pada direktori Livecode ini.

Tugas : 
1. Buatlah ERD untuk kebutuhan database diatas. 

Harap sertakan ERD.png pada Direktori Pengumpulan Livecode 2 ini. 


## RELEASE 2 : Build API Development 

Buatlah REST API yang berfungsi untuk mengelola platform layanan pariwisata. API yang Anda buat harus mencakup fungsionalitas manajemen pengguna dan manajemen pemesanan tur, dengan mempertimbangkan requirement tech stack berikut:

a. Go - Echo: Framework ini akan digunakan untuk mengembangkan REST API. Echo menyediakan fitur routing yang efisien dan middleware yang mendukung autentikasi dan error handling.

b. PostgreSQL dengan ORM GORM: Gunakan PostgreSQL sebagai database dan ORM GORM untuk mengelola interaksi dengan database. ORM ini akan memudahkan dalam melakukan operasi CRUD dan query ke database.

c. JWT untuk Autentikasi: Gunakan JSON Web Token (JWT) untuk mengamankan endpoint yang membutuhkan autentikasi. JWT akan dikirimkan dalam header Authorization pada setiap request yang memerlukan akses aman.



## Details - API User Management

A. Register (POST /users/register):

Endpoint ini digunakan untuk mendaftarkan pengguna baru di platform.
Pengguna harus menyediakan email yang valid dan kata sandi.
Setelah pendaftaran berhasil, data pengguna baru akan disimpan di database, dan response akan mengembalikan informasi pengguna (tanpa kata sandi).

Contoh response sukses (status code: 201):

```sh
{
  "user_id": 1,
  "email": "john.doe@example.com"
}

```

- Error Handling:

400 Bad Request: Jika email atau kata sandi tidak valid atau email sudah terdaftar.



B. Login (POST /users/login)

- Endpoint ini digunakan untuk login pengguna yang sudah terdaftar.
- Pengguna harus menyediakan email dan kata sandi yang valid.
- Setelah autentikasi berhasil, server akan mengembalikan token JWT yang dapat digunakan untuk mengakses endpoint yang memerlukan autentikasi.

Contoh response sukses (status code: 200):

```sh
{
   "token": "jwt_token_value"
}

```

Error Handling:

400 Bad Request: Jika email atau kata sandi tidak valid.

404 Not Found: Jika pengguna tidak ditemukan.

B. Booking Management:

Get All Bookings (GET /bookings):

- Endpoint ini digunakan untuk mendapatkan daftar semua pemesanan yang dilakukan oleh pengguna yang sedang login.

- Endpoint ini hanya dapat diakses oleh pengguna yang telah terautentikasi (JWT token harus disertakan di header request).

Contoh response sukses (status code: 200):

```sh
[
  {
    "booking_id": 1,
    "tour_name": "Adventure to Mount Everest",
    "booking_date": "2024-08-16",
    "booking_status": "Completed"
  },
  ...
]
```

- Error Handling:

404 Not Found: Jika tidak ada pemesanan yang ditemukan untuk pengguna tersebut.


2. Get Unpaid Bookings (GET /bookings/unpaid):

- Endpoint ini digunakan untuk mendapatkan daftar pemesanan yang belum dibayar oleh pengguna yang sedang login.

- Endpoint ini hanya dapat diakses oleh pengguna yang telah terautentikasi.
Contoh response sukses (status code: 200):

```sh
[
  {
    "booking_id": 2,
    "tour_name": "Trip to Bali",
    "booking_date": "2024-08-15",
    "booking_status": "Pending"
  },
  ...
]
```

- Error Handling:

404 Not Found: Jika tidak ada pemesanan yang belum dibayar ditemukan untuk pengguna tersebut.

3. Get Vehicle Availability (GET /vehicles/availability): 

- Endpoint ini digunakan untuk melihat pendapatan dari semua tur yang tersedia.

- Pengguna yang telah login dapat mengakses daftar tur beserta pendapatan yang diperoleh.

Contoh response sukses (status code: 200):

```sh
[
   {
    "tour_id": 1,
    "tour_name": "Adventure to Mount Everest",
    "total_earnings": 5000000
  },
  ...
]
```

Error Handling:

- 404 Not Found: Jika tidak ada tur yang ditemukan.

4. Get Total Customers (GET /reports/total-customers):

- Endpoint ini digunakan untuk mendapatkan jumlah total pelanggan yang terdaftar di platform.

- Endpoint ini hanya dapat diakses oleh pengguna yang telah terautentikasi (JWT token harus disertakan di header request).

Contoh response sukses (status code: 200):

```sh
{
  "total_customers": 150
}
```

5. Get Total Bookings Per Tour (GET /reports/bookings-per-tour):

- Endpoint ini digunakan untuk mendapatkan jumlah total pemesanan untuk setiap tur yang tersedia di platform.

- Endpoint ini hanya dapat diakses oleh pengguna yang telah terautentikasi (JWT token harus disertakan di header request).

```sh
 {
    "tour_name": "Adventure to Mount Everest",
    "total_bookings": 20
  },
  {
    "tour_name": "Trip to Bali",
    "total_bookings": 35
  },
  ...
```

Error Handling:

- 404 Not Found: Jika tidak ada data pemesanan yang ditemukan untuk tur tertentu.

- 500 Internal Server Error: Jika terjadi masalah pada server yang menghalangi pengambilan data.

