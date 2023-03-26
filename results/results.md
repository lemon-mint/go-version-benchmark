# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.7

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
| 1.18 | 462.138072 | 4.525253 | 3517.131603 | 2.005346 |
| 1.19 | 482.068686 | 3.628019 | 3542.262798 | 1.661628 |
| 1.20 | 413.966975 | 4.056053 | 3554.560312 | 1.233804 |
| gotip | 449.461490 | 4.274173 | 3725.104505 | 50.977751 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 114.168036 | 12.837385 | 4012.601201 | 1.721489 |
| 1.19 | 106.726415 | 1.875989 | 3902.365043 | 1.493045 |
| 1.20 | 115.163461 | 2.512877 | 3902.912554 | 6.161583 |
| gotip | 115.055859 | 3.315745 | 4249.686502 | 1.291344 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 126.906735 | 8.828402 | 5710.721153 | 15.267878 |
| 1.19 | 106.809665 | 3.671943 | 5574.385517 | 28.035732 |
| 1.20 | 114.132567 | 3.518783 | 6222.427279 | 25.081538 |
| gotip | 116.077283 | 3.130282 | 6579.621786 | 35.323459 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 129.045017 | 2.923824 | 5074.076401 | 49.296659 |
| 1.19 | 110.122323 | 3.196975 | 8630.087535 | 34.132940 |
| 1.20 | 117.352260 | 3.581513 | 8699.945397 | 49.126889 |
| gotip | 116.764463 | 1.822656 | 5205.471316 | 22.221548 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 122.349700 | 2.370070 | 5733.539587 | 51.298159 |
| 1.19 | 113.127908 | 4.658997 | 7416.221857 | 60.695122 |
| 1.20 | 117.913204 | 2.504098 | 8565.667722 | 99.492750 |
| gotip | 120.896884 | 5.708683 | 8770.811943 | 123.265954 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 136.489005 | 6.145301 | 5028.732298 | 1.305823 |
| 1.19 | 120.888791 | 3.824195 | 5083.795984 | 3.596025 |
| 1.20 | 126.605903 | 2.438059 | 5079.113723 | 1.796540 |
| gotip | 130.305632 | 4.693558 | 5263.743140 | 3.821399 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 128.872260 | 2.935246 | 3574.733295 | 0.765739 |
| 1.19 | 112.040664 | 2.472908 | 369.885577 | 0.735011 |
| 1.20 | 118.497663 | 2.210929 | 369.721769 | 1.485413 |
| gotip | 120.579868 | 4.114896 | 374.622941 | 0.376068 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 138.137497 | 3.911397 | 4983.970643 | 756.267260 |
| 1.19 | 125.774173 | 21.792142 | 4861.140723 | 371.300822 |
| 1.20 | 127.659596 | 6.261566 | 5255.177719 | 182.670845 |
| gotip | 134.061534 | 4.515271 | 5129.949571 | 236.564049 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 117.678902 | 3.601337 | 4348.167731 | 1.568070 |
| 1.19 | 110.455419 | 3.855970 | 2191.700346 | 0.709604 |
| 1.20 | 116.296754 | 4.097572 | 2191.671008 | 2.124020 |
| gotip | 120.745220 | 3.852093 | 2191.122040 | 2.039189 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 151.797251 | 19.062685 | 1147.780200 | 8.084109 |
| 1.19 | 134.718576 | 6.133567 | 1156.113345 | 3.502535 |
| 1.20 | 134.642562 | 3.616540 | 1155.270763 | 5.070945 |
| gotip | 139.870375 | 3.742091 | 1133.383411 | 6.937732 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 156.437547 | 4.757497 | 2357.975151 | 12.206130 |
| 1.19 | 139.466472 | 3.294822 | 2243.206324 | 19.858688 |
| 1.20 | 145.056887 | 6.777366 | 2208.246453 | 10.038769 |
| gotip | 145.130153 | 4.192944 | 2213.933562 | 11.523012 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 186.150484 | 24.997601 | 5727.258543 | 3.936673 |
| 1.19 | 161.285167 | 10.570044 | 3223.595425 | 2.094798 |
| 1.20 | 165.009792 | 4.748380 | 3223.656923 | 1.157641 |
| gotip | 169.312952 | 3.131271 | 2690.430777 | 2.045793 |

![switch_case](./switch_case__725e73000e.png)

