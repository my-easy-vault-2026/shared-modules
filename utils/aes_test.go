package utils

import (
	"encoding/base64"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	privKey := `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQC8DjFPgM2m5z214irOf1ZJmhvQT5C96tJK9YTQXEsFs1+EAVwB
EZAUDP8/nVmM1OJYjy7XHNXL42XxsKrw/PGOOECzjbk4bP7OEr9T0m2U0f5aAyoy
k3mJBxxkvVSIggZU9+KZUzlyXRU+tq48q0ksbU5d2rg8awx/tJ1ZH+Ye7wIDAQAB
AoGAbZ8sWXbw6pkjnPHxGBycdR0zl3O6IStQWMfUGw2h4fnU9QNB2ZGeVkHZTQDx
Yan7I7qbiGPTsR2moIPmscIRIGDd6wqoCxd31vq3bEt/j2Hcwzdy7j4OfxCLrJWZ
8yycrDqLNeDyfDfCxSAzLXd/iUWvJ5g0jcTvhy5BhyqSsGkCQQDLka7Z4scRufhw
AQZ70bsKZmtj8akkfAFig2GR0ngIukLEaiTdxblBosztdiRIQQFpuScErIrSS3ub
E6fbKUiLAkEA7H2ebxq/jC4mAtH7pXsmgPhffjC+391Pd8NEryjvbV0WC8b9AvnF
CI6QFOYQNQCeE363iLHMpUGRbR2AUsprrQJAN1x8P9czqjd4QAWXXM8R3ecyp3CO
PlTXD5KJU134tO7qv33aXtGp6xa6Qo4RmfDL0JPA7714124dxHPY/3o/twJBAIaG
brmE3bFADx9Lk3pus1hp8Og3klyF586YpVl+T7RGX2QTrZkju5FCh3Nb65w63bD5
RC3d9iuLQM5xGa3+t4ECQCaRBr73tzqtwB9SnQYzmw5r0CAiOlZahZviyDCrDQTg
SbcCFGRwHR64uJKI8bRRzSXEjhToVelspIFWEAOfIWw=
-----END RSA PRIVATE KEY-----`

	var ServerAesKey = []byte{0x76, 0xfc, 0xfe, 0x63, 0x18, 0x7f, 0x5c, 0x27, 0xab, 0x7c, 0x3d, 0x4d, 0x7b, 0xf2, 0x59, 0xd0}
	var iv = []byte{0x76, 0xfc, 0xfe, 0x63, 0x18, 0x7f, 0x5c, 0x27, 0xab, 0x7c, 0x3d, 0x4d, 0x7b, 0xf2, 0x59, 0xd0}

	encrypted, err := AesEncrypt([]byte(privKey), ServerAesKey, iv)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(privKey), base64.StdEncoding.EncodeToString(encrypted), len(base64.StdEncoding.EncodeToString(encrypted)))
}

func TestAesDecrypt(t *testing.T) {
	var ServerAesKey = []byte{0x76, 0xfc, 0xfe, 0x63, 0x18, 0x7f, 0x5c, 0x27, 0xab, 0x7c, 0x3d, 0x4d, 0x7b, 0xf2, 0x59, 0xd0}
	var iv = []byte{0x76, 0xfc, 0xfe, 0x63, 0x18, 0x7f, 0x5c, 0x27, 0xab, 0x7c, 0x3d, 0x4d, 0x7b, 0xf2, 0x59, 0xd0}

	privKey := `VXZoP9Pi8XAkD0AxCGuhjx5ipXB2YBDgZxIFlcID+W/tY7k63V+yhGwrjOGSmeQDpn7D5DBDfpiixHwpH10Buh/WxS7OY3/Gqg8OIDy6+gNnwJFFYlWff6CwxrIbu/RJTemejTgD+TN2GScH0f858STyqhixUTU+UBy2OziDiAj94BCJAc3eVjVXCR5K7/1/+q1QqiUAfJ8QKw/w29/h6mRV51MLEGEw94yEhXQsFoZ6tnzS/ONawdwXNQa5tzhIyYOjklnHzNGv5ZtJHF16Nk/XYajrNQ8ya8qZH/pQL3UuLyI1CixaRbuMhAwUt2mt3EGxgv884DsxfTkQKLRJhiatnwLA9C1OehYAqf4bRze4jgb0bJLyDpdOkrwrmxOuY9mf9F9YRKyaXf5zdb5RgEibS/a/hkB8Ez7OForaEEz1mwQnfrOjbyHofFiq5kNlTN5OPCM81JCOb7RHFV9rlTAw/5x6/sKWntWCRQGROaHd5NohTKXaGW2SUQ+1KY1Fj6JItWUbOfoZp57xcUG/Dg1LMQwcMGKsyVvrvVrsYp6e3XXEI4YuY7fI0pGism+jHpKqdpwcXqW3RpPGQX9jjS+q/mn1wNohKbQLs27CuNZF3MK03VXQnfgKM4qSfaJSacTOME7rFLb6/b7KouTrM57nc/TtEHK+4AUyDRdAtdA/MYYs3x2g7YmqHIWn0kXkghfpRJgWzbX8MHfQMD3ylOkOJDs2cP/9ncXAvRE5CRrv7gMyyaPWTHhfocuBsYqlM7/aXmfEUFOKukxmaMtSNe2IoKU4cLPOrDZ8YwIvgesW1vnVMuh8YSNYCJZejYOcaKdVOaQwJa+h02dy/HlmzymwRykm07W9jzWGo9diPjKDAFShjXNkCPRKj9JnEgRfVl2FDd1j8/NEq+Fuw29UmOZCozQrUqUxYXIvOunYy8pOyGOH36RDOnnBYzFsfBunb7hTNsgOkHc/f6bhy+PcJxZN3rPbYOOcPSyxI62YHeAMs6PpxD4EzeQI5mUTPpf2CfIAMb4EbmmUPqvOU/PLHG6T2YGjmLfFCkJid/Pjcx/imxU6zQUDryzkEYEI6pmrbRp7DxH9GFQcelMLumO2iZJ2wTLNaWW8IWkEAi9L0fnU30fAT7e7EoDrve1fp2zDoLWeX9WsP3Dp2nlcGNN7+6lkWB+dIfCq1+FviuPIu4U=`
	data, err := base64.StdEncoding.DecodeString(privKey)

	if err != nil {
		t.Fatal(err)
	}

	decrypted, err := AesDecrypt(data, ServerAesKey, iv)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(decrypted))

	signBytes, err := RsaSignWithSha256([]byte("xxxxxx"), decrypted)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(base64.StdEncoding.EncodeToString(signBytes))
}
