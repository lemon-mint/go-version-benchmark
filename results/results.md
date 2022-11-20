# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.2

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
| 1.17 | 474.870029 | 21.602368 | 3859.987942 | 223.926194 |
| 1.18 | 486.928339 | 66.378958 | 3516.806293 | 3.499835 |
| 1.19 | 542.256106 | 32.299729 | 3540.429836 | 198.144210 |
| gotip | 3779.533715 | 86.930833 | 3555.545980 | 1.338089 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 128.701850 | 16.842076 | 3901.959618 | 1.548856 |
| 1.18 | 131.555685 | 23.514462 | 4012.011386 | 2.285213 |
| 1.19 | 139.003615 | 3.870593 | 3672.251912 | 229.630306 |
| gotip | 2818.146371 | 71.356612 | 3907.472324 | 185.607884 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 131.582377 | 3.101130 | 5628.173505 | 154.898966 |
| 1.18 | 145.172938 | 32.174134 | 5662.646870 | 251.265120 |
| 1.19 | 143.584216 | 4.017813 | 5432.319009 | 111.269951 |
| gotip | 2793.050826 | 114.179747 | 6149.433495 | 293.529472 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 133.290554 | 5.651492 | 8917.948904 | 1745.492480 |
| 1.18 | 143.290951 | 6.433760 | 5186.747454 | 129.014954 |
| 1.19 | 145.188284 | 3.175179 | 8399.942293 | 40.421673 |
| gotip | 2895.098240 | 99.731769 | 8479.785924 | 133.828262 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 132.858538 | 5.178960 | 6026.120289 | 369.037644 |
| 1.18 | 139.495662 | 45.024295 | 5868.485054 | 76.720994 |
| 1.19 | 150.210138 | 7.035782 | 7591.078111 | 198.302333 |
| gotip | 2763.699113 | 60.509293 | 8888.657203 | 388.477479 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 152.279430 | 5.671119 | 5011.658435 | 270.124302 |
| 1.18 | 157.136666 | 5.688010 | 5032.690291 | 271.419431 |
| 1.19 | 150.536706 | 9.268066 | 5086.095370 | 151.975866 |
| gotip | 3255.443411 | 230.389966 | 5080.364632 | 3060.372012 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 132.063513 | 6.594903 | 3278.530233 | 199.130660 |
| 1.18 | 143.027496 | 6.856627 | 3577.309876 | 126.797682 |
| 1.19 | 147.280528 | 5.471833 | 369.550911 | 21.123692 |
| gotip | 3084.085734 | 89.100362 | 348.703402 | 21.437667 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 154.047962 | 74.478616 | 5181.913642 | 512.567491 |
| 1.18 | 159.377827 | 80.131370 | 4686.778627 | 732.741218 |
| 1.19 | 161.833100 | 43.766376 | 4892.720581 | 552.393347 |
| gotip | 3830.970912 | 177.755388 | 5835.567340 | 428.063799 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 128.709250 | 11.202412 | 3131.223296 | 84.032043 |
| 1.18 | 134.488649 | 4.975354 | 4347.948735 | 153.782426 |
| 1.19 | 146.222048 | 15.889284 | 2191.079427 | 77.008605 |
| gotip | 2787.194908 | 143.813122 | 2190.665912 | 75.532098 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 165.396866 | 21.700876 | 1051.078301 | 30.945761 |
| 1.18 | 167.804449 | 13.691917 | 1119.045397 | 24.948174 |
| 1.19 | 173.663008 | 2.088825 | 1151.663871 | 29.803105 |
| gotip | 3912.485835 | 59.870224 | 1167.738117 | 56.016660 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 173.951860 | 6.060478 | 2322.557540 | 62.109285 |
| 1.18 | 174.090119 | 7.287244 | 2313.001178 | 40.730593 |
| 1.19 | 177.064435 | 3.591173 | 2213.806635 | 203.759941 |
| gotip | 4000.771222 | 97.144860 | 2173.845772 | 97.172523 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 200.700364 | 17.256619 | 5204.246192 | 1.942461 |
| 1.18 | 210.181345 | 11.943105 | 5723.590066 | 2.733419 |
| 1.19 | 205.558015 | 12.180313 | 3221.501244 | 185.711091 |
| gotip | 5237.901563 | 364.689478 | 2944.792860 | 180.144885 |

![switch_case](./switch_case__725e73000e.png)

