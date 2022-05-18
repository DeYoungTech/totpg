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

package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/deyoungtech/totpg/internal/validation"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

var (
	algorithm *string
	digits    *int
	period    *int
	secret    *string
	skew      *int
)

func init() {
	algorithm = flag.String("algorithm", "SHA1", "Algorithm to use for HMAC. Valid options: SHA1. SHA256, SHA512, MD5. Defaults to SHA1.")
	digits = flag.Int("digits", 6, "The number of digits for the One Time Passcode. Defaults to 6")
	period = flag.Int("period", 30, "Number of seconds a TOTP hash is valid for. Six and Eight are the most common values. Defaults to 30 seconds.")
	secret = flag.String("secret", "", "The Base32 encoded secret key associated with the Time-based One Time Password.")
	skew = flag.Int("skew", 0, "Periods before or after the current time to allow. Value of 1 allows up to Period of either side of the specified time. Defaults to 0 allowed skews.")
	flag.Parse()
}

func main() {

	passcode, err := totp.GenerateCodeCustom(validation.ValidSecret(secret), time.Now().UTC(), totp.ValidateOpts{
		Period:    uint(*period),
		Skew:      uint(*skew),
		Digits:    otp.Digits(*digits),
		Algorithm: validation.ValidAlgorithm(*algorithm),
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", passcode)
}
