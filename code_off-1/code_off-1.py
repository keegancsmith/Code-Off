#!/usr/bin/env python

from collections import defaultdict

def val(word):
    return sum(ord(x) - ord('A') + 1 if x.isupper() else ord(x) - ord('a')
               for x in word)

def equivalence_class(words):
    equiv = defaultdict(list)
    for word in words:
        equiv[val(word)].append(word)
    return equiv
    
def main():
    fin = file('code_off-1.in')
    n = int(fin.readline())
    words = [fin.readline().strip() for _ in range(n)]
    
    equiv = equivalence_class(words)
    for word in words:
        is_palindrome = word == word[::-1]
        equiv_words = [x for x in equiv[val(word)] if x != word]
        print(word)
        print('true' if is_palindrome else 'false')
        print('\n'.join(equiv_words))
        if equiv_words:
            print('')

if __name__ == '__main__':
    main()
