// +build dev

package build

import (
	"os"
	. "github.com/microgiantya/logger"

	"debug/buildinfo"
)

func init() {
        info, err := buildinfo.ReadFile(os.Args[0])
        if err != nil {
                E(p, err)
                os.Exit(1)
        }
        I(p, info.String())

}
