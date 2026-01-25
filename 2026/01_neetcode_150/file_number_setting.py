import os
import re

def balance_file_numbers(directory=".", padding_width=4):
    # Regex to capture leading digits and the rest of the filename
    # Example: '1_file.c' -> ('1', '_file.c')
    pattern = re.compile(r'^(\d+)(.*)')

    for filename in os.listdir(directory):
        match = pattern.match(filename)
        
        if match:
            original_num = match.group(1)
            rest_of_name = match.group(2)
            
            # Pad the number with leading zeros
            balanced_num = original_num.zfill(padding_width)
            new_filename = f"{balanced_num}{rest_of_name}"
            
            # Only rename if the name actually changes
            if filename != new_filename:
                old_path = os.path.join(directory, filename)
                new_path = os.path.join(directory, new_filename)
                
                print(f"Renaming: {filename} -> {new_filename}")
                os.rename(old_path, new_path)

if __name__ == "__main__":
    # You can change '.' to your specific folder path
    balance_file_numbers(directory=".", padding_width=4)