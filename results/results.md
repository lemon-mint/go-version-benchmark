# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.2

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.437000

CacheSize: 49152

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.437000

CacheSize: 49152

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 461.743783 | 25.814485 | 4018.867020 | 4.653335 |
| 1.18 | 466.437670 | 7.901452 | 3616.372259 | 7.209999 |
| 1.19 | 520.786552 | 9.989322 | 3614.856719 | 4.258908 |
| gotip | 530.385078 | 8.595649 | 3638.611424 | 6.636146 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 121.670620 | 3.066317 | 4791.761201 | 2.903464 |
| 1.18 | 126.613622 | 2.518331 | 4796.445037 | 1.338796 |
| 1.19 | 137.663863 | 6.916269 | 4791.354641 | 0.940633 |
| gotip | 147.560705 | 5.413855 | 4791.448932 | 1.190127 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 132.341014 | 2.677917 | 5175.275218 | 69.573914 |
| 1.18 | 136.442623 | 2.527028 | 5170.515803 | 22.494548 |
| 1.19 | 140.138100 | 4.674484 | 5190.334509 | 25.613900 |
| gotip | 147.771515 | 3.901039 | 6309.864195 | 13.691646 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 134.484268 | 9.779208 | 10050.156623 | 21.952177 |
| 1.18 | 137.606311 | 5.366722 | 5902.809894 | 54.494337 |
| 1.19 | 142.264331 | 2.792543 | 10109.589957 | 45.577062 |
| gotip | 150.522898 | 4.728289 | 10258.869506 | 65.694696 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 125.737165 | 5.288528 | 5884.493111 | 121.653955 |
| 1.18 | 132.346266 | 7.381950 | 5634.956162 | 39.918948 |
| 1.19 | 142.510798 | 3.514666 | 7166.612493 | 29.329049 |
| gotip | 150.607860 | 3.479430 | 7773.142412 | 36.785769 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 143.294499 | 1.964854 | 6065.366061 | 2.329955 |
| 1.18 | 151.282894 | 3.778816 | 5977.622570 | 3.801510 |
| 1.19 | 148.641841 | 3.175901 | 5916.472136 | 2.824573 |
| gotip | 155.422573 | 2.130523 | 6025.027011 | 7.909269 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 133.976433 | 3.649059 | 3626.024136 | 271.476379 |
| 1.18 | 138.382317 | 2.438142 | 3542.840719 | 2.723307 |
| 1.19 | 141.549204 | 2.575255 | 361.181586 | 33.974794 |
| gotip | 149.112457 | 3.622284 | 397.789247 | 1.150836 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 150.476216 | 55.121986 | 5060.740900 | 493.489099 |
| 1.18 | 152.051482 | 54.668358 | 5387.926612 | 992.628620 |
| 1.19 | 155.697379 | 75.237320 | 5038.268197 | 716.950678 |
| gotip | 162.533288 | 54.547204 | 5252.792729 | 211.165287 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 124.024009 | 51.711163 | 3675.410434 | 13.119038 |
| 1.18 | 127.713891 | 3.625087 | 5216.441693 | 1.142591 |
| 1.19 | 141.815774 | 3.472810 | 2705.466009 | 1.295280 |
| gotip | 147.400556 | 2.506327 | 2705.256065 | 1.556278 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 156.294728 | 12.582660 | 1066.391828 | 1.803064 |
| 1.18 | 161.203868 | 8.705708 | 1108.154053 | 5.852979 |
| 1.19 | 165.742829 | 6.734993 | 1140.793934 | 4.521301 |
| gotip | 173.102525 | 6.842374 | 1115.759015 | 7.306640 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 164.528213 | 2.036095 | 2327.750601 | 43.774907 |
| 1.18 | 169.080198 | 2.933807 | 2341.223061 | 30.725067 |
| 1.19 | 172.283536 | 3.031397 | 2193.158723 | 12.892640 |
| gotip | 179.955465 | 5.828763 | 2165.861565 | 10.250289 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 192.151671 | 12.984527 | 5966.018252 | 17.230693 |
| 1.18 | 192.349655 | 13.069452 | 5954.008862 | 28.411622 |
| 1.19 | 196.994156 | 11.098790 | 3506.937655 | 2.280206 |
| gotip | 206.337131 | 11.220583 | 3320.188875 | 11.803356 |

![switch_case](./switch_case__725e73000e.png)

