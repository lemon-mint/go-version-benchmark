# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.439000

CacheSize: 49152

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.439000

CacheSize: 49152

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 457.514487 | 9.259884 | 4018.101221 | 66.118316 |
| 1.18 | 463.513854 | 5.795732 | 3614.512380 | 398.809538 |
| 1.19 | 518.530357 | 6.728668 | 3610.936982 | 2.358347 |
| gotip | 525.754184 | 2.460560 | 3679.255563 | 2.046898 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 121.746875 | 8.857897 | 4791.566249 | 73.502612 |
| 1.18 | 124.415403 | 4.088408 | 4795.328954 | 0.897467 |
| 1.19 | 136.585852 | 3.954437 | 4791.064466 | 1.899340 |
| gotip | 141.047475 | 4.079929 | 4795.477082 | 1.150984 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 128.672653 | 3.001794 | 5154.554748 | 69.738923 |
| 1.18 | 132.666968 | 3.131553 | 5174.080848 | 15.803293 |
| 1.19 | 135.218788 | 10.914565 | 5191.724120 | 26.116164 |
| gotip | 141.964365 | 2.690186 | 6261.306264 | 261.416633 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 130.330815 | 2.956012 | 9765.781406 | 34.236741 |
| 1.18 | 136.539579 | 6.340158 | 5611.894866 | 28.110606 |
| 1.19 | 140.130390 | 3.496280 | 9791.739374 | 33.467175 |
| gotip | 144.070894 | 3.137060 | 5775.764779 | 18.384345 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 124.709203 | 4.135441 | 5864.323365 | 41.610718 |
| 1.18 | 129.016256 | 3.902622 | 5549.484014 | 37.108434 |
| 1.19 | 142.735609 | 4.588747 | 7118.150429 | 24.933606 |
| gotip | 145.955878 | 4.036194 | 7738.693213 | 53.834861 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 138.808582 | 2.308015 | 6067.032063 | 2.962583 |
| 1.18 | 147.406145 | 2.680759 | 5977.993939 | 10.933014 |
| 1.19 | 148.366121 | 2.450550 | 5925.753273 | 2.311500 |
| gotip | 154.150367 | 1.832849 | 5955.581767 | 9.642318 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 130.890134 | 2.288513 | 3627.656012 | 273.738432 |
| 1.18 | 136.251260 | 4.346134 | 3536.910399 | 2.441708 |
| 1.19 | 139.991432 | 3.430639 | 360.367648 | 0.492238 |
| gotip | 147.676906 | 2.939258 | 393.837650 | 3.782604 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 142.947675 | 13.972591 | 4810.085559 | 548.225605 |
| 1.18 | 149.677236 | 40.088997 | 5313.993022 | 634.889682 |
| 1.19 | 154.586542 | 81.944962 | 5013.122950 | 357.156981 |
| gotip | 161.422415 | 6.931243 | 5163.202784 | 364.080707 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 120.093260 | 7.688799 | 3668.748058 | 10.638788 |
| 1.18 | 126.617041 | 3.843801 | 5216.023737 | 2.110905 |
| 1.19 | 139.594644 | 1.623054 | 2705.191279 | 3.965979 |
| gotip | 141.334266 | 2.569868 | 5221.132314 | 0.912754 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 157.295437 | 9.611959 | 1064.274475 | 14.778023 |
| 1.18 | 158.979396 | 40.220505 | 1105.796997 | 5.404813 |
| 1.19 | 165.854863 | 11.743166 | 1142.499159 | 104.193798 |
| gotip | 169.530107 | 3.533650 | 1047.592943 | 4.607674 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 160.591643 | 1.772116 | 2323.699525 | 12.948628 |
| 1.18 | 168.011518 | 3.573946 | 2322.788583 | 29.673351 |
| 1.19 | 169.194112 | 2.574064 | 2204.172897 | 18.743937 |
| gotip | 174.993698 | 4.317360 | 2152.898417 | 13.765400 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 189.359573 | 10.204875 | 5934.925420 | 37.166981 |
| 1.18 | 190.886099 | 10.622505 | 5974.034973 | 23.021938 |
| 1.19 | 196.391139 | 10.438511 | 3505.634649 | 3.138923 |
| gotip | 206.018986 | 2.769974 | 3317.766805 | 1.806565 |

![switch_case](./switch_case__725e73000e.png)

