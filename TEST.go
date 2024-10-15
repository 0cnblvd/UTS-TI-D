//POINTER

package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Println("Sebelum:", x)

	// Menggunakan pointer untuk mengubah nilai x
	ubahNilai(&x)
	fmt.Println("Setelah:", x)
}

func ubahNilai(n *int) {
	*n = 20
}

//karena coding buku sangat lah panjang, saya mencari cara apakah ada variabel yang dapat mempermudah?
