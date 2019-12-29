# Helper Functions

These are my frequently used helper functions.

## Installation
```bash
go get github.com/dkuye/helper
```

## Function List
| Name | Description | Return Value |
| Uuid | Return Uuid string | `string`  |
| MySqlTimeStamp | Return MySql timestamp | `time.Time`  |
| MySqlTimeStampToUnixTimeStamp | Convert MySql timestamp to Unix timestamp | `int64`  |
| FormatNumber | Thousand separator | `string`  |
| FormatMoney | Thousand separator to decimal places | `string`  |
| HashPassword | Return `bcrypt` hash string  | `string`  |
| IntToString | Convert `int` value to `string` value | `string`  |
| Int64ToString | Convert `int64` value to `string` value | `string`  |
| Uint64ToSting | Convert `uint64` value to `string` value | `string`  |
| StringToInt | Convert `string` value to `int` value | `int`  |
| StringToInt | General range slice between to numbers | `[]int`  |
| PickIntRandomly | Pick an item randomly form a `[]int` | `int`  |
| PickStringRandomly | Pick an item randomly form a `[]sting` | `string`  |
| PickIntRandomlyBetween | Pick an item randomly between a range | `int`  |
| ShortNumber | Shorten thousands | `string`  |
| Shuffle | Shuffle `[]int` | `[]int`  |
| ToJson | Convert  `interface{}` to JSON | `string`  |
| UnixTimeStamp | Convert unix timestamp | `int64`  |
| ValidatePassword | Validate hashed password | `(bool, error)`  |

