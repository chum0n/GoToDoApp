# 計算機科学実験4

## 困った時には
GORM関連  
公式  
http://gorm.io/ja_JP/docs/  
Qiita  
https://qiita.com/chan-p/items/cf3e007b82cc7fce2d81  

GIN関連
公式  
https://gin-gonic.com/ja/docs/

## メモ

### データベース
GORMでのpostgresは公式ドキュメントのやり方通り
postgres  
ER図の作成にはLucidchart(http://www.lucidchart.com)を使用しようとしたが、よくわからなかったのでMySQLWorkbenchを使用  
以下の記事が書き方の参考になる  
https://qiita.com/noborus/items/11438d16f790b1d42ad8  
書き方はIE記法を学習（ググったらいっぱい出てくる）

SQL文の書き方は公式ドキュメント見る

## 困ったこと
gormはstructの複数形をテーブル名にする  
それを避けるにはテーブル名を指定しなければならない  
http://gorm.io/ja_JP/docs/conventions.html

integer 型の主キーでのみ動作します  
db.First()は  
http://gorm.io/ja_JP/docs/query.html

can't find importと言われた  
go getに-uオプションをつける

css読み込めない  
mainで静的ファイルを読み込ませる処理をかく  
router.Static("/assets", "./assets")

<h1>Signup</h1>
    <h2>追加</h2>
    <!-- /newにPOSTで投げる -->
    <!-- Go側のcreateのとことつながっている -->
    <form method="post" action="/new">
        <p>ID<input type="text" name="customer_id" size="30" placeholder="入力してください" ></p>
        <p>名前<input type="text" name="customer_name" size="30" placeholder="入力してください" ></p>
        <p>年齢<input type="text" name="age" size="30" placeholder="入力してください" ></p>
        <p>性別(男は1,女は2)<input type="text" name="gender" size="30" placeholder="入力してください" ></p>
        <p><input type="submit" value="Send"></p>
    </form>


<h1>カスタマー情報</h1>
    <h2>追加</h2>
    <!-- /newにPOSTで投げる -->
    <!-- Go側のcreateのとことつながっている -->
    <form method="post" action="/new">
        <p>ID<input type="text" name="customer_id" size="30" placeholder="入力してください" ></p>
        <p>名前<input type="text" name="customer_name" size="30" placeholder="入力してください" ></p>
        <p>年齢<input type="text" name="age" size="30" placeholder="入力してください" ></p>
        <p>性別(男は1,女は2)<input type="text" name="gender" size="30" placeholder="入力してください" ></p>
        <p><input type="submit" value="追加"></p>
    </form>

    <ul>
        {{ range .customers }}
            <li>ID：{{ .Customer_id }}、名前：{{ .Customer_name }}、年齢：{{ .Age }}、性別：{{ .Gender }}
                <label><a href="/detail/{{.Customer_id}}">編集</a></label>
                <label><a href="/delete_check/{{.Customer_id}}">削除</a></label>
            </li>
        {{end}}
    </ul>