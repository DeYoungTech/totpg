/**
 *  Copyright 2022 Brittan DeYoung - DeYoung Technologies LLC.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

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
