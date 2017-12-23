package start8

// Kotlinはデフォルトではクラスを継承できない
// 修飾子openをつける必要がある
// また、プロパティも同様にopenをつけることでoverride可能になる
open class Person(open val name: String) {
  // overrideもデフォルトでは禁止されているのでopenをつける必要がある
  open fun introduceMyself() {
    println("I am $name")
  }
}

class Student(override val name: String, val id: Long) : Person(name) {
  override fun introduceMyself() {
    println("I am $name(id=$id)")
  }
}

abstract class Greeter(val target: String) {
  abstract fun sayHello()
}

class EnglishGreeter(target: String) : Greeter(target) {
  override fun sayHello() {
    println("Hello, $target")
  }
}

class JapaneseGreeter(target: String) : Greeter(target) {
  override fun sayHello() {
    println("こんにちわ、$target")
  }
}

fun main(args: Array<String>) {
  val person: Person = Person("翡翠")
  person.introduceMyself()

  val student: Student = Student("秋葉様", 123)
  println(student.id)
  println(student.name)
  student.introduceMyself()

  // StudentはPersonである
  val person2: Person = Student("琥珀さん", 456)
  person2.introduceMyself()
  //person2.id // コンパイルエラー

  EnglishGreeter("Kohaku").sayHello()
  JapaneseGreeter("琥珀さん").sayHello()
}
