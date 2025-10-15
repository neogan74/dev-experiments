from maximum_total_damage_with_spell_casting import maximumTotalDamage

def test_examples():
    assert maximumTotalDamage([1,1,3,4]) == 6       # берём 1,1 и 4
    assert maximumTotalDamage([7,1,6,6]) == 13      # берём 1 и 6,6

def test_edges():
    assert maximumTotalDamage([5]) == 5
    assert maximumTotalDamage([2,2,2]) == 6         # все 2
    assert maximumTotalDamage([1,2,3,4,5]) >= 6     # напр., 1,4,5 нельзя вместе, лучший 1+4=5 или 2+5=7

def test_custom():
    # близкие значения конфликтуют: можно взять только один кластер каждые 3
    assert maximumTotalDamage([1,1,2,3,3,4,7,7]) == max(1+1+4, 3+3) + 7+7  # = max(6,6) + 14 = 20