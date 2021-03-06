# Release Notes

[Latest release](https://github.com/allaman/go-cli-skeleton/releases/latest))

## 0.5.0

### Added

- new subcommand for parsing directories
  go-cli-skeleton dir
  go-cli-skeleton dir -p /tmp
  go-cli-skeleton dir -r -p /tmp

## 0.4.1

### Changed

- renamed `writeBytesToFile` to `writeToFile`
- `writeToFile` accepts an `io.Reader` and does not decide whether input is from stdin (no side effects regarding that)

### Removed

- `readFromStdintoByte` is now obsolet

## 0.4.0

### Added

- JSON parsing via [gjson](https://github.com/tidwall/gjson) and unmarshalling

### Changed

- Concise naming

## 0.3.1

### Changed

- reorder command order (cosmetic)
- bump go version to 1.18

## 0.3.0

### Changed

- `file` command for file operations replacing `read-file` and `write-file`

### Added

- support for reading from stdin for writing content to a file

## 0.2.0

### Changed

- arguments fir file-reading and file-writing required

## 0.1.0

### Added

- file reading
- file writing

## 0.0.10

### Fixed

- wrong projectname in goreleaser

## 0.0.9

### Changelog

- initial release
