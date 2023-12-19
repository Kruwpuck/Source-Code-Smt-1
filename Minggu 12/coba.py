nilai_siswa = {}
i = 1
jumlah_siswa = int(input("Jumlah siswa : "))
while i <= jumlah_siswa:
    nama = input("Masukkan nama : ")    
    nilai = int(input("Masukkan nilai : "))
    nilai_siswa[nama] = nilai
    i += 1

def hitung_rata_rata(nilai_siswa):
    return sum(nilai_siswa.values())//len(nilai_siswa)
def hitung_tertinggi(nilai_siswa):
    return max(nilai_siswa.values())
def hitung_terendah(nilai_siswa):
    return min(nilai_siswa.values())
print(hitung_rata_rata(nilai_siswa))
print(hitung_tertinggi(nilai_siswa))
print(hitung_terendah(nilai_siswa))
