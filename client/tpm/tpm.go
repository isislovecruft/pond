// Package tpm wraps the Trousers library for accessing the TPM from
// user-space. It currently provides very limited functionality: just NVRAM
// access.
package tpm

// Stubbed out because we have no TPM chip, and don't want one, and don't want
// to install libtpsi-dev

func Present() bool {
	return false
}

type Context struct {}

type NVRAM struct {}

func (c *Context) NewNVRAM() (*NVRAM, error) {
	return nil, nil
}
