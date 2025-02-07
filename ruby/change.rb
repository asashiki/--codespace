age = 20
name = ["John", "Doe"]

def change_age(age)
  age = 30
end

def change_name(name)
    name[0] = "Jane"
end

change_age(age)
change_name(name)

puts age
puts name