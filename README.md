# helper

These are my frequently used helper functions.

## Installation
```bash
go get github.com/dkuye/helper
```

## Function List
| SN | Name | Description | Return Value |
|---| --- | --- | --- |
| 1. | Uuid | Return Uuid string | `string`  |
| 2. | MySqlTimeStamp | Return MySql timestamp | `time.Time`  |
| 3. | MySqlTimeStampToUnixTimeStamp | Convert MySql timestamp to Unix timestamp | `int64`  |
| 4. | FormatNumber | Thousand separator | `string`  |
| 5. | FormatMoney | Thousand separator to decimal places | `string`  |
| 6. | HashPassword | Return `bcrypt` hash string  | `string`  |
| 7. | IntToString | Convert `int` value to `string` value | `string`  |
| 8. | Int64ToString | Convert `int64` value to `string` value | `string`  |
| 9. | Uint64ToSting | Convert `uint64` value to `string` value | `string`  |
| 10. | StringToInt | Convert `string` value to `int` value | `int`  |
| 11. | StringToInt | General range slice between to numbers | `[]int`  |
| 12. | PickIntRandomly | Pick an item randomly form a `[]int` | `int`  |
| 13. | PickStringRandomly | Pick an item randomly form a `[]sting` | `string`  |
| 14. | PickIntRandomlyBetween | Pick an item randomly between a range | `int`  |
| 15. | ShortNumber | Shorten thousands | `string`  |
| 16. | Shuffle | Shuffle `[]int` | `[]int`  |
| 17. | ToJson | Convert  `interface{}` to JSON | `string`  |
| 18. | UnixTimeStamp | Convert unix timestamp | `int64`  |
| 19. | ValidatePassword | Validate hashed password | `(bool, error)`  |
| 20. | TitleCase | Proper string tile case | `string`  |

