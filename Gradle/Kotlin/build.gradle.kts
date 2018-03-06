task("Task1")
project.task("Task2")

val pj:String = "Abc"

task(pj)

task("Task3"){
    print("This is task3")
}

Task3.doLast{
    println("Hello")
}