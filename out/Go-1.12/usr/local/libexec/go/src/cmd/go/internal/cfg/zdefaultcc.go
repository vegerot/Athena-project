// Code generated by go tool dist; DO NOT EDIT.

package cfg

const DefaultPkgConfig = `pkg-config`
func DefaultCC(goos, goarch string) string {
	switch goos+`/`+goarch {
	case "linux/amd64":
		return "gcc"
	}
	return "aarch64-apple-darwin18.2-clang"
}
func DefaultCXX(goos, goarch string) string {
	switch goos+`/`+goarch {
	case "linux/amd64":
		return "g++"
	}
	return "aarch64-apple-darwin18.2-clang++"
}