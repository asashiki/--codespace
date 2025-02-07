var age: Int = 39

let baseAge: Int = 35

func tooOld(age: Int) -> Bool {
    return age > 35
}

if tooOld(age: age) {
    print("Too old")
} else {
    print("OK")
}