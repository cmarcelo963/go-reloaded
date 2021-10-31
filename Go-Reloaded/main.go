package main

import (
	"gr"
)

func main() {
	file := gr.ReadFile()
	file = gr.ToHex(file)
	file = gr.ToBin(file)
	file = gr.ToUp(file)
	file = gr.ToLow(file)
	file = gr.ToCap(file)
	file = gr.Punc(file)
	file = gr.A(file)
	
	gr.CreateFile(file)
}
