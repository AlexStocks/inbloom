/******************************************************
# DESC    :
# AUTHOR  : Alex Stocks
# VERSION : 1.0
# LICENCE : Apache License 2.0
# EMAIL   : alexstocks@foxmail.com
# MOD     : 2017-03-04 17:48
# FILE    : a.go
******************************************************/

package main

import (
	"encoding/base64"
	"fmt"
	"github.com/AlexStocks/inbloom/go/inbloom"
)

func main() {
	data := "yoID6AAAABMNJALf42lALcpcu2WH9sZcLPWh/g+ynjcSVaWfDxVuudTRAA=="
	bf, err := inbloom.UnmarshalBase64(data)
	fmt.Printf("%v %v\n", bf, err)

	uids := []string{"u0667601477730140020006032",
		"u0824011478254848030001156", "u1768771480923934030001199",
		"u2413521474991179020001113", "u3686801478240154030005408",
		"u3757851481174023020001166", "u3795351482317225020001244",
		"u4555881480912962010001166", "u4710551480917824010001244",
		"u4881041468996697020001184", "u4881041474530826010001095",
		"u4905391484641165010001141", "u5425051481012754020001255",
		"u6450591484584368010001166", "u6937721484632783020001218",
		"u7414411480402599030001267", "u7675311485144153020001211",
		"u9587291477909978030001277", "u9805521470914539020001149"}
	for _, uid := range uids {
		fmt.Println(bf.Contains(uid))
	}
	fmt.Println(bf.Contains("u0"))

	//  stdEncoding
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Printf("bloom: could not decode base64 data: %s", err)
		return
	}

	bf, err = inbloom.Unmarshal(b)
	fmt.Printf("%v %v\n", bf, err)
	for _, uid := range uids {
		fmt.Println(bf.Contains(uid))
	}
	fmt.Println(bf.Contains("u0"))

	//  URLEncoding
	// 这个方式是EverytingMe/inbloom 的inbloom.UnmarshalBase64官方实现，
        // 无法解析bloom.py通过base64编码的结果，我已经在我自己的fork
        // (github.com/AlexStocks/inbloom)里面fix掉了这个bug
	_, err = base64.URLEncoding.DecodeString(data)
	if err != nil {
		// bloom: could not decode base64 data: illegal base64 data at input byte 36Compile over.
		fmt.Printf("bloom: could not decode base64 data: %s", err)
		return
	}
}
