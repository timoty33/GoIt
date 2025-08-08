package file

import "os"

// FileExists verifica se o arquivo ou diretório existe no caminho dado
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
