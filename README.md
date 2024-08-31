# Uy vazifasi: Go da Worker Pools ni Amalga Oshirish

## Maqsad
Bu vazifaning maqsadi - `Go` da `worker pool`ni amalga oshirish orqali ko'p sonli vazifalarni bir vaqtda bajarish, resurslardan samarali foydalanish va ish samaradorligini oshirish.

## Talablar
- Gorutinlar va kanallardan foydalanib, bir vaqtda vazifalarni bajara oladigan oddiy `worker pool` ni amalga oshiring.
- Vazifa muvaffaqiyatsiz bo'lgan taqdirda, xatolarni boshqarishni amalga oshiring.
- Worker pool yuklashga qarab ishchilar sonini dinamik tarzda sozlashi kerak.
- `Redis`'da vazifa natijalari yoki xatolarni saqlash uchun foydalaning.
- `Go` ning `net/http` paketidan foydalanib, vazifalarni yuborish va ularning holatini tekshirish uchun `RESTful API` interfeysini amalga oshiring.

## Batafsil Bosqichlar
1. **Vazifalarni Yuborish API**: 
    - Clientga vazifa yuborishga ruxsat beruvchi `HTTP` interfeysini (`/submit-task`) amalga oshiring.
    - Har bir vazifa `worker pool` tomonidan qayta ishlanishi kerak.

2. **Worker Pool Amalga Oshirish**: 
    - Goroutinlar yordamida `worker pool` yarating.
    - Kanallardan foydalanib, vazifalarni taqsimlash va workerlar bilan asosiy jarayon o'rtasida muloqotni boshqaring.
    - Navbatdagi vazifalar soniga qarab `worker pool` ni dinamik ravishda kengaytirishni amalga oshiring.

3. **Xatolarni Boshqarish va Qayta Urinishlar**:
    - Vazifalar muvaffaqiyatsiz bo'lganida qayta urinishlar sonini ko'rsatish va boshqarish mexanizmini amalga oshiring
    - Muvaffaqiyatsiz bo'lgan vazifa tafsilotlarini tahlil qilish uchun `Redis` da saqlang.

4. **Vazifa Holati API**:
    - Ma'lum bir vazifaning holatini tekshirish uchun `API` interfeysini (`/task-status/{id}`) amalga oshiring.
    - Vazifa holati `Redis`'dan olinishi kerak.

5. **Joylashtirish**:
    - Dasturni `Docker` yordamida joylashtiring.