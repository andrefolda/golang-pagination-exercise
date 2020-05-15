# golang-pagination-exercise

## Exercise
We want a footer pagination so the user can navigate through the webapp.

We have these variables available, you need to use them:
- current_page - actual page
- total_pages - total available pages
- boundaries - how many pages we want to link in the beginning and in the end
- around - how many pages we want to link before and after the actual page

For the pages that we don’t want to link we want to show “…”

#### Examples:

current_page = 4; total_pages = 5; boundaries = 1; around = 0

`
1 ... 4 5
`

current_page = 4; total_pages = 10; boundaries = 2; around = 2

`
1 2 3 4 5 6 ... 9 10
`

This exercise have to print the result.

**You need to do unit testing on your solution**

HINTS:
 - Focus on what we’re asking, don’t play around.
 - Keep it simple, don’t over-engineer. If you’re writing a lot of code, maybe you’re doing it
wrong.
 - Keep it readable, we’re not playing golf. Code must be readable.
 - Test your solution, play with corner-cases. We’ll play your code against our unit test
engine.
