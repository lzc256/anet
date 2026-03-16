//go:build !(android && cgo)

package anet

func androidDeviceApiLevel() int {
	return 30
	return -1
}
