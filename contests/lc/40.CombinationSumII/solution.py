class Solution:
    def combinationSum2(self, candidates: list[int], target: int) -> list[list[int]]:
        res = []
        candidates.sort() # Step 1: Sort to handle duplicates and allow early pruning

        def backtrack(start, current_target, path):
            if current_target == 0: # Step 2: Base case - target reached
                res.append(list(path))
                return

            for i in range(start, len(candidates)):
                # Step 3: Skip duplicate elements at the same recursion level
                if i > start and candidates[i] == candidates[i-1]:
                    continue

                # Step 4: Pruning
                if candidates[i] > current_target:
                    break

                path.append(candidates[i]) # Step 5: Recurse
                backtrack(i + 1, current_target - candidates[i], path)
                path.pop() # Step 6: Backtrack

        backtrack(0, target, [])
        return res
