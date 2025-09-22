import pprint

dog_name = 'rex'
a = 10
b = 2
# just some operations
print(a+b)
print(a-b)
print(a*b)
print(a/b)
print(a%b)
""""
Cool staf here
"""
for i in range(5):
    print(i)

pprint(list(range(5)))

i = 0 
if i == 45:
    print("i is 45")
elif i == 35:
    print("i is 35")
elif i > 10:
    print("i is greater than 10")
elif i%3 == 0:
    print("i is multiple of 3")
else:
    print("i is less than 10")

cat = 'Murzik''
if 'z' in cat:
    print("z is in cat name")
    if cat == 'Murzik':
        print("cat name is Murzik")
    else:
        print("Some other cat")
else:
    print("z is not contains in the cat name")

print("just some while loop")
while count < 5:
    print(count)
    count += 1

count = 0
print("änother loop")
while True:
    print(f"the count is {count}")
    if count > 5:
        break   
    count += 1

# exceptions
thinkers = ['One', 'two', 'three']

while True:
    try:
        thinker = thinkers.pop()
        print(f"the thinker is {thinker}")
    except IndexError as e:
        print("No more thinkers")
        print(e)
        break

