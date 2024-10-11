// Copyright (c) 2012-2016 The Revel Framework Authors, All rights reserved.
// Revel Framework source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package validation

import (
	"regexp"

	boovalidation "github.com/boo-admin/boo/validation"
)

var emptyMessageArgs = []interface{}{}

type Validator = boovalidation.Validator

type Required = boovalidation.Required

func ValidRequired() Required {
	return boovalidation.ValidRequired()
}

type Min = boovalidation.Min

func ValidMin(min int) Min {
	return boovalidation.ValidMin(min)
}

func ValidMinFloat(min float64) Min {
	return boovalidation.ValidMinFloat(min)
}

type Max = boovalidation.Max

func ValidMax(max int) Max {
	return boovalidation.ValidMax(max)
}

func ValidMaxFloat(max float64) Max {
	return boovalidation.ValidMaxFloat(max)
}

// Range requires an integer to be within Min, Max inclusive.
type Range = boovalidation.Range

func ValidRange(min, max int) Range {
	return boovalidation.ValidRange(min, max)
}

func ValidRangeFloat(min, max float64) Range {
	return boovalidation.ValidRangeFloat(min, max)
}

// MinSize requires an array or string to be at least a given length.
type MinSize = boovalidation.MinSize

func ValidMinSize(min int) MinSize {
	return boovalidation.ValidMinSize(min)
}

// MaxSize requires an array or string to be at most a given length.
type MaxSize = boovalidation.MaxSize

func ValidMaxSize(max int) MaxSize {
	return boovalidation.ValidMaxSize(max)
}

// Length requires an array or string to be exactly a given length.
type Length = boovalidation.Length

func ValidLength(n int) Length {
	return boovalidation.ValidLength(n)
}

// Match requires a string to match a given regex.
type Match = boovalidation.Match

func ValidMatch(regex *regexp.Regexp) Match {
	return boovalidation.ValidMatch(regex)
}

type Email = boovalidation.Email

func ValidEmail() Email {
	return boovalidation.ValidEmail()
}

const (
	None               = boovalidation.None
	IPAny              = boovalidation.IPAny
	IPv4               = boovalidation.IPv4
	IPv6               = boovalidation.IPv6
	IPv4MappedIPv6     = boovalidation.IPv4MappedIPv6
	IPv4CIDR           = boovalidation.IPv4CIDR
	IPv6CIDR           = boovalidation.IPv6CIDR
	IPv4MappedIPv6CIDR = boovalidation.IPv4MappedIPv6CIDR
)

// Requires a string(IP Address) to be within IP Pattern type inclusive.
type IPAddr = boovalidation.IPAddr

// Requires an IP Address string to be exactly a given  validation type (IPv4, IPv6, IPv4MappedIPv6, IPv4CIDR, IPv6CIDR, IPv4MappedIPv6CIDR OR IPAny)
func ValidIPAddr(cktypes ...int) IPAddr {
	return boovalidation.ValidIPAddr(cktypes...)
}

// Requires a MAC Address string to be exactly
type MacAddr = boovalidation.MacAddr

func ValidMacAddr() MacAddr {
	return boovalidation.ValidMacAddr()
}

// Requires a Domain string to be exactly
type Domain = boovalidation.Domain

func ValidDomain() Domain {
	return boovalidation.ValidDomain()
}

type URL = boovalidation.URL

func ValidURL() URL {
	return boovalidation.ValidURL()
}

/*
NORMAL BenchmarkRegex-8   	2000000000	         0.24 ns/op
STRICT BenchmarkLoop-8    	2000000000	         0.01 ns/op
*/
const (
	NORMAL = boovalidation.NORMAL
	STRICT = boovalidation.STRICT
)

// Requires a string to be without invisible characters
type PureText = boovalidation.PureText

func ValidPureText(m int) PureText {
	return boovalidation.ValidPureText(m)
}

const (
	ONLY_FILENAME       = boovalidation.ONLY_FILENAME
	ALLOW_RELATIVE_PATH = boovalidation.ALLOW_RELATIVE_PATH
)

// Requires an string to be sanitary file path
type FilePath = boovalidation.FilePath

func ValidFilePath(m int) FilePath {
	return boovalidation.ValidFilePath(m)
}

// 开始日期不能大于结束日期
type TimeStartEndCheck = boovalidation.TimeStartEndCheck
