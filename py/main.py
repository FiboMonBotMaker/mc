import os

google="https://www.google.com/search?q={}\n"

with open("list.txt", "w") as f:
    s = ""
    for n in os.listdir("data"):
        if ".jar" in n:
            s += google.format(n)
    f.write(s)
    
