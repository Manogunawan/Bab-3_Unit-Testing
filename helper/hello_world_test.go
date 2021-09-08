package helper

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

//func TestHelloWorldAssert(t *testing.T) {
//	result := HelloWorld("Manogunawan")
//	assert.Equal(t, "Hello Manogunawan", result, "Result must be 'Hello Manogunawan'")
//	fmt.Println("TestHelloWorld with Assert Done")
//}

//BenchmarkTable
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{	name: "Manogunawan",
			request: "Mano",
		},
		{
			name: "Resqi",
			request: "Res",
		},
	}
	for _, benchmark := range benchmarks{
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

//SubBench
func BenchmarkSub(b *testing.B) {
	b.Run("Dian", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Dian")
		}
	})
	b.Run("Novtri", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Novtri")
		}
	})
}

//Bench
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Mano")
	}
}

func BenchmarkHelloWorldMano(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Manogunawan")
	}
}

//Table Test
func TestTableHelloWorld(t *testing.T)  {
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "Manogunawan",
			request: "Manogunawan",
			expected: "Hello Manogunawan",
		},
		{
			name: "Dian",
			request: "Dian",
			expected: "Hello Dian",
		},
		{
			name: "Belly",
			request: "Belly",
			expected: "Hello Belly",
		},
	}
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

//Test SubTest
func TestSubTest(t *testing.T)  {
	t.Run("Manogunawan", func(t *testing.T) {
		result := HelloWorld("Manogunawan")
		require.Equal(t, "Hello Manogunawan", result, "Result must be 'Hello Manogunawan'")
	})
	t.Run("Gultom", func(t *testing.T) {
		result := HelloWorld("Gultom")
		require.Equal(t, "Hello Gultom", result, "Result must be 'Hello Manogunawan'")
	})
	t.Run("Dian", func(t *testing.T) {
		result := HelloWorld("Dian")
		require.Equal(t, "Hello Dian", result, "Result must be 'Hello Manogunawan'")
	})
}


//Test Main
func TestMain(m *testing.M)  {
	//before
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	//after
	fmt.Println("AFTER UNIT TEST")
}

//Test Skip
func TestSkip(t *testing.T)  {
	if runtime.GOOS == "mano" {
		t.Skip("Can not run on Mac Os")
	}

	result := HelloWorld("Manogunawan")
	require.Equal(t, "Hello Manogunawan", result, "Result must be 'Hello Manogunawan'")
}

//Assertion
func TestHelloWorldAssertRequire(t *testing.T) {
	result := HelloWorld("Manogunawan")
	require.Equal(t, "Hello Manogunawan", result, "Result must be 'Hello Manogunawan'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Manogunawan")

	if result != "Hello Manogunawan"{
		panic("Result is not 'Hello World'")
	}
}

func TestHelloWorldMan(t *testing.T) {
	result := HelloWorld("Nores")

	if result != "Hello Nores" {
		t.Fatal("Result mus be 'Hello Nores'")
	}
	fmt.Println("TestHelloWorldMan Done")
}