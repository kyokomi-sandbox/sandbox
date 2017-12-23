package start12

fun square(i: Int): Int = i * i

fun main(args: Array<String>) {
//  val s : String = null // nullは非nullのStringの値になることはないエラーs
  val s: String? = null
  println(s)
//  println(s.toUpperCase()) // nullになりえるエラー
  println(s?.toUpperCase()) // safe call

  val a: Int? = 5
  // let関数は任意のT型の引数をblockに取って、
  // そのblockに対してletのレシーバーとなるオブジェクトを引数に与えて呼び出す
  val aSquare = a?.let { square(it) }
//  val aSquare =
//      if (a != null) square(a)
//      else null
  println(aSquare)

  val foo: String? = "Hello"
  val bar: String = foo!! // 強制的にNotNullに変換する（危険）
  println(bar.toUpperCase())

//  val foo2: String? = null
//  val bar2: String = foo2!! // 強制的にNotNullに変換する（危険）
//  println(bar2.toUpperCase()) // NullPointer
  val foo2: String? = null
//  val bar2: String = requireNotNull(foo2, { "foo2はNullであってはダメ" }) // java.lang.IllegalArgumentException: foo2はNullであってはダメ
//  println(bar2.toUpperCase())

  // エルビス演算子でnullなら指定のデフォルト値を扱うことができるs
  println((foo2 ?: "default").toUpperCase())
}
