package common_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/Cyb3r-Jak3/common/v3"
	"github.com/Cyb3r-Jak3/common/v3/colors"
)

func BenchmarkJSONResponse(b *testing.B) {
	r, _ := http.NewRequest("GET", "/", nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeRequest(r, JSONTest)
	}
}

func BenchmarkWOAllowedMethod(b *testing.B) {
	r, _ := http.NewRequest("GET", "/", nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeRequest(r, StringTest)
	}
}

func BenchmarkAllowedMethod(b *testing.B) {
	r, _ := http.NewRequest("GET", "/", nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeRequest(r, common.AllowedMethod(StringTest, "GET"))
	}
}

func BenchmarkContentResponse(b *testing.B) {
	r, _ := http.NewRequest("GET", "/", nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeRequest(r, ContentTest)
	}
}

func BenchmarkStringResponse(b *testing.B) {
	r, _ := http.NewRequest("GET", "/", nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeRequest(r, StringTest)
	}
}

func BenchmarkJSONMarshall(b *testing.B) {
	r, _ := http.NewRequest("GET", "/", nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeRequest(r, JSONMarshalTest)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = common.GenerateRandInt(5)
	}
}

func BenchmarkJSONParse(b *testing.B) {
	testStruct := new(testStruct)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = common.ParseYamlOrJSON("../testData/parsetest.json", testStruct)
	}

}

func BenchmarkYAMLParse(b *testing.B) {
	testStruct := new(testStruct)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = common.ParseYamlOrJSON("../testData/parsetest.yml", testStruct)
	}

}

func BenchmarkStringSearch2(b *testing.B) {
	array := []string{"hello", "world"}
	for i := 0; i < b.N; i++ {
		common.StringSearch("hello", array)
	}

}
func BenchmarkStringSearch10(b *testing.B) {
	array := []string{"hello", "world", "hi", "mom", "and", "here", "are", "ten", "random", "words"}
	for i := 0; i < b.N; i++ {
		common.StringSearch("hello", array)
	}

}
func BenchmarkFloatSearch2(b *testing.B) {
	array := []float64{1.1, 1.2}
	for i := 0; i < b.N; i++ {
		common.FloatSearch(1.9, array)
	}
}

func BenchmarkFloatSearch10(b *testing.B) {
	array := []float64{1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0}
	for i := 0; i < b.N; i++ {
		common.FloatSearch(1.9, array)
	}
}

func BenchmarkIntSearch2(b *testing.B) {
	array := []int{1, 2}
	for i := 0; i < b.N; i++ {
		common.IntSearch(1, array)
	}
}

func BenchmarkIntSearch10(b *testing.B) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		common.IntSearch(1, array)
	}
}

func BenchmarkGetEnv(b *testing.B) {
	os.Setenv("test", "value")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = common.GetEnv("test", "")
	}
}

func BenchmarkGetEnvMissing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		common.GetEnv("missing", "value")
	}
}

func BenchmarkSHA256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = common.HashFile("256", "hash.go")
	}
}

func BenchmarkSHA384(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = common.HashFile("384", "hash.go")
	}
}

func BenchmarkSHA512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = common.HashFile("512", "hash.go")
	}
}

func BenchmarkToHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		colors.ToHex(colors.Aqua)
	}
}

func BenchmarkSkipRoot(b *testing.B) {
	jsonString := `{"root": {"key": "value"}}`
	for i := 0; i < b.N; i++ {
		common.SkipRoot([]byte(jsonString))
	}
}

func BenchmarkEnvironMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		common.EnvironMap()
	}
}
