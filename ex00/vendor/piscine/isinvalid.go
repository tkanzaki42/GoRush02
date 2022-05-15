package piscine

func IsInvalid(args [][]rune) bool {
	if len(args) == 0 { // 空でないか？0ならそもそもmapが入ってきてない。
		return true //不正の場合
	}
	// 変な文字が入ってくる場合
	
	for _, row := range args { // "....." コマンドライン引数
		for _, r := range row { // rowに分解して1行を取り出す。さらに分解で1文字。
			if !(r == '.' || r == '#') {
				return true // 不正文字があった場合。
			}
		}
	}
	// 長方形の形を成していない
	for i := range args[1:] { // インデックスが1なのは、基準を0行目にして比べたいから
		if len(args[i]) != len(args[0]) { // 1行目とそれ以外を比べて、異なった場合は不正
			return true // 不正である場合　
		}
	}
	return false // 
	//3つのチェック
	//空のマップでないか？
	//変な文字が入ってくる場合
	//長方形かどうか
	//ファイルがあるかどうか
}
