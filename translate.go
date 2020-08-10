// パッケージの読み込み
package main
// ライブラリのインポート
import(
	"strings"
)

// 塩基配列を翻訳するtranslate関数を定義
func Translate(nuc string, codonMap map[string]string) string {

	// 塩基配列をすべて大文字に変換
	nuc = strings.ToUpper(nuc)

	// 配列長を取得
	length := len(nuc)

	// 翻訳したアミノ酸配列を代入していく箱を用意
	prot := ""

	// 変数codonを定義
	var codon string

	// 翻訳の実行
	for i := 2; i <= length-1; i += 3 {
		codon = string(nuc[i-2]) + string(nuc[i-1]) + string(nuc[i])

		// コドンを翻訳
		if aa, ok := codonMap[codon]; ok {
			prot += aa
		} else {
			prot += "X"
		}
	}

	// 配列長が3の倍数でない場合に最後の余ったコドンを翻訳
	if length % 3 != 0 {
		prot += "X"
	}

	return prot
}