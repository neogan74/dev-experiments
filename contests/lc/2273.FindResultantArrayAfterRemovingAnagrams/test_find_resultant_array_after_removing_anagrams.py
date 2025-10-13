from find_resultant_array_after_removing_anagrams import Solution

def test_examples():
    s = Solution()
    assert s.removeAnagrams(["abba","baba","bbaa","cd","cd"]) == ["abba","cd"]
    assert s.removeAnagrams(["a","b","c","d","e"]) == ["a","b","c","d","e"]

def test_single_and_empty():
    s = Solution()
    assert s.removeAnagrams(["abc"]) == ["abc"]
    assert s.removeAnagrams([]) == []

def test_mixed_runs():
    s = Solution()
    assert s.removeAnagrams(["abc","cba","bac","abcd","dcba","ab","ba","ab"]) == ["abc","abcd","ab"]

def test_repeated_same_word():
    s = Solution()
    assert s.removeAnagrams(["aa","aa","aa"]) == ["aa"]