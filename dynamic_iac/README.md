# SMKDEV Scholarships Batch 2 - Final Submissions

Ini merupakan Final Submission yang akan menguji pemahaman pembelajaran dan pemahaman kamu selama live session berlangsung. Kamu dapat melihat live session yang telah diberikan sebelumnya untuk membantu kamu dalam menyelesaikan Final Submission ini

## Storyboard

Disini kamu perlu untuk menyelesaikan setiap fungsi yang kosong dengan solusi yang dapat kamu implementasikan. Kamu juga bisa kembali melihat Live Session sebelumnya dimana implementasi yang ada pada Submission ini sangat sering dibahas disana. Jadi kamu perlu menyelesaikan Submission ini berdasarkan Live Session atau gaya penyelesaian yang berbeda selama implementasi Projek yang diberikan dapat diajalankan dengan baik dan benar

Untuk proses inisiasi Final Submission ini, dimohon untuk menginisasi projek dengan `go mod init nama-projek` sehingga dapat membantu kamu dalam proses pengerjaan Final Project ini. Package yang kamu perlukan dalam mengerjakan Final Project ini adalah sebaga berikut
- Echo Framework
- Viper
- Gorm
- PostgreSQL Driver dari Gorm

> Tidak menutup kemungkinan untuk package tambahan agar projek ini dapat kamu kerjakan dengan baik

## Aturan

- Kamu perlu mengimplementasikan solusi yang ada pada beberapa layer dari business logic yang telah didefinisikan pada `/internal`. Secara detail kamu perlu memperhatikan dan menyelesaikan, yakni
  - `/internal/app/handlers`
  - `/internal/app/services`
  - `/internal/app/repositories`
  - `env.yaml`: Terkhusus untuk environment variables yang akan digunakan, baik koneksi postgres dengan Viper boleh disesuaikan dengan environment atau konfigurasi OS masing-masing. Jadi implementasi dapat flexible dilakukan. Kamu dapat menonton kembali Live Class kita pada saat melakukan konfigurasi dan koneksi ke PostgreSQL dengan Viper pada `/internal/configs`

## Objectiveness

Selesaikan keseluruhan kode kosong dengan implementasi kamu. Untuk pengembangan dan API Testing, kamu dapat menggunakan Postman atau Thunder Client untuk menguji setiap fungsi/routes yang telah kamu selesaikan.

Kamu juga bisa melihat contoh dokumentasi API yang telah terimplementasi selama Live Class berlangsung disini: https://documenter.getpostman.com/view/33035417/2sA2r81PZF


## Submissions

Jikalau kamu telah menyelesaikan Final Submission ini, kamu perlu submit solusi tersebut dalam bentuk **.zip dan diupload ke Google Drive dan kemudian upload link file** tersebut ke Google Form yang bisa diakses pada link berikut: https://docs.google.com/forms/d/e/1FAIpQLSe0qKXNFSoT7ljz0_UjmBhNISSZXTUTGDfd4KbTs8M2JRPzQw/viewform?usp=sharing

> Jangan lupa atur General Access menjadi **Anyone with the link** menjadi **Viewer** agar bisa diakses dan ditesting.

### Penilaian Final Project

- Sumber Solusi yang telah disubmit akan dilakukan testing, namun akan dijalankan di local computer mentor/reviewer. Kalau ingin melihat contoh testing code (beta) yang akan digunakan, dapat dilihat pada link berikut: https://github.com/smkdev-id/promotion_tracking_dashboard/tree/main/tests.
- Testing yang ada pada dokumentasi di link tersebut belum sepenuhnya update dikarenakan bersifat **internal** selama timeline Final Project ini berlangsung. Testing Code yang stable akan dipublikasi setelah masa bootcamp ini berakhir
- Jikalau kamu menghiraukan hasil coverage yang ada, kamu bisa improve lebih baik lagi. Namun tidak menjadi nilai tambah untuk submission ini.
