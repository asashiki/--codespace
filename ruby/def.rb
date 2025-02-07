#age = 39

def base_age   #ruby 方法定义没有变量可以不用括号
    39
end

def too_old?(age)  #?代表布尔值
  age > 35
end

if too_old?(base_age)
  puts "Too old"
else
  puts "OK"
end