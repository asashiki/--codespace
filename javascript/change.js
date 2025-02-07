var age = 36;
var names = ['John', 'Jane', 'Doe'];

function changeAge(age) {
//  age = 39;
  return 39;
}

function changeNames(names) {
    names[0] = 'John Doe';
    return names;
    }

age = change(age);
changeNames(names);

console.log(age); 
console.log(names[0]);