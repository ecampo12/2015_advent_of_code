import sys
import os
import glob
import re

def main():
    curr_path = sys.path[0]
    os.chdir(curr_path)
    list_of_files = glob.glob('day*')
    if len(list_of_files) == 0:
        day = 1
    else:
        day = max([int(re.findall(r'\d+', day)[0]) for day in list_of_files]) + 1
    print(f"Latest folder: {day}")
    print(f"Creating day{day} folder")
    os.mkdir(f"{curr_path}/day{day}")
    os.system(f"cp {curr_path}/AOC_temp.py {curr_path}/day{day}/AOC.py")
    os.system(f"cp {curr_path}/test_temp.py {curr_path}/day{day}/test.py")
    os.chdir(f"{curr_path}/day{day}")
    open("input.txt", "x")
    open("test_input.txt", "x")
    
    print(f"Created day{day} folder")
    # I know I could use gitpython, but I don't want to install it
    os.system("git add .")
    os.system(f"git commit -m \"Created by script: Day{day}\"")
    
if __name__ == "__main__":
    main()