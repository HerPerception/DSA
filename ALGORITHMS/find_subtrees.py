# First implementation. Provided by AI. At the time this commit was made, I'm yet to wholly understand the implementation. 
# 1133ms runtime and 111.31MB memory use.
class Solution:

    def findSubtreeSizes(self, parent: List[int], s: str) -> List[int]:
        n = len(parent)
        children = [[] for _ in range(n)]

        for i in range(1, n):
            children[parent[i]].append(i)

        last_seen = {}
        new_parent = list(parent)
        def dfs(node):
            letter = s[node]
            if letter in last_seen:
                new_parent[node] = last_seen[letter]
            seen = last_seen.get(letter)
            last_seen[letter] = node

            for child in children[node]:
                dfs(child)

            if seen is None:
                del last_seen[letter]
            else:
                last_seen[letter] = seen
        dfs(0)
        new_children = [[] for _ in range(n)]
        for  i in range(1, n):
            new_children[new_parent[i]].append(i)
        answer = [1] * n
        def count_sizes(node):
            for child in new_children[node]:
                count_sizes(child)               # First, count the child's subtree
                answer[node] += answer[child]    # Then, add that total to the parent

        count_sizes(0)  # Start counting from the CEO (node 0)
        return answer
