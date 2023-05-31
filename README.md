# Portofolio
Aplikasi Simple Manajemen Pembiayaan Kendaraan menggunakan Go  


Aplikasi Simple Manajemen Pembiayaan Kendaraan menggunakan Go  


Fitur yang tersedia 
Pendataan Nasabah menggunakan Database Mysql
Management data nasabah
Pendataan Transaksi Setiap Nasabah
Validasi transaksi Nasabah
Pendataan Pembayaran Nasabah

API yang bisa Digunakan

http://127.0.0.1:5025/api/v1/sks/public/list/customer 

http://127.0.0.1:5025/api/v1/sks/public/register/customer
{
    "id": 3,
    "nik": "14045",
    "full_name": "Roman",
    "legal_name": "Roman Legal",
    "place_of_birth": "bandung",
    "birth_date": "1990-01-01",
    "gaji": 2000000,
    "foto_ktp": "foto_ktp.jpg",
    "foto_selfie": "foto_selfie.jpg"
}

http://127.0.0.1:5025/api/v1/sks/public/insert/limit
{
    "id":3
}

http://127.0.0.1:5025/api/v1/sks/public/list/limit
{
    "id":3
}

http://127.0.0.1:5025/api/v1/sks/public/transaction
{
    
    "customer_id": 2,
    "contract_no": "K03",
    "tenor": 6,
    "otr": 700000,
    "asset_name": "Mobil Bensin",
    "transaction_type": "Dealer"
}

http://127.0.0.1:5025/api/v1/sks/public/update/balance
{
    "total":1000,
    "customer_id":2,
    "tenor":6,
    "limit_ammount":2000000
}

http://127.0.0.1:5025/api/v1/sks/public/payments
{
    "customer_id":1,
    "contract_no":24,
    "installment":141665,
    "ammount":849994
}

