from typing import List

def minTime(skill: List[int], mana: List[int]) -> int:
    """
    No-wait flow shop с временем обработки p[i][j] = skill[i] * mana[j].
    Идея: обрабатываем «столбцы» (potions) слева направо и «спускаемся» по магам снизу вверх,
    поддерживая два значения:
      - prevWizardDone: время завершения ПОСЛЕДНЕГО мага на текущем зелье j
      - prevPotionDone: время завершения (i+1)-го мага на предыдущем зелье j-1
    Переход:
      prevPotionDone -= skill[i+1] * mana[j-1]
      prevWizardDone = max(prevPotionDone, prevWizardDone - skill[i] * mana[j])
    а затем добавляем суммарное время текущего столбца.
    """
    n, m = len(skill), len(mana)
    if n == 0 or m == 0:
        return 0

    sumSkill = sum(skill)
    prevWizardDone = sumSkill * mana[0]  # время окончания зелья 0 у последнего мага

    for j in range(1, m):
        prevPotionDone = prevWizardDone
        for i in range(n - 2, -1, -1):
            # earliest start на маге i для зелья j с учётом no-wait ограничений
            prevPotionDone -= skill[i + 1] * mana[j - 1]
            prevWizardDone = max(prevPotionDone, prevWizardDone - skill[i] * mana[j])
        prevWizardDone += sumSkill * mana[j]

    return prevWizardDone
