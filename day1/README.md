# Day 1 — String Manipulation

**Theme:** Closing the "Fear of String Problems"  
**Goal:** Familiarize with Go's built-in string functions instead of overcomplicating problems.

---

## Problems Solved

### LC7 — Reverse Integer
**Approach:** Convert int to string, reverse the byte slice, convert back to int.  
**Key lessons:**
- Use `strconv.Itoa` / `strconv.Atoi` to bridge int ↔ string.
- Strip the `-` sign before reversing, then restore it after — avoids sign getting buried in the middle.
- Always check 32-bit overflow bounds (`math.Pow(2, 31)`) after reversing; return `0` if exceeded.
- `strings.Trim` can cleanly remove a specific character (like `"-"`) from both ends.

---

### LC125 — Valid Palindrome
**Approach:** Normalize → filter → two-pointer check.  
**Key lessons:**
- Always normalize first: `strings.ToLower` + `strings.TrimSpace` before any logic.
- Use `[]rune` (not `[]byte`) when iterating strings — handles multi-byte/Unicode characters safely.
- Filter non-alphanumeric chars with a simple range check (`'a'`–`'z'`, `'0'`–`'9'`) rather than regex, keeping it fast and readable.
- Two-pointer (`left`, `right` walking inward) is the idiomatic palindrome check — O(n) time, O(1) extra space on the cleaned slice.

---

### LC151 — Reverse Words in a String
**Approach:** Trim → split on whitespace → fill result slice in reverse → join.  
**Key lessons:**
- `strings.Fields` is the right tool for splitting on arbitrary whitespace (handles multiple spaces, leading/trailing spaces automatically) — prefer it over `strings.Split(s, " ")`.
- Filling a pre-allocated slice from the end (`iterator = len-1; iterator--`) is cleaner than reversing after the fact.
- `strings.Join(slice, " ")` collapses everything back with a single space, giving the expected output without extra trimming.

---

## General Takeaways

| Situation | Preferred tool |
|---|---|
| int ↔ string | `strconv.Itoa` / `strconv.Atoi` |
| Case normalization | `strings.ToLower` / `strings.ToUpper` |
| Whitespace handling | `strings.TrimSpace`, `strings.Fields` |
| Split on delimiter | `strings.Split` (known delim) or `strings.Fields` (whitespace) |
| Rejoin slice | `strings.Join` |
| Char-by-char iteration | `range str` over `[]rune`, not `[]byte` |
| Palindrome / two-pointer | Walk `left` and `right` inward, stop when they meet |

> String problems rarely need manual pointer arithmetic from scratch — lean on the standard library first, then optimize only if needed.
