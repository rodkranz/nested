# Changelog

## [Unreleased]

## [1.1.0] - 2020-02-13
### Added
 - Added method `NewFromInterface`. try to cast from `interface` to `map[string]interface`

## [1.0.0]
### Added
 - `New` - method cast `map[string]interface` to `Map`
 - `Interface` - retrieve any data from `map` looking for the position as `interface` or `map[string]interface` type
 - `String` - retrieve `string` from `map` looking for the position as `string` type
 - `Int` - retrieve `int` from `map` looking for the position as `int` type
 - `Time` - retrieve `time` from `map` looking for the position as `time.Time` type
