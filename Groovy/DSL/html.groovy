import java.lang.reflect.Array

def invokeMethod(String name, args) {
    print "<${name}"
    args.each { arg ->
        if (arg instanceof Map) {
            arg.each {
                print " ${it.key} ='${it.value}' "
            }
        } else if(arg instanceof ArrayList){
            print arg
            arg.each {print it}
        }
        else if(arg instanceof Closure) {
            print '>'
            arg.delegate = this
            def value = arg.call()
            if (value) {
                print "${value}"
            }
        }
    }

    println "</${name}>"
}





html  {
    head {
        meta {

        }
    }
    body {
        table(style: 'margin:5px;') {
            tr('class': 'trClass', style: 'padding:5px;') {
                td { 'Cell1' }
                td { 'Cell2' }
                td { 'Cell3' }
            }
        }
    }
}