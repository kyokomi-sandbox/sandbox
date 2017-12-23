package start11

class Container<T>(val value: T)

fun <T> box(value: T): Container<T> = Container(value)

val <T> T.string: String
  get() = toString()

fun main(args: Array<String>) {
  val container: Container<Int> = box(5)
  println(container.string)
}
