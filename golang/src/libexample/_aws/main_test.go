package main

import (
	"testing"
	"github.com/stripe/aws-go/gen/s3"
)

func BenchmarkS3GetObject1(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	var s3Client *s3.S3
	s3Client = credentials()
	
	for i := 0; i < b.N; i++ {
		printGetObject(s3Client)
	}
}

func BenchmarkS3GetObject2(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var s3Client *s3.S3
		s3Client = credentials()
		printGetObject(s3Client)
	}
}
