package start10

interface Greeter {
  fun sayHello(target: String)
  fun sayHello()
}

class EnglishGreeter : Greeter {
  override fun sayHello(target: String) {
    println("Hello $target!")
  }

  override fun sayHello() {
    println("Hello no name")
  }
}

open class Superclass

interface Foo {
  fun aaa()
  fun bbb()
}

// interfaceの継承はOK
interface Bar : Foo {
  override fun aaa() {}
  fun ccc()
}

class MyClass : Superclass(), Foo, Bar {
  override fun bbb() {}
  override fun ccc() {}
}

open class JapaneseGreeter : Greeter {
  override fun sayHello(target: String) {
    println("こんにちわ $target")
  }

  override fun sayHello() {
    sayHello("匿名さん")
  }
}

class JapaneseGreeterWithRecording : JapaneseGreeter() {
  private val _targets: MutableSet<String> = mutableSetOf()

  val targets: Set<String> get() = _targets

  override fun sayHello(target: String) {
    _targets += target
    super.sayHello(target)
  }
}

fun main(args: Array<String>) {
  val greeter: Greeter = EnglishGreeter()
  greeter.sayHello("test")

  val g = JapaneseGreeterWithRecording()
  g.sayHello("片山実波（CV: 田中美海）")
  g.sayHello("岡本未夕（CV: 高木美佑）")
  g.sayHello()
  println(g.targets)
}
