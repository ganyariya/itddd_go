# itddd_go

[![](https://www.seshop.com/static/images/product/20675/L.png)](https://www.shoeisha.co.jp/book/detail/9784798150727)

https://github.com/nrslib/itddd

ドメイン駆動設計入門を参考に go 言語で DDD を実装しています。
本の内容のまとめは以下 Zenn Scrap になります。

https://zenn.dev/ganariya/scraps/55191eccb40b44

### 5 章までのクリーンアーキテクチャ構成

5 章までは、見慣れているクリーンアーキテクチャで実装していました。

https://github.com/ganyariya/itddd_go/tree/2d16615e2f70c9099eaf02fd6256aef1387b170f

6 章以降はパッケージ構成を変更し、レイヤードアーキテクチャとなっています。
（UseCase InputPort などを利用せず、 ApplicationService のみとなっています。）
