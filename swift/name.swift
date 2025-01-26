// var name: String = "Swift";
// name = "42"; // Error: Cannot assign value of type 'Int' to type 'String'

// print(name);


import Foundation

print("Enter a number:")
if let input = readLine(), let number = Int(input) {
    if number > 35 {
        print("Too old")
    } else {
        print("OK")
    }
    // print("You entered: \(number)")
} else {
    print("Invalid input")
}