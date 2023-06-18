# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.10

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.907000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.907000

CacheSize: 36608

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 481.869464 | 23.439378 | 3518.293828 | 51.163055 |
| 1.19 | 502.699200 | 4.744232 | 3542.909107 | 1.809980 |
| 1.20 | 431.797677 | 5.367756 | 3555.586906 | 2.684452 |
| gotip | 472.106668 | 6.264536 | 3535.111486 | 4.155169 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 122.560876 | 2.895956 | 4014.937591 | 2.454683 |
| 1.19 | 109.632747 | 3.994709 | 3904.461815 | 1.971809 |
| 1.20 | 118.063625 | 3.186822 | 3909.191502 | 5.937154 |
| gotip | 117.291212 | 1.960158 | 4246.949867 | 1.217338 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 130.048133 | 6.148545 | 5732.741807 | 25.600471 |
| 1.19 | 112.240891 | 4.271006 | 5578.791443 | 21.401902 |
| 1.20 | 123.438990 | 4.996583 | 6234.629070 | 25.432913 |
| gotip | 117.905715 | 4.483825 | 6558.230896 | 22.952999 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 135.084542 | 14.906339 | 5504.455404 | 47.018594 |
| 1.19 | 114.257270 | 2.261227 | 9156.287014 | 40.116846 |
| 1.20 | 127.262197 | 4.659488 | 9256.589843 | 84.625434 |
| gotip | 120.550755 | 3.672143 | 5654.917214 | 109.988508 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 128.059506 | 2.959457 | 6104.922211 | 77.665078 |
| 1.19 | 115.438432 | 1.358644 | 7784.840936 | 67.758259 |
| 1.20 | 123.968249 | 3.880355 | 8891.685833 | 43.351359 |
| gotip | 119.208983 | 5.161161 | 9267.397614 | 97.911077 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 144.200848 | 2.398362 | 5028.433442 | 1.789305 |
| 1.19 | 122.051020 | 10.070570 | 5090.288178 | 3.830824 |
| 1.20 | 130.641459 | 3.594773 | 5082.159727 | 2.069934 |
| gotip | 132.440959 | 4.112221 | 5136.005288 | 3.632331 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 133.035365 | 3.493062 | 3576.868063 | 1.813638 |
| 1.19 | 116.184369 | 7.157695 | 370.260819 | 0.467335 |
| 1.20 | 122.047924 | 4.159776 | 370.166355 | 0.541969 |
| gotip | 119.775089 | 4.451633 | 375.183406 | 0.374535 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 146.179656 | 8.939989 | 4102.965301 | 333.878769 |
| 1.19 | 123.764348 | 11.307547 | 4620.366287 | 843.523515 |
| 1.20 | 134.744498 | 7.977020 | 5323.695855 | 300.015150 |
| gotip | 130.814043 | 4.358005 | 5507.302830 | 275.443213 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 123.501891 | 7.371750 | 4348.323938 | 0.624626 |
| 1.19 | 112.930248 | 10.777563 | 2191.609607 | 0.920169 |
| 1.20 | 120.171556 | 3.061751 | 2193.969035 | 0.963512 |
| gotip | 119.749940 | 3.010273 | 4348.867286 | 1.917034 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 159.887684 | 6.162474 | 1193.609956 | 16.964325 |
| 1.19 | 137.262259 | 6.362798 | 1220.898182 | 5.181350 |
| 1.20 | 148.344370 | 5.101356 | 1200.490143 | 9.234333 |
| gotip | 143.622319 | 5.498261 | 1175.545136 | 33.790270 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 168.089215 | 15.310503 | 2413.186691 | 16.184724 |
| 1.19 | 143.209409 | 3.297956 | 2299.751277 | 5.185477 |
| 1.20 | 154.671666 | 4.873997 | 2268.821506 | 7.932549 |
| gotip | 147.213466 | 3.512388 | 2198.096421 | 7.359966 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 195.752100 | 7.495450 | 5842.087757 | 4.707965 |
| 1.19 | 162.473766 | 12.501259 | 3221.111279 | 1.826855 |
| 1.20 | 172.921425 | 5.308119 | 3219.828604 | 1.824030 |
| gotip | 169.022965 | 5.484233 | 2774.562259 | 2.148510 |

![switch_case](./switch_case__725e73000e.png)

