import sys

# name ="David";
# name = 42;

# print(name);

if len(sys.argv) > 1:
    try:
        number = int(sys.argv[1])
        if number > 35:
            print("TOO OLD")
        else:
            print("TOO YOUNG")
    except ValueError:
        print("Please enter a valid number.")
else:
    print("No number entered.")