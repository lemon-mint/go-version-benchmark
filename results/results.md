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
| 1.18 | 492.930214 | 7.877782 | 3521.012322 | 2.327470 |
| 1.19 | 514.464645 | 5.006468 | 3544.727878 | 0.866317 |
| 1.20 | 434.907869 | 6.335543 | 3556.365292 | 1.764489 |
| gotip | 477.910103 | 4.791412 | 3537.195052 | 3.515493 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 124.996685 | 2.407519 | 4014.854522 | 0.992227 |
| 1.19 | 113.840709 | 3.013365 | 3905.959963 | 2.639560 |
| 1.20 | 118.160457 | 2.310455 | 3920.036085 | 9.272445 |
| gotip | 122.525834 | 2.887904 | 4248.764714 | 2.438096 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 133.237343 | 5.978704 | 5735.214468 | 15.693838 |
| 1.19 | 117.109935 | 8.313065 | 5571.896033 | 25.678557 |
| 1.20 | 118.676445 | 3.587387 | 6224.805726 | 24.482646 |
| gotip | 120.549313 | 5.040075 | 6571.438625 | 35.240479 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 134.224794 | 4.879955 | 5673.136099 | 69.160576 |
| 1.19 | 118.612843 | 3.671816 | 9241.478967 | 26.539340 |
| 1.20 | 122.483318 | 3.155709 | 9086.027382 | 21.566729 |
| gotip | 119.501901 | 2.241801 | 5746.487320 | 55.003162 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 126.189355 | 3.372008 | 6201.331430 | 71.887707 |
| 1.19 | 115.902276 | 11.319537 | 7902.996544 | 123.208017 |
| 1.20 | 121.716167 | 3.085062 | 9302.784086 | 99.733630 |
| gotip | 125.251138 | 4.240681 | 10408.581113 | 269.523072 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 146.349931 | 2.982704 | 5034.005115 | 2.628625 |
| 1.19 | 123.241175 | 2.544844 | 5087.647329 | 6.464964 |
| 1.20 | 128.072744 | 4.319928 | 5084.504378 | 1.953321 |
| gotip | 134.421200 | 2.857192 | 5134.703981 | 2.621392 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 134.615197 | 2.329466 | 3578.139189 | 9.296613 |
| 1.19 | 114.950755 | 4.511059 | 370.769985 | 0.507158 |
| 1.20 | 126.839245 | 3.014159 | 370.177683 | 0.481483 |
| gotip | 125.843293 | 2.595672 | 374.937818 | 0.518874 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 153.114224 | 9.284860 | 4771.139992 | 513.197047 |
| 1.19 | 128.758911 | 25.311056 | 4919.214084 | 342.312914 |
| 1.20 | 131.092380 | 4.392126 | 5372.586563 | 293.450840 |
| gotip | 134.570251 | 3.001073 | 5458.993498 | 675.155666 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 122.719021 | 0.941564 | 4348.199965 | 0.872537 |
| 1.19 | 112.042464 | 12.840298 | 2193.001922 | 1.480971 |
| 1.20 | 120.736782 | 3.640212 | 2192.857275 | 1.547860 |
| gotip | 122.526522 | 1.808346 | 4351.026299 | 1.029251 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 170.556343 | 8.027336 | 1222.961095 | 20.160634 |
| 1.19 | 143.054209 | 8.405238 | 1234.236869 | 10.733635 |
| 1.20 | 147.360697 | 4.222501 | 1219.534962 | 4.920251 |
| gotip | 143.656213 | 2.184271 | 1173.537102 | 5.457793 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 168.914332 | 6.356744 | 2415.441098 | 8.958059 |
| 1.19 | 146.334130 | 3.560395 | 2299.918341 | 10.386557 |
| 1.20 | 151.305064 | 3.266148 | 2269.324938 | 15.068381 |
| gotip | 151.430965 | 3.967621 | 2203.973427 | 10.677632 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 195.240362 | 6.153019 | 5841.743813 | 3.332233 |
| 1.19 | 168.493580 | 5.543553 | 3223.771276 | 4.587595 |
| 1.20 | 177.386882 | 4.437837 | 3221.964754 | 0.970894 |
| gotip | 171.560345 | 4.723066 | 2775.103583 | 2.582676 |

![switch_case](./switch_case__725e73000e.png)

