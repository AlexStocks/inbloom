#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# ******************************************************
# DESC    :
# AUTHOR  : Alex Stocks
# VERSION : 1.0
# LICENCE : Apache License 2.0
# EMAIL   : alexstocks@foxmail.com
# MOD     : 2017-03-03 17:22
# FILE    : bloom.py
# ******************************************************

import inbloom
import base64
import requests

# Basic usage
uids = ['u0667601477730140020006032', 'u0824011478254848030001156', 'u1768771480923934030001199', 'u2413521474991179020001113', 'u3686801478240154030005408', 'u3757851481174023020001166', 'u3795351482317225020001244', 'u4555881480912962010001166', 'u4710551480917824010001244', 'u4881041468996697020001184', 'u4881041474530826010001095', 'u4905391484641165010001141', 'u5425051481012754020001255', 'u6450591484584368010001166', 'u6937721484632783020001218', 'u7414411480402599030001267', 'u7675311485144153020001211', 'u9587291477909978030001277', 'u9805521470914539020001149']
bf = inbloom.Filter(entries=len(uids), error=0.001)
for uid in uids:
    bf.add(uid)

res = base64.b64encode(inbloom.dump(bf))
# yoID6AAAABMNJALf42lALcpcu2WH9sZcLPWh/g+ynjcSVaWfDxVuudTRAA==
print len(bf.buffer())
print res, len(res)
bf = inbloom.load(base64.b64decode(res))
print bf.contains(uids[0])
print bf.contains('u0')
