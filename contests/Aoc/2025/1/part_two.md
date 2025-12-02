## --- Part Two ---

You're sure that's the right password, but the door won't open. You knock, but nobody answers. You build a snowman while you think.

As you're rolling the snowballs for your snowman, you find another security document that must have fallen into the snow:

> "Due to newer security protocols, please use password method `0x434C49434B` until further notice."

You remember from the training seminar that "method `0x434C49434B`" means you're actually supposed to count the number of times any click causes the dial to point at 0, regardless of whether it happens during a rotation or at the end of one.

Following the same rotations as in the above example, the dial points at zero a few extra times during its rotations:

- The dial starts by pointing at `50`.
- The dial is rotated `L68` to point at `82`; during this rotation, it points at 0 once.
- The dial is rotated `L30` to point at `52`.
- The dial is rotated `R48` to point at `0`.
- The dial is rotated `L5` to point at `95`.
- The dial is rotated `R60` to point at `55`; during this rotation, it points at `0` once.
- The dial is rotated `L55` to point at 0.
- The dial is rotated `L1` to point at `99`.
- The dial is rotated `L99` to point at `0`.
- The dial is rotated `R14` to point at `14`.
- The dial is rotated `L82` to point at `32`; during this rotation, it points at `0` once.

In this example, the dial points at `0` three times at the end of a rotation, plus three more times during a rotation. So, in this example, the new password would be `6`.

**Be careful**: if the dial were pointing at 50, a single rotation like R1000 would cause the dial to point at 0 ten times before returning back to 50!

Using password method **0x434C49434B**, what is the password to open the door?

> `0x434C49434B` → **CLICK** → "count every click, when the pointer is reaching 0"

-----
Вы уверены, что это правильный пароль, но дверь не открывается. Вы стучите — никто не отвечает. Пока думаете, вы строите снеговика.

Пока катаете комья для снеговика, вы находите ещё один документ безопасности, который, видимо, упал в снег:

> "Из-за новых протоколов безопасности, пожалуйста, используйте метод пароля 0x434C49434B до дальнейших указаний."

Вы вспоминаете, что на семинаре по технике безопасности вам объясняли: "метод 0x434C49434B" означает, что теперь нужно считать каждый щелчок, при котором указатель останавливается на 0. Это включает промежуточные остановки во время вращения, а не только конечное положение.

### Пример

Рассмотрим, как это работает на примере:

1. Указатель начинает с позиции `50`.
2. `L68`: указатель перемещается на `82`, и **проходит через 0 один раз** во время вращения.
3. `L30`: указатель перемещается на **52**.
4. R48: указатель перемещается на 0 (остановка в конце вращения).
5. L5: указатель перемещается на 95.
6. R60: указатель перемещается на 55, и проходит через 0 один раз во время вращения.
7. L55: указатель перемещается на 0 (остановка в конце вращения).
8. L1: указатель перемещается на 99.
9. L99: указатель перемещается на 0 (остановка в конце вращения).
10. R14: указатель перемещается на 14.
11. L82: указатель перемещается на 32, и проходит через 0 один раз во время вращения.

**Результат**:
- 3 остановки на 0 в конце вращений.
- 3 остановки на 0 во время вращений.
- Итого: 6 — это и есть новый пароль.

**Особое замечание:**

Если указатель начнёт с позиции `50` и выполнит вращение `R1000`, он пройдёт через `0` **десять раз**, прежде чем вернётся к `50`.

**Задача**:

Используя **метод** `0x434C49434B`, вычислите, сколько всего раз указатель проходит через `0` (включая конечные и промежуточные позиции) при выполнении всех вращений из вашего входного файла.

> `0x434C49434B` → **CLICK** → "считайте каждый щелчок, при котором указатель останавливается на 0"
