# Лабораторна робота №4: Командна розробка та GitHub Workflow

**Тема:** Стратегії розгалуження, Code Review та захист релізу.
**Мета:** Освоїти індустріальні стандарти командної розробки: налаштування захищених гілок (Branch Protection), застосування стратегії Trunk-based Development та впровадження культури Code Review для забезпечення стабільності main-гілки.

---

## Етап 1: Налаштування Branch Protection
Для головної гілки `main` у налаштуваннях репозиторію GitHub було встановлено правила захисту (Branch Protection Rules). Це необхідно для запобігання випадковим змінам та забезпечення стабільності коду.

**Налаштовані правила:**
- **Require a pull request before merging:** заборонено прямий пуш у гілку main.
- **Require approvals:** необхідне підтвердження змін від іншого учасника (або імітація процесу через PR).
- **Require conversation resolution before merging:** усі гілки обговорення в коді мають бути закриті (вирішені).

> <img width="1004" height="466" alt="image" src="https://github.com/user-attachments/assets/9c6b1687-8e93-45fd-8dc0-5adff69c7fa6" />
<img width="916" height="769" alt="image" src="https://github.com/user-attachments/assets/326be0a9-9cb9-4868-888a-31b84202a670" />



---

## Етап 2: Створення Feature-гілки та розробка
Згідно зі стратегією Trunk-based Development, для впровадження нового функціоналу було створено короткоживучу гілку `feat-improve-logging`.

**Внесені зміни:**
У файл `internal/processor/metadata.go` додано логіку лічильника оброблених зображень та логування прогресу кожні 100 ітерацій.

```go
// Приклад доданого коду:
processedCount++
if processedCount % 100 == 0 {
    log.Printf("Статус: Успішно оброблено %d зображень", processedCount)
}
```

> <img width="1004" height="496" alt="image" src="https://github.com/user-attachments/assets/7e857534-9876-415d-8aa1-d56ced62fcd7" />
<img width="1004" height="473" alt="image" src="https://github.com/user-attachments/assets/50041b5e-3e85-4cde-a17c-57e2b17a5e36" />
<img width="1004" height="473" alt="image" src="https://github.com/user-attachments/assets/38f76ad3-fa92-4347-9601-9bd31793909e" />


---

## Етап 3: Pull Request та Code Review
Було створено Pull Request (PR) для злиття гілки `feat-improve-logging` у `main`. У рамках цього етапу проведено імітацію процесу Code Review.

**Під час рев'ю:**
1. Залишено коментар щодо винесення магічного числа `100` у константу для покращення читаємості.
2. Запрошено додавання Unit-тесту для перевірки коректності роботи лічильника.

**Виправлення:**
Зауваження були опрацьовані, код виправлено у feature-гілці, після чого PR було схвалено (Approve).

> <img width="1004" height="472" alt="image" src="https://github.com/user-attachments/assets/6b775b9c-27ba-4ef3-a58b-da01c6e92dbf" />


---


## Відповіді на контрольні питання

1. **Чому Trunk-based Development вважається кращим для CI/CD, ніж класичний Git-flow?**
   Ця стратегія використовує короткоживучі гілки та часте злиття в `main`. Це забезпечує швидке виявлення конфліктів інтеграції, постійну готовність коду до релізу та мінімізує накопичення "технічного боргу".

2. **Що таке "Linear History" і чому Squash Merge допомагає її підтримувати?**
   Linear History — це історія комітів без складних розгалужень та зайвих merge-комітів. Squash Merge об'єднує всі коміти з feature-гілки в один чистий коміт, що робить історію `main` зрозумілою та легкою для читання.

3. **Чим небезпечний "Force Push" у спільну гілку і в яких випадках він допустимий?**
   Force push перезаписує історію гілки на сервері. У спільній гілці це призведе до втрати комітів інших розробників. Допустимий лише у власних feature-гілках для очищення історії перед мерджем.

4. **Яка різниця між статусами Comment, Approve та Request Changes у процесі Code Review?**
   - **Comment:** звичайне зауваження або порада, не блокує злиття.
   - **Approve:** схвалення змін, дає дозвіл на мердж.
   - **Request Changes:** вимога обов'язкових виправлень, блокує злиття до моменту їх внесення.
