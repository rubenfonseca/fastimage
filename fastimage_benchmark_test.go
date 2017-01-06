package fastimage

import "testing"

// func BenchmarkSample(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		if x := fmt.Sprintf("%d", 42); x != "42" {
// 			b.Fatalf("Unexpected string: %s", x)
// 		}
// 	}
// }

func BenchmarkCustomTimeout(b *testing.B) {
	url := "http://upload.wikimedia.org/wikipedia/commons/9/9a/SKA_dishes_big.jpg"

	for i := 0; i < b.N; i++ {
		DetectImageTypeWithTimeout(url, 500)
	}
}
