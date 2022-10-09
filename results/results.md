# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.1

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz

Cores: 1

Mhz: 2095.078000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz

Cores: 1

Mhz: 2095.078000

CacheSize: 36608

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 569.565417 | 12.351995 | 4558.310779 | 18.068102 |
| 1.18 | 572.406066 | 15.204457 | 4153.388928 | 11.569724 |
| 1.19 | 635.080614 | 14.639801 | 4189.987537 | 19.162114 |
| gotip | 644.619585 | 13.747400 | 4188.183850 | 18.736524 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 148.540154 | 5.713700 | 4618.623024 | 16.452204 |
| 1.18 | 154.188057 | 3.559823 | 4724.672952 | 30.924625 |
| 1.19 | 170.107649 | 5.233367 | 4388.250439 | 51.086089 |
| gotip | 169.039324 | 3.970700 | 4534.856837 | 80.050607 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 149.650332 | 6.209579 | 6482.604709 | 96.861931 |
| 1.18 | 162.289762 | 8.935545 | 6626.603552 | 81.190797 |
| 1.19 | 165.853324 | 6.008219 | 6494.379000 | 75.095175 |
| gotip | 180.831427 | 5.449354 | 7574.955480 | 34.263940 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 168.829316 | 8.170694 | 10580.247733 | 98.658545 |
| 1.18 | 170.663888 | 3.470699 | 6207.893620 | 103.982018 |
| 1.19 | 183.492710 | 10.494234 | 10423.951632 | 76.742531 |
| gotip | 184.629047 | 5.868704 | 6281.930892 | 131.396890 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 153.318935 | 4.088900 | 7049.846368 | 286.282248 |
| 1.18 | 159.241736 | 6.070519 | 6690.512862 | 85.204792 |
| 1.19 | 171.002250 | 7.238165 | 8414.350056 | 113.237232 |
| gotip | 172.132767 | 5.261241 | 10204.065680 | 124.482374 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 170.136845 | 3.091820 | 5609.193096 | 62.837260 |
| 1.18 | 172.970849 | 7.105060 | 5669.535599 | 52.434803 |
| 1.19 | 178.885163 | 3.591308 | 5693.273822 | 61.735417 |
| gotip | 178.362940 | 7.954665 | 5631.739777 | 85.392521 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 154.411394 | 6.483191 | 4272.122420 | 117.344833 |
| 1.18 | 168.155148 | 3.587531 | 3989.797344 | 73.734320 |
| 1.19 | 164.252326 | 3.832213 | 416.208637 | 9.417645 |
| gotip | 171.617707 | 5.479753 | 409.772459 | 15.152231 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 178.046672 | 17.772700 | 6744.453183 | 923.818309 |
| 1.18 | 185.432277 | 101.210068 | 4465.643599 | 475.288965 |
| 1.19 | 188.322420 | 140.751062 | 4833.762383 | 658.115907 |
| gotip | 196.369761 | 6.787073 | 5607.644073 | 625.836538 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 147.224194 | 5.277992 | 3790.062939 | 29.641970 |
| 1.18 | 156.557207 | 5.929060 | 5192.813084 | 12.729392 |
| 1.19 | 169.768168 | 5.188187 | 2599.762725 | 19.677210 |
| gotip | 173.154372 | 7.179902 | 5143.793158 | 22.606743 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 186.291506 | 15.289126 | 1260.909860 | 23.646704 |
| 1.18 | 196.066064 | 14.593912 | 1329.729744 | 38.483561 |
| 1.19 | 202.676543 | 27.774346 | 1390.195119 | 21.252019 |
| gotip | 199.963915 | 4.241482 | 1315.245072 | 42.501954 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 193.752702 | 7.951096 | 2660.203503 | 65.307543 |
| 1.18 | 197.786576 | 7.410867 | 2683.519595 | 42.055348 |
| 1.19 | 197.872501 | 5.222135 | 2565.656157 | 39.776191 |
| gotip | 215.350176 | 8.794100 | 2511.943874 | 41.040868 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 219.869649 | 15.024590 | 5735.609253 | 97.103772 |
| 1.18 | 225.961342 | 24.010285 | 6361.644915 | 73.498629 |
| 1.19 | 226.811620 | 25.050975 | 3567.350094 | 41.560687 |
| gotip | 240.331021 | 13.865497 | 3618.802988 | 50.549375 |

![switch_case](./switch_case__725e73000e.png)

