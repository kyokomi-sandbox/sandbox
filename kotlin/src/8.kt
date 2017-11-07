class Greeter {
  fun greet(name: String) {
    println("Hello $name!")
  }
}

class Person {
  var name: String = ""
  var age: Int = 0
}

class MyClass {
  // DIやユニットテストなどで初期化のタイミングを遅らせたいときに使う
  // あまり使わないほうがよい
  lateinit var foo: String
}

// コンストラクタで引数をうけて初期化
class Rational2 constructor(n: Int, d: Int) {
  val numerator: Int = n
  val denominator: Int = d
}

// 引数をそのままプロパティにする
class Rational3(val numerator: Int, val denominator: Int) {
  // セカンダリコンストラクタ（0個以上持つことができる）
  constructor(numerator: Int) : this(numerator, 1)
}


// 引数をそのままプロパティにする
class Rational4(val numerator: Int, val denominator: Int = 1) {
  // イニシャライザ（インスタンス生成の際に実施しておきたい処理を定義できる）
  init {
    require(denominator != 0)
  }
}

fun main(args: Array<String>) {
  val greeter = Greeter()
  greeter.greet("Kotlin")

  val hanako = Person()
  println(hanako.name)
  println(hanako.age)

  //val a = MyClass()
  //println(a.foo) // => UninitializedPropertyAccessException

  val half = Rational2(1, 2)
  println(half.numerator)
  println(half.denominator)
  val half2 = Rational3(1, 2)
  println(half2.numerator)
  println(half2.denominator)
  val half3 = Rational3(5)
  println(half3.numerator)
  println(half3.denominator)
  //val half4 = Rational4(5, 0) // IllegalArgumentException



}
