human = True

if human:
    print("I'm a Human")
else:
    print("I'm not a Human")

match human:
    case True:
        print("I'm a Human")
    case False:
        print("I'm not a Human")

for i in range(10):
    print(i)

while True:
    print("Hello World")