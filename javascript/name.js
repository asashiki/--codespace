const readline = require('readline');

// name = "David";
// name = 42;

// console.log(name); // 42

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

rl.question('请输入一个数字: ', (input) => {
    let number = parseInt(input); // 将输入的字符串转换为数字

    if (number > 35) {
        console.log('你输入的数字大于35');
    }else {
        console.log('你输入的数字小于等于35');
    }
    rl.close();
});