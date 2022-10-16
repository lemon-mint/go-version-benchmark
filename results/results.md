# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.1

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.906000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.906000

CacheSize: 36608

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 491.085032 | 20.537121 | 3594.411672 | 221.404895 |
| 1.18 | 498.510829 | 13.963980 | 3108.162253 | 203.216146 |
| 1.19 | 536.976098 | 12.219373 | 3459.078884 | 210.745115 |
| gotip | 537.128204 | 11.839381 | 3143.776248 | 205.765388 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 127.669747 | 10.850481 | 3445.733385 | 201.268353 |
| 1.18 | 132.868683 | 4.077856 | 3779.510458 | 236.389243 |
| 1.19 | 144.097309 | 5.267731 | 3439.100165 | 231.050192 |
| gotip | 146.467338 | 3.462997 | 3445.160126 | 146.320631 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 134.514894 | 4.284949 | 5386.982578 | 94.554115 |
| 1.18 | 145.178286 | 4.235005 | 5392.485118 | 90.644774 |
| 1.19 | 140.483120 | 7.176659 | 4948.871444 | 138.798362 |
| gotip | 148.254667 | 4.467487 | 5819.675007 | 127.936873 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 135.726608 | 2.925329 | 8710.819842 | 146.278573 |
| 1.18 | 147.849960 | 3.598848 | 5274.640507 | 85.796164 |
| 1.19 | 149.416258 | 4.959958 | 8810.554946 | 97.162229 |
| gotip | 160.723855 | 5.706408 | 8840.150432 | 85.817939 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 128.424658 | 4.326448 | 6077.879322 | 135.589002 |
| 1.18 | 144.679844 | 3.288771 | 6170.401900 | 138.132789 |
| 1.19 | 153.194509 | 5.257591 | 7836.347781 | 269.769733 |
| gotip | 165.998825 | 6.460679 | 8962.349111 | 343.041038 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 144.382103 | 2.751398 | 5016.001500 | 238.013675 |
| 1.18 | 160.147801 | 6.909331 | 4440.042937 | 269.109196 |
| 1.19 | 145.960823 | 6.503874 | 4488.898992 | 239.548602 |
| gotip | 158.607876 | 2.835313 | 4494.133091 | 275.827616 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 126.466864 | 3.740319 | 3277.126986 | 131.323570 |
| 1.18 | 145.089664 | 3.498857 | 3301.860868 | 202.088781 |
| 1.19 | 144.731957 | 2.372924 | 328.090595 | 21.163117 |
| gotip | 145.317914 | 3.551555 | 327.472687 | 0.927524 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 157.349462 | 5.326098 | 5286.208559 | 972.006034 |
| 1.18 | 164.894351 | 104.336723 | 4413.420434 | 490.162963 |
| 1.19 | 156.406789 | 112.245879 | 5124.253121 | 394.580580 |
| gotip | 170.350317 | 5.879065 | 5556.167208 | 617.750747 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 130.008045 | 5.015596 | 2970.254999 | 60.051934 |
| 1.18 | 137.146200 | 4.445310 | 4346.725507 | 251.127522 |
| 1.19 | 147.272753 | 5.951506 | 2192.426204 | 118.308595 |
| gotip | 166.280158 | 3.804763 | 2195.616036 | 115.820826 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 165.723467 | 17.446069 | 1047.969888 | 18.712453 |
| 1.18 | 166.252028 | 14.408463 | 1091.360143 | 25.197891 |
| 1.19 | 176.744603 | 12.907556 | 1107.158676 | 17.794914 |
| gotip | 186.755132 | 6.667626 | 1116.136094 | 23.217473 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 175.772199 | 3.819842 | 2158.846196 | 81.191611 |
| 1.18 | 174.220180 | 5.509912 | 2244.432246 | 41.828470 |
| 1.19 | 178.172130 | 6.207125 | 2129.308974 | 35.315345 |
| gotip | 179.556644 | 4.178200 | 2111.653232 | 27.279220 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 192.845709 | 18.618945 | 4596.046977 | 183.064539 |
| 1.18 | 205.690742 | 22.801422 | 5053.085139 | 329.738992 |
| 1.19 | 206.539166 | 20.277896 | 2847.713439 | 152.178189 |
| gotip | 216.546670 | 4.690771 | 2891.569556 | 114.978536 |

![switch_case](./switch_case__725e73000e.png)

