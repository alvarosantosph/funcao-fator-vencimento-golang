/*
 * Cálculo fator de vencimento padrão febraban (Federação Brasileira de Bancos)
 * Linguagem: GoLang
 * Autor: Alvaro Santos
 * Data: 13/11/2019
 * GitHub: https://github.com/alvarosantosph
 * Email: alvaro_webmaster@hotmail.com
 */

// Funcao para Gerar Fator de Vencimento de Acordo com a Data de Vencimento
// Utilizar a data no formato DD/MM/AAAA
func gerarFatorVencimento(stringData string) string {
	runes := []rune(stringData)
	dataParte1 := string(runes[0:2])
	dataParte2 := string(runes[3:5])
	dataParte3 := string(runes[6:10])
	int1, _ := strconv.Atoi(dataParte1)
	int2, _ := strconv.Atoi(dataParte2)
	int3, _ := strconv.Atoi(dataParte3)
	st := time.Date(1997, time.Month(10), 7, 00, 00, 00, 0, time.UTC)
	en := time.Date(int3, time.Month(int2), int1, 00, 00, 00, 0, time.UTC)
	diff := (en.Sub(st))
	calculoMilisegundo := diff.Milliseconds() + 3600000
	fatorVencimento := calculoMilisegundo / 86400000
	retorno := strconv.FormatInt(fatorVencimento, 10)
	return retorno
}
