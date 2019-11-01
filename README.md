# 計算機科学実験4

## メモ

### データベース
postgres  
ER図の作成にはLucidchart(http://www.lucidchart.com)を使用しようとしたが、よくわからなかったのでMySQLWorkbenchを使用  
以下の記事が書き方の参考になる  
https://qiita.com/noborus/items/11438d16f790b1d42ad8  
書き方はIE記法を学習（ググったらいっぱい出てくる）

## 困ったこと
gormはstructの複数形をテーブル名にする  
それを避けるにはテーブル名を指定しなければならない  
http://gorm.io/ja_JP/docs/conventions.html

integer 型の主キーでのみ動作します  
db.First()は  
http://gorm.io/ja_JP/docs/query.html