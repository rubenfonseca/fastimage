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
	// url := "http://upload.wikimedia.org/wikipedia/commons/9/9a/SKA_dishes_big.jpg"
	url := "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_120x44dp.png"
	// url := "http://loremflickr.com/500/500"

	for i := 0; i < b.N; i++ {
		it, is, err := DetectImageTypeWithTimeout(url, 10000)
		b.Logf("type:%v, size:%v, err:%v", it, is, err)
	}
}
