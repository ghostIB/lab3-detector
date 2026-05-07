# Лабораторна робота №3: Профілювання та дебаг Go-сервісу

## Виконані завдання:
1. **Аналіз пам'яті (Memory Leak):** Виявлено витік у `LeakCache` за допомогою `pprof`.
   <img width="1004" height="445" alt="image" src="https://github.com/user-attachments/assets/bb3c9fcc-d613-44f5-8ce6-eeca8a870f74" />

   <img width="1004" height="473" alt="image" src="https://github.com/user-attachments/assets/5a17d313-7b1f-4fc2-8a77-55a717d18ec6" />

   <img width="1004" height="292" alt="image" src="https://github.com/user-attachments/assets/7feb48cf-4afb-440a-9c2c-e617b8db9ec1" />

2. **Race Condition:** Виправлено падіння сервісу за допомогою `sync.Mutex`.
   <img width="1004" height="227" alt="image" src="https://github.com/user-attachments/assets/ae707386-7904-4d3b-9f41-7b20d031665f" />
   <img width="1004" height="227" alt="image" src="https://github.com/user-attachments/assets/b097499b-5aed-4dc5-b433-e162adae0682" />

3. **Оптимізація CPU:** Винесено компіляцію `regexp` у глобальну змінну, що значно знизило навантаження.
   <img width="1004" height="296" alt="image" src="https://github.com/user-attachments/assets/6f5591d4-89c3-4dbd-9f2e-2ecc0eac7245" />
   <img width="1004" height="289" alt="image" src="https://github.com/user-attachments/assets/6fcaa9dc-74bd-4b33-90de-62d40e89285b" />

4. **Debugging:** Налаштовано `dlv` у режимі headless для віддаленого підключення.
   <img width="1004" height="96" alt="image" src="https://github.com/user-attachments/assets/d1d8234f-9955-4602-862c-7e14775261b5" />
   <img width="1004" height="516" alt="image" src="https://github.com/user-attachments/assets/d7f126ad-ccc9-4f61-94f0-4426fb400061" />

