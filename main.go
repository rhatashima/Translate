// パッケージの読み込み
package main
// ライブラリのインポート
import(
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"flag"
)


func main() {

	// ヘルプメッセージのカスタマイズ
	flag.Usage = func() {
    usage := `translate ver1.0: This program translate dna to protein.
Usage: translate <dna.fa> [> protein.fa]`

    fmt.Fprintf(os.Stderr, "%s\n", usage)
  }

	// プログラム引数をスライスargsとして取得
	flag.Parse()
	args := flag.Args()

	// fastaファイルが指定されていなければエラー出力
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Error! Fasta file is missing.\n")
		flag.Usage()
		os.Exit(0)
	}

	// get the file path
	filePath := args[0]

	//　open the file
	file, err := os.Open(filePath)

	// 関数の終了時にファイルを閉じる
	defer file.Close()

	// handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	// コドン表を定義
	var codonMap = map[string]string{
		"TTT" : "F", "TCT" : "S", "TAT" : "Y", "TGT" : "C",
		"TTC" : "F", "TCC" : "S", "TAC" : "Y", "TGC" : "C",
		"TTA" : "L", "TCA" : "S", "TAA" : "*", "TGA" : "*",
		"TTG" : "L", "TCG" : "S", "TAG" : "*", "TGG" : "W",

		"CTT" : "L", "CCT" : "P", "CAT" : "H", "CGT" : "R",
		"CTC" : "L", "CCC" : "P", "CAC" : "H", "CGC" : "R",
		"CTA" : "L", "CCA" : "P", "CAA" : "Q", "CGA" : "R",
		"CTG" : "L", "CCG" : "P", "CAG" : "Q", "CGG" : "R",

		"ATT" : "I", "ACT" : "T", "AAT" : "N", "AGT" : "S",
		"ATC" : "I", "ACC" : "T", "AAC" : "N", "AGC" : "S",
		"ATA" : "I", "ACA" : "T", "AAA" : "K", "AGA" : "R",
		"ATG" : "M", "ACG" : "T", "AAG" : "K", "AGG" : "R",

		"GTT" : "V", "GCT" : "A", "GAT" : "D", "GGT" : "G",
		"GTC" : "V", "GCC" : "A", "GAC" : "D", "GGC" : "G",
		"GTA" : "V", "GCA" : "A", "GAA" : "E", "GGA" : "G",
		"GTG" : "V", "GCG" : "A", "GAG" : "E", "GGG" : "G"}

	// make a scanner for file reading
	fileScanner := bufio.NewScanner(file)

	// 翻訳前の塩基配列を保持する箱を作成
	var seq string

	// read line by line
	for i := 1; fileScanner.Scan(); i++ {

		// 一行読み込む
		line := fileScanner.Text()

		// ID行の場合
		if strings.HasPrefix(line, ">") {

			// 1行目でなければ翻訳してアミノ酸配列を出力
			if i != 1 {
				fmt.Println(Translate(seq, codonMap))
			}

			// IDを出力
			fmt.Println(line)

			// 変数seqの初期化
			seq = ""

		// 配列行の場合
		} else {
			// seqに追記
			seq += line
		}
	}

	// 最終配列を翻訳して出力
	fmt.Println(Translate(seq, codonMap))
}