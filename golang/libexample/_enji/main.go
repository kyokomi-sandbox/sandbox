package main

import "github.com/ajhager/engi"

type Game struct {
	*engi.Game
	gopher engi.Drawable
	batch  *engi.Batch
}

func (game *Game) Preload() {
	engi.Files.Add("gopher", "data/gopher.png")
	game.batch = engi.NewBatch(engi.Width(), engi.Height())
}

func (game *Game) Setup() {
	engi.SetBg(0x2d3739) // 背景の色
	game.gopher = engi.Files.Image("gopher")
	//	game.font = engi.NewGridFont(engi.Files.Image("font"), 20, 20)
}

func (game *Game) Render() {
	game.batch.Begin()
	game.batch.Draw(game.gopher, 1024/2, 640/2, 0.5, 0.5, 1, 1, 0, 0xffffff, 1) // ロゴを表示
	game.batch.End()
}

func main() {
	// タイトル, height, width, fullscreen, Gameオブジェクト
	engi.Open("Hello", 1024, 640, false, &Game{})
}
