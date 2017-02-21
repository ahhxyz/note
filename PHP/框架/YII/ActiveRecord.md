rules()：

里面的字段都是属性名，AR会自动将数据库字段设置为属性。
如果要处理数据库没有字段，那么就要声明属性。






###where()
___

`where(['OR', ['material_id' => $rootIds], ['child_material_id' => $childrenIds]])`


SQL : `WHERE ("material_id"=30) OR ("child_material_id"=32)`