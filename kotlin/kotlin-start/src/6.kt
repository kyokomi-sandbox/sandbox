fun square(i: Int): Int = i * 2

// 最初にKが出現する位置を返す関数
fun firstK(str: String): Int {
  fun isK(c: Char): Boolean = c == 'K'
  return first(str, ::isK)
//  tailrec fun go(str: String, index: Int): Int =
//      when {
//        str.isEmpty() -> -1
//        str.first() == 'K' -> index
//        else -> go(str.drop(1), index + 1)
//      }
//  return go(str, 0)
}

fun firstUpperCase(str: String): Int {
  fun isUpperCase(c: Char): Boolean = c.isUpperCase()
  return first(str, ::isUpperCase)
//  tailrec fun go(str: String, index: Int): Int =
//      when {
//        str.isEmpty() -> -1
//        str.first().isUpperCase() -> index
//        else -> go(str.drop(1), index + 1)
//      }
//  return go(str, 0)
}

fun first(str: String, predicate: (Char) -> Boolean): Int {
  tailrec fun go(str: String, index: Int): Int =
      when {
        str.isEmpty() -> -1
        predicate(str.first()) -> index
        else -> go(str.drop(1), index + 1)
      }
  return go(str, 0)
}

fun firestWhitespace(str: String): Int {
  val isWhitespace: (Char) -> Boolean = {
    it.isWhitespace()
  }
  return first(str, isWhitespace)
  // ラムダ式を引数に直接指定した例
//  return first(str, { it.isWhitespace()})
}

// ラムダ式用の特別な構文 ↑と同じ意味になる
fun firstWhitespace2(str: String): Int = first(str) { it.isWhitespace() }

// クロージャー
fun getCounter(): () -> Int {
  var count = 0
  return {
    count++
  }
}

// インライン関数。引数の関数オブジェクトがコンパイル時にインラインに展開される
// 高階関数は一般に呼び出しのコストが高い傾向があるため、その対策
inline fun log(debug: Boolean = true, message: () -> String) {
  if (debug) {
    println(message())
  }
}

// ラムダ式内でreturn文を呼ぶ（非ローカルリターンする）ためには、inline展開されている必要がある
inline fun forEach(str: String, f: (Char) -> Unit) {
  for (c in str) {
    f(c)
  }
}

fun containsDigit(str: String): Boolean {
  forEach(str) {
    if (it.isDigit()) {
      return true
    }
  }
  return false
}

fun containsDigit2(str: String): Boolean {
  // ラベルをつけてラベルへのリターン
  forEach(str) here@ {
    if (it.isDigit()) {
      return@here
    }
  }
  return false
}

fun containsDigit3(str: String): Boolean {
  forEach(str) {
    if (it.isDigit()) {
      return@forEach // 推測可能であれば関数名を指定してリターン可能
    }
  }
  return false
}

fun main(args: Array<String>) {
  // Kotlinでは関数もobjectとして扱える
  println(::square)
  // Intの引数を1つ受けてIntを返す関数の型を定義して初期値にsquareの関数objectを代入している
  val functionObject: (Int) -> Int = ::square
  println(functionObject(5))

  println(firstK("argcKadsvj"))
  println(firstUpperCase("aRgcKadsvj"))

  // ラムダ
  val square: (Int) -> Int = { i: Int ->
    i * i
  }
  println(square(100))

  // ラムダは暗黙の変数itが使える
  val square2: (Int) -> Int = {
    it * it
  }
  println(square2(30))

  val counter1 = getCounter()
  println(counter1())
  println(counter1())
  println(counter1())

  log { "このメッセージは表示される" }
  log(false) { "このメッセージは表示されない" }

  print(containsDigit("dadaegimcaa3adaa"))
}
