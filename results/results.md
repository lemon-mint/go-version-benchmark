# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.905000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.905000

CacheSize: 36608

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 503.754753 | 13.938978 | 3863.358460 | 4.902641 |
| 1.18 | 512.588496 | 9.497590 | 3516.503350 | 45.635377 |
| 1.19 | 538.397280 | 10.676108 | 3541.103983 | 0.942954 |
| gotip | 544.191253 | 4.136828 | 3534.186269 | 2.905709 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 125.138878 | 3.087323 | 3902.241977 | 3.528802 |
| 1.18 | 137.512056 | 4.794827 | 4014.325385 | 2.220080 |
| 1.19 | 145.232079 | 11.382633 | 3901.622409 | 2.658473 |
| gotip | 146.469040 | 2.695426 | 4012.668872 | 4.230537 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 133.640038 | 4.935518 | 5688.783888 | 30.867588 |
| 1.18 | 140.703289 | 2.727336 | 5716.067733 | 33.057779 |
| 1.19 | 143.956698 | 3.446320 | 5544.484780 | 20.863484 |
| gotip | 154.923910 | 8.621602 | 6343.160792 | 17.847290 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 137.499863 | 1.920304 | 8975.273769 | 56.859465 |
| 1.18 | 143.793678 | 6.141765 | 5352.833731 | 75.993013 |
| 1.19 | 149.986795 | 2.334586 | 8971.260987 | 42.781823 |
| gotip | 152.051139 | 3.704606 | 5481.027398 | 69.676283 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 131.029024 | 2.319208 | 6107.201673 | 102.626522 |
| 1.18 | 142.386706 | 2.900123 | 6091.999942 | 101.588678 |
| 1.19 | 152.762361 | 6.902013 | 7802.223402 | 143.923391 |
| gotip | 155.890771 | 5.238942 | 8999.670233 | 130.455710 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 146.990778 | 1.509434 | 5012.333733 | 1.342012 |
| 1.18 | 150.838891 | 5.098312 | 5028.935843 | 6.269215 |
| 1.19 | 153.684849 | 3.246817 | 5082.406897 | 3.593122 |
| gotip | 158.631855 | 3.468104 | 5036.356031 | 2.246805 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 136.746105 | 4.571940 | 3712.445741 | 2.928765 |
| 1.18 | 143.812503 | 1.846923 | 3575.945276 | 1.599698 |
| 1.19 | 148.285897 | 2.152570 | 369.743565 | 0.474602 |
| gotip | 156.694140 | 2.738440 | 369.583630 | 0.551071 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 161.786031 | 12.959678 | 5025.864965 | 659.818003 |
| 1.18 | 172.199596 | 68.435858 | 4687.232057 | 359.140874 |
| 1.19 | 173.039129 | 125.818364 | 4766.643765 | 509.691623 |
| gotip | 178.062181 | 4.257405 | 5423.856409 | 260.338988 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 132.139678 | 5.277601 | 3231.369856 | 23.931399 |
| 1.18 | 140.221517 | 4.587086 | 4352.380279 | 2.032632 |
| 1.19 | 155.812154 | 3.650121 | 2193.742586 | 1.550067 |
| gotip | 155.121488 | 2.820576 | 4348.470283 | 2.534466 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 163.062690 | 5.825802 | 1091.742841 | 10.485347 |
| 1.18 | 168.724429 | 12.170024 | 1147.632398 | 5.439022 |
| 1.19 | 174.256749 | 15.639540 | 1168.014189 | 5.535409 |
| gotip | 177.360038 | 2.617740 | 1073.204939 | 8.673456 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 169.524923 | 3.738427 | 2360.019580 | 7.914429 |
| 1.18 | 174.356679 | 2.203753 | 2368.254270 | 12.157428 |
| 1.19 | 179.329855 | 3.101397 | 2252.070805 | 11.483271 |
| gotip | 181.491367 | 4.419620 | 2222.642173 | 14.916601 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 195.482705 | 3.312234 | 5205.500030 | 5.823247 |
| 1.18 | 211.625101 | 14.180559 | 5724.474749 | 1.874178 |
| 1.19 | 203.795471 | 54.408569 | 3223.398820 | 1.633174 |
| gotip | 219.186163 | 2.853912 | 3223.570135 | 1.056925 |

![switch_case](./switch_case__725e73000e.png)

