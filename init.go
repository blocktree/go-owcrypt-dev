package owcrypt

func init() {
	initonce.Do(initCurves)
}
