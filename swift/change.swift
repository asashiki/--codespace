var age: Int = 39
var name: [String] = ["John", "Doe"]

func change_age(age: inout Int) {
    age = 20
}
func change_name(name: inout [String]) {
    name = ["Jane", "Doe"]
}

change_age(age: &age)
change_name(name: &name)

print(age)
print(name)
