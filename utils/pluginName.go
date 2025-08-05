package utils

import (
	"regexp"
	"strings"
)

func GetPluginNameFromUrl(url string) string {
	// Regex para extrair o nome do plugin do URL
	re := regexp.MustCompile(`github\.com/([^/]+)/([^/]+)`)
	matches := re.FindStringSubmatch(url)
	if len(matches) >= 3 {
		return matches[2] // Retorna o nome do repositório
	}
	return ""
}

func SearchPlugin(urls []string, nomePlugin string) string {
	for _, url := range urls {
		if strings.HasSuffix(url, nomePlugin) {
			urlGithubPlugin := url
			return urlGithubPlugin
		}
	}
	return ""
}
