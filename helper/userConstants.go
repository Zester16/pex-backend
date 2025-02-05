package helper

import "os"

func GetTableNames(inp string) string {

	constants := map[string]string{
		"": "",
	}

	return constants[inp]

}

func GetJWTSecretKey() []byte {
	return []byte(os.Getenv("JWT_KEY"))
}
