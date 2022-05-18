package validation

import (
	"flag"
	"regexp"

	"github.com/pquerna/otp"
)

func ValidAlgorithm(algorithm string) otp.Algorithm {
	switch algorithm {
	case "SHA1":
		return otp.AlgorithmSHA1
	case "SHA256":
		return otp.AlgorithmSHA256
	case "SHA512":
		return otp.AlgorithmSHA512
	case "MD5":
		return otp.AlgorithmMD5
	}
	flag.PrintDefaults()
	panic("Unsupported algorithm provided")
}

func ValidSecret(secret *string) string {
	match, _ := regexp.Match("^(?:[A-Z2-7]{8})*(?:[A-Z2-7]{2}={6}|[A-Z2-7]{4}={4}|[A-Z2-7]{5}={3}|[A-Z2-7]{7}=)?$", []byte(*secret))

	if match {
		return *secret
	} else {
		flag.PrintDefaults()
		panic("Unsupported secret provided")
	}
}
