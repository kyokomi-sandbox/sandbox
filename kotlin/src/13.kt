package start13

// data classはtoStringとか色々実装してくれて便利
data class User(val id: Long, val name: String) {
  // companionをつけることでシングルトンなオブジェクトをclass内に定義できる
  companion object {
    val DUMMY = User(0, "dummy")
  }

  // 修飾子innerをつけると外側のUserのオブジェクトの参照もつかむことができる
  inner class Action {
    fun show(): String = "$name($id)"
  }
}

interface Greeter {
  fun greet(name: String)
}

// オブジェクト宣言でシングルトン的なクラスが定義可能
object JapaneseGreeter : Greeter {
  override fun greet(name: String) {
    println("こんにちわ、${name}さん!")
  }
}

fun main(args: Array<String>) {
  val u1 = User(1, "七瀬佳乃")
  println(u1)
  println(u1.Action().show())

  val greeter = object : Greeter {
    override fun greet(name: String) {
      println("Hello $name")
    }
  }
  greeter.greet("Test")

  JapaneseGreeter.greet("テスト")

  println(User.DUMMY.Action().show())
}
