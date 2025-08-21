module example

replace github.com/xjslang/xjs => ../../xjs
replace github.com/xjslang/try-parser => ../

go 1.21.6

require github.com/xjslang/try-parser v0.0.0-00010101000000-000000000000
require github.com/xjslang/xjs v0.0.0-00010101000000-000000000000 // indirect
